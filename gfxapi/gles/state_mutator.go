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

func (ϟa *GlBlendBarrierKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_0_ext := ExtensionId_GL_KHR_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_0_ext
	return nil
}
func (ϟa *GlBlendEquationSeparateiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_1_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_1_ext
	return nil
}
func (ϟa *GlBlendEquationiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_2_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_2_ext
	return nil
}
func (ϟa *GlBlendFuncSeparateiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_3_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_3_ext
	return nil
}
func (ϟa *GlBlendFunciEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_4_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_4_ext
	return nil
}
func (ϟa *GlColorMaskiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_5_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_5_ext
	return nil
}
func (ϟa *GlCopyImageSubDataEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_6_ext := ExtensionId_GL_EXT_copy_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_6_ext
	return nil
}
func (ϟa *GlDebugMessageCallbackKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_7_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_7_ext
	return nil
}
func (ϟa *GlDebugMessageControlKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_8_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_8_ext
	return nil
}
func (ϟa *GlDebugMessageInsertKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_9_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_9_ext
	return nil
}
func (ϟa *GlDisableiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_10_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_10_ext
	return nil
}
func (ϟa *GlEnableiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_11_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_11_ext
	return nil
}
func (ϟa *GlFramebufferTextureEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_12_ext := ExtensionId_GL_EXT_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_12_ext
	return nil
}
func (ϟa *GlGetDebugMessageLogKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_13_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_13_ext
	return nil
}
func (ϟa *GlGetObjectLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_14_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_14_ext
	return nil
}
func (ϟa *GlGetObjectPtrLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_15_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_15_ext
	return nil
}
func (ϟa *GlGetPointervKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_16_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_16_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_17_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_17_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_18_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_18_ext
	return nil
}
func (ϟa *GlGetTexParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_19_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_19_ext
	return nil
}
func (ϟa *GlGetTexParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_20_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_20_ext
	return nil
}
func (ϟa *GlIsEnablediEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_21_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_21_ext
	return nil
}
func (ϟa *GlMinSampleShadingOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_22_ext := ExtensionId_GL_OES_sample_shading // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_22_ext
	return nil
}
func (ϟa *GlObjectLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_23_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_23_ext
	return nil
}
func (ϟa *GlObjectPtrLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_24_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_24_ext
	return nil
}
func (ϟa *GlPatchParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_25_ext := ExtensionId_GL_EXT_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_25_ext
	return nil
}
func (ϟa *GlPopDebugGroupKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_26_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_26_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_27_ext := ExtensionId_GL_EXT_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_27_ext
	return nil
}
func (ϟa *GlPushDebugGroupKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_28_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_28_ext
	return nil
}
func (ϟa *GlSamplerParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_29_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_29_ext
	return nil
}
func (ϟa *GlSamplerParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_30_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_30_ext
	return nil
}
func (ϟa *GlTexBufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_31_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_31_ext
	return nil
}
func (ϟa *GlTexBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_32_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_32_ext
	return nil
}
func (ϟa *GlTexParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_33_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_33_ext
	return nil
}
func (ϟa *GlTexParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_34_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_34_ext
	return nil
}
func (ϟa *GlTexStorage3DMultisampleOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_35_ext := ExtensionId_GL_OES_texture_storage_multisample_2d_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_35_ext
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
		glErrorInvalidEnum_38_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_38_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_36_major, minRequiredVersion_36_minor
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_39_major := uint32(3)                               // u32
	minRequiredVersion_39_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_41_msg := "No context bound" // string
		return
		_ = error_41_msg
	}
	GetContext_40_result := context // Contextʳ
	ctx := GetContext_40_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_39_major, minRequiredVersion_39_minor, q, context, GetContext_40_result, ctx
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_42_major := uint32(3) // u32
	minRequiredVersion_42_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_43_major := uint32(3) // u32
		minRequiredVersion_43_minor := uint32(2) // u32
		_, _ = minRequiredVersion_43_major, minRequiredVersion_43_minor
	default:
		glErrorInvalidEnum_44_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_44_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_42_major, minRequiredVersion_42_minor
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_45_major := uint32(3)                               // u32
	minRequiredVersion_45_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_47_msg := "No context bound" // string
		return
		_ = error_47_msg
	}
	GetContext_46_result := context // Contextʳ
	ctx := GetContext_46_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = &Query{}
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_45_major, minRequiredVersion_45_minor, q, context, GetContext_46_result, ctx
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_48_major := uint32(3) // u32
	minRequiredVersion_48_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_QUERY_RESULT, GLenum_GL_QUERY_RESULT_AVAILABLE:
	default:
		glErrorInvalidEnum_49_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_49_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_48_major, minRequiredVersion_48_minor
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_50_major := uint32(3) // u32
	minRequiredVersion_50_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_51_major := uint32(3) // u32
		minRequiredVersion_51_minor := uint32(2) // u32
		_, _ = minRequiredVersion_51_major, minRequiredVersion_51_minor
	default:
		glErrorInvalidEnum_52_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_52_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_CURRENT_QUERY:
	default:
		glErrorInvalidEnum_53_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_53_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_50_major, minRequiredVersion_50_minor
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_54_major := uint32(3)     // u32
	minRequiredVersion_54_minor := uint32(0)     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_56_msg := "No context bound" // string
		return
		_ = error_56_msg
	}
	GetContext_55_result := context // Contextʳ
	ctx := GetContext_55_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_54_major, minRequiredVersion_54_minor, context, GetContext_55_result, ctx
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_57_major := uint32(2) // u32
	minRequiredVersion_57_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_58_major := uint32(3) // u32
		minRequiredVersion_58_minor := uint32(0) // u32
		_, _ = minRequiredVersion_58_major, minRequiredVersion_58_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_59_major := uint32(3) // u32
		minRequiredVersion_59_minor := uint32(1) // u32
		_, _ = minRequiredVersion_59_major, minRequiredVersion_59_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_60_major := uint32(3) // u32
		minRequiredVersion_60_minor := uint32(2) // u32
		_, _ = minRequiredVersion_60_major, minRequiredVersion_60_minor
	default:
		glErrorInvalidEnum_61_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_61_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_63_msg := "No context bound" // string
		return
		_ = error_63_msg
	}
	GetContext_62_result := context // Contextʳ
	ctx := GetContext_62_result     // Contextʳ
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_57_major, minRequiredVersion_57_minor, context, GetContext_62_result, ctx
	return nil
}
func (ϟa *GlBindBufferBase) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_64_major := uint32(3) // u32
	minRequiredVersion_64_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_65_major := uint32(3) // u32
		minRequiredVersion_65_minor := uint32(1) // u32
		_, _ = minRequiredVersion_65_major, minRequiredVersion_65_minor
	default:
		glErrorInvalidEnum_66_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_66_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_64_major, minRequiredVersion_64_minor
	return nil
}
func (ϟa *GlBindBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_67_major := uint32(3) // u32
	minRequiredVersion_67_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_68_major := uint32(3) // u32
		minRequiredVersion_68_minor := uint32(1) // u32
		_, _ = minRequiredVersion_68_major, minRequiredVersion_68_minor
	default:
		glErrorInvalidEnum_69_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_69_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_67_major, minRequiredVersion_67_minor
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_70_major := uint32(2) // u32
	minRequiredVersion_70_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_71_major := uint32(3) // u32
		minRequiredVersion_71_minor := uint32(0) // u32
		_, _ = minRequiredVersion_71_major, minRequiredVersion_71_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_72_major := uint32(3) // u32
		minRequiredVersion_72_minor := uint32(1) // u32
		_, _ = minRequiredVersion_72_major, minRequiredVersion_72_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_73_major := uint32(3) // u32
		minRequiredVersion_73_minor := uint32(2) // u32
		_, _ = minRequiredVersion_73_major, minRequiredVersion_73_minor
	default:
		glErrorInvalidEnum_74_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_74_param
	}
	switch ϟa.Usage {
	case GLenum_GL_DYNAMIC_DRAW, GLenum_GL_STATIC_DRAW, GLenum_GL_STREAM_DRAW:
	case GLenum_GL_DYNAMIC_COPY, GLenum_GL_DYNAMIC_READ, GLenum_GL_STATIC_COPY, GLenum_GL_STATIC_READ, GLenum_GL_STREAM_COPY, GLenum_GL_STREAM_READ:
		minRequiredVersion_75_major := uint32(3) // u32
		minRequiredVersion_75_minor := uint32(0) // u32
		_, _ = minRequiredVersion_75_major, minRequiredVersion_75_minor
	default:
		glErrorInvalidEnum_76_param := ϟa.Usage // GLenum
		return
		_ = glErrorInvalidEnum_76_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_78_msg := "No context bound" // string
		return
		_ = error_78_msg
	}
	GetContext_77_result := context       // Contextʳ
	ctx := GetContext_77_result           // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target) // BufferId
	b := ctx.Instances.Buffers.Get(id)    // Bufferʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_70_major, minRequiredVersion_70_minor, context, GetContext_77_result, ctx, id, b
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_79_major := uint32(2) // u32
	minRequiredVersion_79_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_80_major := uint32(3) // u32
		minRequiredVersion_80_minor := uint32(0) // u32
		_, _ = minRequiredVersion_80_major, minRequiredVersion_80_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_81_major := uint32(3) // u32
		minRequiredVersion_81_minor := uint32(1) // u32
		_, _ = minRequiredVersion_81_major, minRequiredVersion_81_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_82_major := uint32(3) // u32
		minRequiredVersion_82_minor := uint32(2) // u32
		_, _ = minRequiredVersion_82_major, minRequiredVersion_82_minor
	default:
		glErrorInvalidEnum_83_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_83_param
	}
	ϟa.Data.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_79_major, minRequiredVersion_79_minor
	return nil
}
func (ϟa *GlCopyBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_84_major := uint32(3) // u32
	minRequiredVersion_84_minor := uint32(0) // u32
	switch ϟa.ReadTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_85_major := uint32(3) // u32
		minRequiredVersion_85_minor := uint32(2) // u32
		_, _ = minRequiredVersion_85_major, minRequiredVersion_85_minor
	default:
		glErrorInvalidEnum_86_param := ϟa.ReadTarget // GLenum
		return
		_ = glErrorInvalidEnum_86_param
	}
	switch ϟa.WriteTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_87_major := uint32(3) // u32
		minRequiredVersion_87_minor := uint32(2) // u32
		_, _ = minRequiredVersion_87_major, minRequiredVersion_87_minor
	default:
		glErrorInvalidEnum_88_param := ϟa.WriteTarget // GLenum
		return
		_ = glErrorInvalidEnum_88_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_84_major, minRequiredVersion_84_minor
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_89_major := uint32(2)                               // u32
	minRequiredVersion_89_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_91_msg := "No context bound" // string
		return
		_ = error_91_msg
	}
	GetContext_90_result := context // Contextʳ
	ctx := GetContext_90_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Buffers, b.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_89_major, minRequiredVersion_89_minor, b, context, GetContext_90_result, ctx
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_92_major := uint32(2)                               // u32
	minRequiredVersion_92_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_94_msg := "No context bound" // string
		return
		_ = error_94_msg
	}
	GetContext_93_result := context // Contextʳ
	ctx := GetContext_93_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // BufferId
		ctx.Instances.Buffers[id] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
		b.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_92_major, minRequiredVersion_92_minor, b, context, GetContext_93_result, ctx
	return nil
}
func (ϟa *GlGetBufferParameteri64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_95_major := uint32(3) // u32
	minRequiredVersion_95_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_96_major := uint32(3) // u32
		minRequiredVersion_96_minor := uint32(2) // u32
		_, _ = minRequiredVersion_96_major, minRequiredVersion_96_minor
	default:
		glErrorInvalidEnum_97_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_97_param
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET, GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	default:
		glErrorInvalidEnum_98_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_98_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_95_major, minRequiredVersion_95_minor
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_99_major := uint32(2) // u32
	minRequiredVersion_99_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_100_major := uint32(3) // u32
		minRequiredVersion_100_minor := uint32(0) // u32
		_, _ = minRequiredVersion_100_major, minRequiredVersion_100_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_101_major := uint32(3) // u32
		minRequiredVersion_101_minor := uint32(2) // u32
		_, _ = minRequiredVersion_101_major, minRequiredVersion_101_minor
	default:
		glErrorInvalidEnum_102_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_102_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET:
		minRequiredVersion_103_major := uint32(3) // u32
		minRequiredVersion_103_minor := uint32(0) // u32
		_, _ = minRequiredVersion_103_major, minRequiredVersion_103_minor
	default:
		glErrorInvalidEnum_104_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_104_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_106_msg := "No context bound" // string
		return
		_ = error_106_msg
	}
	GetContext_105_result := context      // Contextʳ
	ctx := GetContext_105_result          // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target) // BufferId
	b := ctx.Instances.Buffers.Get(id)    // Bufferʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_99_major, minRequiredVersion_99_minor, context, GetContext_105_result, ctx, id, b
	return nil
}
func (ϟa *GlGetBufferPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_107_major := uint32(3) // u32
	minRequiredVersion_107_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_108_major := uint32(3) // u32
		minRequiredVersion_108_minor := uint32(2) // u32
		_, _ = minRequiredVersion_108_major, minRequiredVersion_108_minor
	default:
		glErrorInvalidEnum_109_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_109_param
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_MAP_POINTER:
	default:
		glErrorInvalidEnum_110_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_110_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_107_major, minRequiredVersion_107_minor
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_111_major := uint32(2)    // u32
	minRequiredVersion_111_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_113_msg := "No context bound" // string
		return
		_ = error_113_msg
	}
	GetContext_112_result := context // Contextʳ
	ctx := GetContext_112_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Buffers.Contains(ϟa.Buffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_111_major, minRequiredVersion_111_minor, context, GetContext_112_result, ctx
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_114_major := uint32(3) // u32
	minRequiredVersion_114_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_115_major := uint32(3) // u32
		minRequiredVersion_115_minor := uint32(1) // u32
		_, _ = minRequiredVersion_115_major, minRequiredVersion_115_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_116_major := uint32(3) // u32
		minRequiredVersion_116_minor := uint32(2) // u32
		_, _ = minRequiredVersion_116_major, minRequiredVersion_116_minor
	default:
		glErrorInvalidEnum_117_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_117_param
	}
	supportsBits_118_seenBits := ϟa.Access                                                                                                                                                                                                                                      // GLbitfield
	supportsBits_118_validBits := (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT) | ((GLbitfield_GL_MAP_READ_BIT) | ((GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT) | (GLbitfield_GL_MAP_WRITE_BIT))))) // GLbitfield
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
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_120_msg := "No context bound" // string
		return
		_ = error_120_msg
	}
	GetContext_119_result := context                                // Contextʳ
	ctx := GetContext_119_result                                    // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_114_major, minRequiredVersion_114_minor, supportsBits_118_seenBits, supportsBits_118_validBits, context, GetContext_119_result, ctx, b, ptr
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_121_major := uint32(3) // u32
	minRequiredVersion_121_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_122_major := uint32(3) // u32
		minRequiredVersion_122_minor := uint32(1) // u32
		_, _ = minRequiredVersion_122_major, minRequiredVersion_122_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_123_major := uint32(3) // u32
		minRequiredVersion_123_minor := uint32(2) // u32
		_, _ = minRequiredVersion_123_major, minRequiredVersion_123_minor
	default:
		glErrorInvalidEnum_124_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_124_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_126_msg := "No context bound" // string
		return
		_ = error_126_msg
	}
	GetContext_125_result := context                                // Contextʳ
	ctx := GetContext_125_result                                    // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	b.Data.Slice(uint64(b.MappingOffset), uint64((b.MappingOffset)+(int32(b.MappingData.Count))), ϟs).Copy(b.MappingData, ϟa, ϟs, ϟd, ϟl, ϟb)
	externs{ϟa, ϟs, ϟd, ϟl, ϟb}.unmapMemory(b.MappingData)
	b.MappingOffset = int32(0)
	b.MappingData = U8ˢ{}
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _ = minRequiredVersion_121_major, minRequiredVersion_121_minor, context, GetContext_125_result, ctx, b
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_127_major := uint32(2) // u32
	minRequiredVersion_127_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_128_major := uint32(3) // u32
		minRequiredVersion_128_minor := uint32(2) // u32
		_, _ = minRequiredVersion_128_major, minRequiredVersion_128_minor
	default:
		glErrorInvalidEnum_129_param := ϟa.DrawMode // GLenum
		return
		_ = glErrorInvalidEnum_129_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_131_msg := "No context bound" // string
		return
		_ = error_131_msg
	}
	GetContext_130_result := context                                                // Contextʳ
	ctx := GetContext_130_result                                                    // Contextʳ
	last_index := (uint32(ϟa.FirstIndex)) + ((uint32(ϟa.IndexCount)) - (uint32(1))) // u32
	ReadVertexArrays_132_ctx := ctx                                                 // Contextʳ
	ReadVertexArrays_132_first_index := uint32(ϟa.FirstIndex)                       // u32
	ReadVertexArrays_132_last_index := last_index                                   // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_132_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_132_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
		if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
			vertexAttribTypeSize_133_t := arr.Type // GLenum
			vertexAttribTypeSize_133_result := func() (result uint32) {
				switch vertexAttribTypeSize_133_t {
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
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_133_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_133_result) * (arr.Size) // u32
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
			for v := uint32(ReadVertexArrays_132_first_index); v < (ReadVertexArrays_132_last_index)+(uint32(1)); v++ {
				offset := (elstride) * (v) // u32
				arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_133_t, vertexAttribTypeSize_133_result, elsize, elstride
		}
		_ = arr
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_127_major, minRequiredVersion_127_minor, context, GetContext_130_result, ctx, last_index, ReadVertexArrays_132_ctx, ReadVertexArrays_132_first_index, ReadVertexArrays_132_last_index
	return nil
}
func (ϟa *GlDrawArraysIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_134_major := uint32(3) // u32
	minRequiredVersion_134_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_135_major := uint32(3) // u32
		minRequiredVersion_135_minor := uint32(2) // u32
		_, _ = minRequiredVersion_135_major, minRequiredVersion_135_minor
	default:
		glErrorInvalidEnum_136_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_136_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_134_major, minRequiredVersion_134_minor
	return nil
}
func (ϟa *GlDrawArraysInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_137_major := uint32(3) // u32
	minRequiredVersion_137_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_138_major := uint32(3) // u32
		minRequiredVersion_138_minor := uint32(2) // u32
		_, _ = minRequiredVersion_138_major, minRequiredVersion_138_minor
	default:
		glErrorInvalidEnum_139_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_139_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_137_major, minRequiredVersion_137_minor
	return nil
}
func (ϟa *GlDrawBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_140_major := uint32(3) // u32
	minRequiredVersion_140_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_140_major, minRequiredVersion_140_minor
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_141_major := uint32(2) // u32
	minRequiredVersion_141_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_142_major := uint32(3) // u32
		minRequiredVersion_142_minor := uint32(2) // u32
		_, _ = minRequiredVersion_142_major, minRequiredVersion_142_minor
	default:
		glErrorInvalidEnum_143_param := ϟa.DrawMode // GLenum
		return
		_ = glErrorInvalidEnum_143_param
	}
	switch ϟa.IndicesType {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_UNSIGNED_INT:
		minRequiredVersion_144_major := uint32(3) // u32
		minRequiredVersion_144_minor := uint32(0) // u32
		_, _ = minRequiredVersion_144_major, minRequiredVersion_144_minor
	default:
		glErrorInvalidEnum_145_param := ϟa.IndicesType // GLenum
		return
		_ = glErrorInvalidEnum_145_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_147_msg := "No context bound" // string
		return
		_ = error_147_msg
	}
	GetContext_146_result := context                           // Contextʳ
	ctx := GetContext_146_result                               // Contextʳ
	count := uint32(ϟa.ElementCount)                           // u32
	id := ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	if (id) != (BufferId(uint32(0))) {
		index_data := ctx.Instances.Buffers.Get(id).Data                                                           // U8ˢ
		offset := uint32(uint64(ϟa.Indices.Address))                                                               // u32
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count)  // u32
		ReadVertexArrays_148_ctx := ctx                                                                            // Contextʳ
		ReadVertexArrays_148_first_index := first                                                                  // u32
		ReadVertexArrays_148_last_index := last                                                                    // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_148_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_148_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_149_t := arr.Type // GLenum
				vertexAttribTypeSize_149_result := func() (result uint32) {
					switch vertexAttribTypeSize_149_t {
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
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_149_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_149_result) * (arr.Size) // u32
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
				for v := uint32(ReadVertexArrays_148_first_index); v < (ReadVertexArrays_148_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_149_t, vertexAttribTypeSize_149_result, elsize, elstride
			}
			_ = arr
		}
		_, _, _, _, _, _, _ = index_data, offset, first, last, ReadVertexArrays_148_ctx, ReadVertexArrays_148_first_index, ReadVertexArrays_148_last_index
	} else {
		index_data := U8ᵖ(ϟa.Indices)                                                               // U8ᵖ
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(index_data, ϟa.IndicesType, uint32(0), count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(index_data, ϟa.IndicesType, uint32(0), count)  // u32
		ReadVertexArrays_150_ctx := ctx                                                             // Contextʳ
		ReadVertexArrays_150_first_index := first                                                   // u32
		ReadVertexArrays_150_last_index := last                                                     // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_150_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_150_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_151_t := arr.Type // GLenum
				vertexAttribTypeSize_151_result := func() (result uint32) {
					switch vertexAttribTypeSize_151_t {
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
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_151_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_151_result) * (arr.Size) // u32
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
				for v := uint32(ReadVertexArrays_150_first_index); v < (ReadVertexArrays_150_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_151_t, vertexAttribTypeSize_151_result, elsize, elstride
			}
			_ = arr
		}
		IndexSize_152_indices_type := ϟa.IndicesType // GLenum
		IndexSize_152_result := func() (result uint32) {
			switch IndexSize_152_indices_type {
			case GLenum_GL_UNSIGNED_BYTE:
				return uint32(1)
			case GLenum_GL_UNSIGNED_SHORT:
				return uint32(2)
			case GLenum_GL_UNSIGNED_INT:
				return uint32(4)
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_152_indices_type, ϟa))
				return result
			}
		}() // u32
		index_data.Slice(uint64(uint32(0)), uint64((uint32(ϟa.ElementCount))*(IndexSize_152_result)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_150_ctx, ReadVertexArrays_150_first_index, ReadVertexArrays_150_last_index, IndexSize_152_indices_type, IndexSize_152_result
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_141_major, minRequiredVersion_141_minor, context, GetContext_146_result, ctx, count, id
	return nil
}
func (ϟa *GlDrawElementsBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_153_major := uint32(3) // u32
	minRequiredVersion_153_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
	default:
		glErrorInvalidEnum_154_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_154_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_155_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_155_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_153_major, minRequiredVersion_153_minor
	return nil
}
func (ϟa *GlDrawElementsIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_156_major := uint32(3) // u32
	minRequiredVersion_156_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_157_major := uint32(3) // u32
		minRequiredVersion_157_minor := uint32(2) // u32
		_, _ = minRequiredVersion_157_major, minRequiredVersion_157_minor
	default:
		glErrorInvalidEnum_158_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_158_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_159_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_159_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_156_major, minRequiredVersion_156_minor
	return nil
}
func (ϟa *GlDrawElementsInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_160_major := uint32(3) // u32
	minRequiredVersion_160_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_161_major := uint32(3) // u32
		minRequiredVersion_161_minor := uint32(2) // u32
		_, _ = minRequiredVersion_161_major, minRequiredVersion_161_minor
	default:
		glErrorInvalidEnum_162_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_162_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_163_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_163_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_160_major, minRequiredVersion_160_minor
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_164_major := uint32(3) // u32
	minRequiredVersion_164_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
	default:
		glErrorInvalidEnum_165_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_165_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_166_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_166_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_164_major, minRequiredVersion_164_minor
	return nil
}
func (ϟa *GlDrawRangeElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_167_major := uint32(3) // u32
	minRequiredVersion_167_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_168_major := uint32(3) // u32
		minRequiredVersion_168_minor := uint32(2) // u32
		_, _ = minRequiredVersion_168_major, minRequiredVersion_168_minor
	default:
		glErrorInvalidEnum_169_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_169_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_170_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_170_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_167_major, minRequiredVersion_167_minor
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_171_major := uint32(3) // u32
	minRequiredVersion_171_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
	default:
		glErrorInvalidEnum_172_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_172_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_173_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_173_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_171_major, minRequiredVersion_171_minor
	return nil
}
func (ϟa *GlPatchParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_174_major := uint32(3) // u32
	minRequiredVersion_174_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_PATCH_VERTICES:
	default:
		glErrorInvalidEnum_175_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_175_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_174_major, minRequiredVersion_174_minor
	return nil
}
func (ϟa *GlPrimitiveBoundingBox) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_176_major := uint32(3) // u32
	minRequiredVersion_176_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_176_major, minRequiredVersion_176_minor
	return nil
}
func (ϟa *GlActiveShaderProgramEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_177_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_177_ext
	return nil
}
func (ϟa *GlAlphaFuncQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_178_ext := ExtensionId_GL_QCOM_alpha_test // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_178_ext
	return nil
}
func (ϟa *GlApplyFramebufferAttachmentCMAAINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_179_ext := ExtensionId_GL_INTEL_framebuffer_CMAA // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_179_ext
	return nil
}
func (ϟa *GlBeginConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_180_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_180_ext
	return nil
}
func (ϟa *GlBeginPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_181_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_181_ext
	return nil
}
func (ϟa *GlBeginPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_182_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_182_ext
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_183_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_184_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_183_ext, requiresExtension_184_ext
	return nil
}
func (ϟa *GlBindProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_185_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_185_ext
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_186_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_188_msg := "No context bound" // string
		return
		_ = error_188_msg
	}
	GetContext_187_result := context // Contextʳ
	ctx := GetContext_187_result     // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = &VertexArray{}
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = requiresExtension_186_ext, context, GetContext_187_result, ctx
	return nil
}
func (ϟa *GlBlendBarrierNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_189_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_189_ext
	return nil
}
func (ϟa *GlBlendEquationSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_190_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_190_ext
	return nil
}
func (ϟa *GlBlendEquationiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_191_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_191_ext
	return nil
}
func (ϟa *GlBlendFuncSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_192_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_192_ext
	return nil
}
func (ϟa *GlBlendFunciOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_193_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_193_ext
	return nil
}
func (ϟa *GlBlendParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_194_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_194_ext
	return nil
}
func (ϟa *GlBlitFramebufferANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_195_ext := ExtensionId_GL_ANGLE_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_195_ext
	return nil
}
func (ϟa *GlBlitFramebufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_196_ext := ExtensionId_GL_NV_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_196_ext
	return nil
}
func (ϟa *GlBufferStorageEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_197_ext := ExtensionId_GL_EXT_buffer_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_197_ext
	return nil
}
func (ϟa *GlClientWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_198_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_198_ext
	return nil
}
func (ϟa *GlColorMaskiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_199_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_199_ext
	return nil
}
func (ϟa *GlCompressedTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_200_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_200_ext
	return nil
}
func (ϟa *GlCompressedTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_201_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_201_ext
	return nil
}
func (ϟa *GlCopyBufferSubDataNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_202_ext := ExtensionId_GL_NV_copy_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_202_ext
	return nil
}
func (ϟa *GlCopyImageSubDataOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_203_ext := ExtensionId_GL_OES_copy_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_203_ext
	return nil
}
func (ϟa *GlCopyPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_204_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_204_ext
	return nil
}
func (ϟa *GlCopyTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_205_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_205_ext
	return nil
}
func (ϟa *GlCopyTextureLevelsAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_206_ext := ExtensionId_GL_APPLE_copy_texture_levels // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_206_ext
	return nil
}
func (ϟa *GlCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_207_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_207_ext
	return nil
}
func (ϟa *GlCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_208_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_208_ext
	return nil
}
func (ϟa *GlCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_209_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_209_ext
	return nil
}
func (ϟa *GlCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_210_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_210_ext
	return nil
}
func (ϟa *GlCoverageMaskNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_211_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_211_ext
	return nil
}
func (ϟa *GlCoverageModulationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_212_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_212_ext
	return nil
}
func (ϟa *GlCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_213_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_213_ext
	return nil
}
func (ϟa *GlCoverageOperationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_214_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_214_ext
	return nil
}
func (ϟa *GlCreatePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_215_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_215_ext
	return nil
}
func (ϟa *GlCreateShaderProgramvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_216_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_216_ext
	return nil
}
func (ϟa *GlDeleteFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_217_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_217_ext
	return nil
}
func (ϟa *GlDeletePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_218_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_218_ext
	return nil
}
func (ϟa *GlDeletePerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_219_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_219_ext
	return nil
}
func (ϟa *GlDeletePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_220_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_220_ext
	return nil
}
func (ϟa *GlDeleteProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_221_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_221_ext
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_222_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_223_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_225_msg := "No context bound" // string
		return
		_ = error_225_msg
	}
	GetContext_224_result := context // Contextʳ
	ctx := GetContext_224_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_222_ext, requiresExtension_223_ext, q, context, GetContext_224_result, ctx
	return nil
}
func (ϟa *GlDeleteSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_226_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_226_ext
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_227_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_229_msg := "No context bound" // string
		return
		_ = error_229_msg
	}
	GetContext_228_result := context                                      // Contextʳ
	ctx := GetContext_228_result                                          // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = requiresExtension_227_ext, context, GetContext_228_result, ctx, a
	return nil
}
func (ϟa *GlDepthRangeArrayfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_230_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_230_ext
	return nil
}
func (ϟa *GlDepthRangeIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_231_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_231_ext
	return nil
}
func (ϟa *GlDisableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_232_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_232_ext
	return nil
}
func (ϟa *GlDisableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_233_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_233_ext
	return nil
}
func (ϟa *GlDisableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_234_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_234_ext
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_235_ext := ExtensionId_GL_EXT_discard_framebuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_235_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_236_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_236_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_237_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_237_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_238_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_239_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_238_ext, requiresExtension_239_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_240_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_240_ext
	return nil
}
func (ϟa *GlDrawBuffersEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_241_ext := ExtensionId_GL_EXT_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_241_ext
	return nil
}
func (ϟa *GlDrawBuffersIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_242_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_242_ext
	return nil
}
func (ϟa *GlDrawBuffersNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_243_ext := ExtensionId_GL_NV_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_243_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_244_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_244_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_245_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_245_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_246_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_246_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_247_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_247_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_248_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_248_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_249_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_249_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_250_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_250_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_251_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_252_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_251_ext, requiresExtension_252_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_253_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_253_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_254_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_254_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_255_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_255_ext
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_256_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_256_ext
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_257_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_257_ext
	return nil
}
func (ϟa *GlEnableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_258_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_258_ext
	return nil
}
func (ϟa *GlEnableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_259_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_259_ext
	return nil
}
func (ϟa *GlEnableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_260_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_260_ext
	return nil
}
func (ϟa *GlEndConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_261_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_261_ext
	return nil
}
func (ϟa *GlEndPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_262_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_262_ext
	return nil
}
func (ϟa *GlEndPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_263_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_263_ext
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_264_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_265_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_264_ext, requiresExtension_265_ext
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_266_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_266_ext
	return nil
}
func (ϟa *GlExtGetBufferPointervQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_267_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_267_ext
	return nil
}
func (ϟa *GlExtGetBuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_268_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_268_ext
	return nil
}
func (ϟa *GlExtGetFramebuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_269_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_269_ext
	return nil
}
func (ϟa *GlExtGetProgramBinarySourceQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_270_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_270_ext
	return nil
}
func (ϟa *GlExtGetProgramsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_271_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_271_ext
	return nil
}
func (ϟa *GlExtGetRenderbuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_272_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_272_ext
	return nil
}
func (ϟa *GlExtGetShadersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_273_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_273_ext
	return nil
}
func (ϟa *GlExtGetTexLevelParameterivQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_274_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_274_ext
	return nil
}
func (ϟa *GlExtGetTexSubImageQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_275_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_275_ext
	return nil
}
func (ϟa *GlExtGetTexturesQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_276_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_276_ext
	return nil
}
func (ϟa *GlExtIsProgramBinaryQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_277_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_277_ext
	return nil
}
func (ϟa *GlExtTexObjectStateOverrideiQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_278_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_278_ext
	return nil
}
func (ϟa *GlFenceSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_279_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_279_ext
	return nil
}
func (ϟa *GlFinishFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_280_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_280_ext
	return nil
}
func (ϟa *GlFlushMappedBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_281_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_281_ext
	return nil
}
func (ϟa *GlFragmentCoverageColorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_282_ext := ExtensionId_GL_NV_fragment_coverage_to_color // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_282_ext
	return nil
}
func (ϟa *GlFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_283_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_283_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_284_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_284_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_285_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_285_ext
	return nil
}
func (ϟa *GlFramebufferTexture3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_286_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_286_ext
	return nil
}
func (ϟa *GlFramebufferTextureMultiviewOVR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_287_ext := ExtensionId_GL_OVR_multiview // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_287_ext
	return nil
}
func (ϟa *GlFramebufferTextureOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_288_ext := ExtensionId_GL_OES_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_288_ext
	return nil
}
func (ϟa *GlGenFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_289_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_289_ext
	return nil
}
func (ϟa *GlGenPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_290_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_290_ext
	return nil
}
func (ϟa *GlGenPerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_291_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_291_ext
	return nil
}
func (ϟa *GlGenProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_292_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_292_ext
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_293_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_294_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_296_msg := "No context bound" // string
		return
		_ = error_296_msg
	}
	GetContext_295_result := context // Contextʳ
	ctx := GetContext_295_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = &Query{}
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = requiresExtension_293_ext, requiresExtension_294_ext, q, context, GetContext_295_result, ctx
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_297_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_299_msg := "No context bound" // string
		return
		_ = error_299_msg
	}
	GetContext_298_result := context // Contextʳ
	ctx := GetContext_298_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = &VertexArray{}
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _ = requiresExtension_297_ext, a, context, GetContext_298_result, ctx
	return nil
}
func (ϟa *GlGetBufferPointervOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_300_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_300_ext
	return nil
}
func (ϟa *GlGetCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_301_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_301_ext
	return nil
}
func (ϟa *GlGetDriverControlStringQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_302_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_302_ext
	return nil
}
func (ϟa *GlGetDriverControlsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_303_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_303_ext
	return nil
}
func (ϟa *GlGetFenceivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_304_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_304_ext
	return nil
}
func (ϟa *GlGetFirstPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_305_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_305_ext
	return nil
}
func (ϟa *GlGetFloati_vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_306_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_306_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_307_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_307_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_308_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_308_ext
	return nil
}
func (ϟa *GlGetImageHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_309_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_309_ext
	return nil
}
func (ϟa *GlGetInteger64vAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_310_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_310_ext
	return nil
}
func (ϟa *GlGetIntegeri_vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_311_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_311_ext
	return nil
}
func (ϟa *GlGetInternalformatSampleivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_312_ext := ExtensionId_GL_NV_internalformat_sample_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_312_ext
	return nil
}
func (ϟa *GlGetNextPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_313_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_313_ext
	return nil
}
func (ϟa *GlGetObjectLabelEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_314_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_314_ext
	return nil
}
func (ϟa *GlGetPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_315_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_315_ext
	return nil
}
func (ϟa *GlGetPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_316_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_316_ext
	return nil
}
func (ϟa *GlGetPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_317_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_317_ext
	return nil
}
func (ϟa *GlGetPathLengthNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_318_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_318_ext
	return nil
}
func (ϟa *GlGetPathMetricRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_319_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_319_ext
	return nil
}
func (ϟa *GlGetPathMetricsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_320_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_320_ext
	return nil
}
func (ϟa *GlGetPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_321_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_321_ext
	return nil
}
func (ϟa *GlGetPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_322_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_322_ext
	return nil
}
func (ϟa *GlGetPathSpacingNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_323_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_323_ext
	return nil
}
func (ϟa *GlGetPerfCounterInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_324_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_324_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterDataAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_325_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_325_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterInfoAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_326_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_326_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_327_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_327_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_328_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_328_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_329_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_329_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_330_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_330_ext
	return nil
}
func (ϟa *GlGetPerfQueryDataINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_331_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_331_ext
	return nil
}
func (ϟa *GlGetPerfQueryIdByNameINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_332_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_332_ext
	return nil
}
func (ϟa *GlGetPerfQueryInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_333_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_333_ext
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_334_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_334_ext, l
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLogEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_335_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_335_ext
	return nil
}
func (ϟa *GlGetProgramPipelineivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_336_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_336_ext
	return nil
}
func (ϟa *GlGetProgramResourcefvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_337_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_337_ext
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_338_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_338_ext
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_339_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_339_ext
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_340_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_340_ext
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_341_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_342_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_341_ext, requiresExtension_342_ext
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_343_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_344_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_343_ext, requiresExtension_344_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_345_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_345_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_346_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_346_ext
	return nil
}
func (ϟa *GlGetSyncivAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_347_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_347_ext
	return nil
}
func (ϟa *GlGetTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_348_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_348_ext
	return nil
}
func (ϟa *GlGetTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_349_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_349_ext
	return nil
}
func (ϟa *GlGetTextureHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_350_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_350_ext
	return nil
}
func (ϟa *GlGetTextureSamplerHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_351_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_351_ext
	return nil
}
func (ϟa *GlGetTranslatedShaderSourceANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_352_ext := ExtensionId_GL_ANGLE_translated_shader_source // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_352_ext
	return nil
}
func (ϟa *GlGetnUniformfvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_353_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_353_ext
	return nil
}
func (ϟa *GlGetnUniformfvKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_354_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_354_ext
	return nil
}
func (ϟa *GlGetnUniformivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_355_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_355_ext
	return nil
}
func (ϟa *GlGetnUniformivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_356_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_356_ext
	return nil
}
func (ϟa *GlGetnUniformuivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_357_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_357_ext
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_358_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_358_ext
	return nil
}
func (ϟa *GlInterpolatePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_359_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_359_ext
	return nil
}
func (ϟa *GlIsEnablediNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_360_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_360_ext
	return nil
}
func (ϟa *GlIsEnablediOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_361_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_361_ext
	return nil
}
func (ϟa *GlIsFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_362_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_362_ext
	return nil
}
func (ϟa *GlIsImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_363_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_363_ext
	return nil
}
func (ϟa *GlIsPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_364_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_364_ext
	return nil
}
func (ϟa *GlIsPointInFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_365_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_365_ext
	return nil
}
func (ϟa *GlIsPointInStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_366_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_366_ext
	return nil
}
func (ϟa *GlIsProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_367_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_367_ext
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_368_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_369_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_371_msg := "No context bound" // string
		return
		_ = error_371_msg
	}
	GetContext_370_result := context // Contextʳ
	ctx := GetContext_370_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = requiresExtension_368_ext, requiresExtension_369_ext, context, GetContext_370_result, ctx
	return nil
}
func (ϟa *GlIsSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_372_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_372_ext
	return nil
}
func (ϟa *GlIsTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_373_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_373_ext
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_374_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_376_msg := "No context bound" // string
		return
		_ = error_376_msg
	}
	GetContext_375_result := context // Contextʳ
	ctx := GetContext_375_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.VertexArrays.Contains(ϟa.Array) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _ = requiresExtension_374_ext, context, GetContext_375_result, ctx
	return nil
}
func (ϟa *GlLabelObjectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_377_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_377_ext
	return nil
}
func (ϟa *GlMakeImageHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_378_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_378_ext
	return nil
}
func (ϟa *GlMakeImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_379_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_379_ext
	return nil
}
func (ϟa *GlMakeTextureHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_380_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_380_ext
	return nil
}
func (ϟa *GlMakeTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_381_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_381_ext
	return nil
}
func (ϟa *GlMapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_382_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_382_ext
	return nil
}
func (ϟa *GlMapBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_383_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_383_ext
	return nil
}
func (ϟa *GlMatrixLoad3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_384_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_384_ext
	return nil
}
func (ϟa *GlMatrixLoad3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_385_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_385_ext
	return nil
}
func (ϟa *GlMatrixLoadTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_386_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_386_ext
	return nil
}
func (ϟa *GlMatrixMult3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_387_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_387_ext
	return nil
}
func (ϟa *GlMatrixMult3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_388_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_388_ext
	return nil
}
func (ϟa *GlMatrixMultTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_389_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_389_ext
	return nil
}
func (ϟa *GlMultiDrawArraysEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_390_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_390_ext
	return nil
}
func (ϟa *GlMultiDrawArraysIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_391_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_391_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_392_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_392_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_393_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_393_ext
	return nil
}
func (ϟa *GlMultiDrawElementsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_394_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_394_ext
	return nil
}
func (ϟa *GlMultiDrawElementsIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_395_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_395_ext
	return nil
}
func (ϟa *GlNamedFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_396_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_396_ext
	return nil
}
func (ϟa *GlPatchParameteriOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_397_ext := ExtensionId_GL_OES_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_397_ext
	return nil
}
func (ϟa *GlPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_398_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_398_ext
	return nil
}
func (ϟa *GlPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_399_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_399_ext
	return nil
}
func (ϟa *GlPathCoverDepthFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_400_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_400_ext
	return nil
}
func (ϟa *GlPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_401_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_401_ext
	return nil
}
func (ϟa *GlPathGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_402_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_402_ext
	return nil
}
func (ϟa *GlPathGlyphIndexRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_403_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_403_ext
	return nil
}
func (ϟa *GlPathGlyphRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_404_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_404_ext
	return nil
}
func (ϟa *GlPathGlyphsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_405_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_405_ext
	return nil
}
func (ϟa *GlPathMemoryGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_406_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_406_ext
	return nil
}
func (ϟa *GlPathParameterfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_407_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_407_ext
	return nil
}
func (ϟa *GlPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_408_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_408_ext
	return nil
}
func (ϟa *GlPathParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_409_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_409_ext
	return nil
}
func (ϟa *GlPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_410_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_410_ext
	return nil
}
func (ϟa *GlPathStencilDepthOffsetNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_411_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_411_ext
	return nil
}
func (ϟa *GlPathStencilFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_412_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_412_ext
	return nil
}
func (ϟa *GlPathStringNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_413_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_413_ext
	return nil
}
func (ϟa *GlPathSubCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_414_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_414_ext
	return nil
}
func (ϟa *GlPathSubCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_415_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_415_ext
	return nil
}
func (ϟa *GlPointAlongPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_416_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_416_ext
	return nil
}
func (ϟa *GlPolygonModeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_417_ext := ExtensionId_GL_NV_polygon_mode // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_417_ext
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_418_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_418_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_419_ext := ExtensionId_GL_OES_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_419_ext
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_420_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_420_ext
	return nil
}
func (ϟa *GlProgramParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_421_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_421_ext
	return nil
}
func (ϟa *GlProgramPathFragmentInputGenNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_422_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_422_ext
	return nil
}
func (ϟa *GlProgramUniform1fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_423_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_423_ext
	return nil
}
func (ϟa *GlProgramUniform1fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_424_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_424_ext
	return nil
}
func (ϟa *GlProgramUniform1iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_425_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_425_ext
	return nil
}
func (ϟa *GlProgramUniform1ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_426_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_426_ext
	return nil
}
func (ϟa *GlProgramUniform1uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_427_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_427_ext
	return nil
}
func (ϟa *GlProgramUniform1uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_428_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_428_ext
	return nil
}
func (ϟa *GlProgramUniform2fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_429_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_429_ext
	return nil
}
func (ϟa *GlProgramUniform2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_430_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_430_ext
	return nil
}
func (ϟa *GlProgramUniform2iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_431_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_431_ext
	return nil
}
func (ϟa *GlProgramUniform2ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_432_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_432_ext
	return nil
}
func (ϟa *GlProgramUniform2uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_433_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_433_ext
	return nil
}
func (ϟa *GlProgramUniform2uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_434_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_434_ext
	return nil
}
func (ϟa *GlProgramUniform3fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_435_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_435_ext
	return nil
}
func (ϟa *GlProgramUniform3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_436_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_436_ext
	return nil
}
func (ϟa *GlProgramUniform3iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_437_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_437_ext
	return nil
}
func (ϟa *GlProgramUniform3ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_438_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_438_ext
	return nil
}
func (ϟa *GlProgramUniform3uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_439_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_439_ext
	return nil
}
func (ϟa *GlProgramUniform3uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_440_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_440_ext
	return nil
}
func (ϟa *GlProgramUniform4fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_441_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_441_ext
	return nil
}
func (ϟa *GlProgramUniform4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_442_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_442_ext
	return nil
}
func (ϟa *GlProgramUniform4iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_443_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_443_ext
	return nil
}
func (ϟa *GlProgramUniform4ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_444_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_444_ext
	return nil
}
func (ϟa *GlProgramUniform4uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_445_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_445_ext
	return nil
}
func (ϟa *GlProgramUniform4uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_446_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_446_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_447_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_447_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_448_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_448_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_449_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_449_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_450_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_450_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_451_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_451_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_452_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_452_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_453_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_453_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_454_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_454_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_455_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_455_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_456_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_456_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_457_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_457_ext
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_458_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_458_ext
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_459_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_459_ext
	return nil
}
func (ϟa *GlRasterSamplesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_460_ext := ExtensionId_GL_EXT_raster_multisample       // ExtensionId
	requiresExtension_461_ext := ExtensionId_GL_EXT_texture_filter_minmax    // ExtensionId
	requiresExtension_462_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_460_ext, requiresExtension_461_ext, requiresExtension_462_ext
	return nil
}
func (ϟa *GlReadBufferIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_463_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_463_ext
	return nil
}
func (ϟa *GlReadBufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_464_ext := ExtensionId_GL_NV_read_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_464_ext
	return nil
}
func (ϟa *GlReadnPixelsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_465_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_465_ext
	return nil
}
func (ϟa *GlReadnPixelsKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_466_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_466_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_467_ext := ExtensionId_GL_ANGLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_467_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_468_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_468_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_469_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_469_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_470_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_470_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_471_ext := ExtensionId_GL_NV_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_471_ext
	return nil
}
func (ϟa *GlResolveDepthValuesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_472_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_472_ext
	return nil
}
func (ϟa *GlResolveMultisampleFramebufferAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_473_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_473_ext
	return nil
}
func (ϟa *GlSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_474_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_474_ext
	return nil
}
func (ϟa *GlSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_475_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_475_ext
	return nil
}
func (ϟa *GlScissorArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_476_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_476_ext
	return nil
}
func (ϟa *GlScissorIndexedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_477_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_477_ext
	return nil
}
func (ϟa *GlScissorIndexedvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_478_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_478_ext
	return nil
}
func (ϟa *GlSelectPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_479_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_479_ext
	return nil
}
func (ϟa *GlSetFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_480_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_480_ext
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_481_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_481_ext
	return nil
}
func (ϟa *GlStencilFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_482_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_482_ext
	return nil
}
func (ϟa *GlStencilFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_483_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_483_ext
	return nil
}
func (ϟa *GlStencilStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_484_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_484_ext
	return nil
}
func (ϟa *GlStencilStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_485_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_485_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_486_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_486_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_487_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_487_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_488_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_488_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_489_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_489_ext
	return nil
}
func (ϟa *GlSubpixelPrecisionBiasNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_490_ext := ExtensionId_GL_NV_conservative_raster // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_490_ext
	return nil
}
func (ϟa *GlTestFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_491_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_491_ext
	return nil
}
func (ϟa *GlTexBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_492_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_492_ext
	return nil
}
func (ϟa *GlTexBufferRangeOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_493_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_493_ext
	return nil
}
func (ϟa *GlTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_494_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_494_ext
	return nil
}
func (ϟa *GlTexPageCommitmentARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_495_ext := ExtensionId_GL_EXT_sparse_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_495_ext
	return nil
}
func (ϟa *GlTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_496_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_496_ext
	return nil
}
func (ϟa *GlTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_497_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_497_ext
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_498_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_498_ext
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_499_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_499_ext
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_500_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_500_ext
	return nil
}
func (ϟa *GlTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_501_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_501_ext
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_502_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_502_ext
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_503_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_503_ext
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_504_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_504_ext
	return nil
}
func (ϟa *GlTextureViewEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_505_ext := ExtensionId_GL_EXT_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_505_ext
	return nil
}
func (ϟa *GlTextureViewOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_506_ext := ExtensionId_GL_OES_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_506_ext
	return nil
}
func (ϟa *GlTransformPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_507_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_507_ext
	return nil
}
func (ϟa *GlUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_508_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_508_ext
	return nil
}
func (ϟa *GlUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_509_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_509_ext
	return nil
}
func (ϟa *GlUniformMatrix2x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_510_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_510_ext
	return nil
}
func (ϟa *GlUniformMatrix2x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_511_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_511_ext
	return nil
}
func (ϟa *GlUniformMatrix3x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_512_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_512_ext
	return nil
}
func (ϟa *GlUniformMatrix3x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_513_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_513_ext
	return nil
}
func (ϟa *GlUniformMatrix4x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_514_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_514_ext
	return nil
}
func (ϟa *GlUniformMatrix4x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_515_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_515_ext
	return nil
}
func (ϟa *GlUnmapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_516_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_516_ext
	return nil
}
func (ϟa *GlUseProgramStagesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_517_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_517_ext
	return nil
}
func (ϟa *GlValidateProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_518_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_518_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_519_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_519_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_520_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_520_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_521_ext := ExtensionId_GL_NV_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_521_ext
	return nil
}
func (ϟa *GlViewportArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_522_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_522_ext
	return nil
}
func (ϟa *GlViewportIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_523_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_523_ext
	return nil
}
func (ϟa *GlViewportIndexedfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_524_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_524_ext
	return nil
}
func (ϟa *GlWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_525_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_525_ext
	return nil
}
func (ϟa *GlWeightPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_526_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_526_ext
	return nil
}
func (ϟa *GlBlendBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_527_major := uint32(3) // u32
	minRequiredVersion_527_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_527_major, minRequiredVersion_527_minor
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_528_major := uint32(2)    // u32
	minRequiredVersion_528_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_530_msg := "No context bound" // string
		return
		_ = error_530_msg
	}
	GetContext_529_result := context // Contextʳ
	ctx := GetContext_529_result     // Contextʳ
	ctx.Blending.BlendColor = Color{Red: ϟa.Red, Green: ϟa.Green, Blue: ϟa.Blue, Alpha: ϟa.Alpha}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_528_major, minRequiredVersion_528_minor, context, GetContext_529_result, ctx
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_531_major := uint32(2) // u32
	minRequiredVersion_531_minor := uint32(0) // u32
	switch ϟa.Equation {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_532_major := uint32(3) // u32
		minRequiredVersion_532_minor := uint32(0) // u32
		_, _ = minRequiredVersion_532_major, minRequiredVersion_532_minor
	default:
		glErrorInvalidEnum_533_param := ϟa.Equation // GLenum
		return
		_ = glErrorInvalidEnum_533_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_535_msg := "No context bound" // string
		return
		_ = error_535_msg
	}
	GetContext_534_result := context // Contextʳ
	ctx := GetContext_534_result     // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_531_major, minRequiredVersion_531_minor, context, GetContext_534_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_536_major := uint32(2) // u32
	minRequiredVersion_536_minor := uint32(0) // u32
	switch ϟa.Rgb {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_537_major := uint32(3) // u32
		minRequiredVersion_537_minor := uint32(0) // u32
		_, _ = minRequiredVersion_537_major, minRequiredVersion_537_minor
	default:
		glErrorInvalidEnum_538_param := ϟa.Rgb // GLenum
		return
		_ = glErrorInvalidEnum_538_param
	}
	switch ϟa.Alpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_539_major := uint32(3) // u32
		minRequiredVersion_539_minor := uint32(0) // u32
		_, _ = minRequiredVersion_539_major, minRequiredVersion_539_minor
	default:
		glErrorInvalidEnum_540_param := ϟa.Alpha // GLenum
		return
		_ = glErrorInvalidEnum_540_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_542_msg := "No context bound" // string
		return
		_ = error_542_msg
	}
	GetContext_541_result := context // Contextʳ
	ctx := GetContext_541_result     // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_536_major, minRequiredVersion_536_minor, context, GetContext_541_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparatei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_543_major := uint32(3) // u32
	minRequiredVersion_543_minor := uint32(2) // u32
	switch ϟa.ModeRGB {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		glErrorInvalidEnum_544_param := ϟa.ModeRGB // GLenum
		return
		_ = glErrorInvalidEnum_544_param
	}
	switch ϟa.ModeAlpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		glErrorInvalidEnum_545_param := ϟa.ModeAlpha // GLenum
		return
		_ = glErrorInvalidEnum_545_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_543_major, minRequiredVersion_543_minor
	return nil
}
func (ϟa *GlBlendEquationi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_546_major := uint32(3) // u32
	minRequiredVersion_546_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		glErrorInvalidEnum_547_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_547_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_546_major, minRequiredVersion_546_minor
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_548_major := uint32(2) // u32
	minRequiredVersion_548_minor := uint32(0) // u32
	switch ϟa.SrcFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_549_param := ϟa.SrcFactor // GLenum
		return
		_ = glErrorInvalidEnum_549_param
	}
	switch ϟa.DstFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_550_param := ϟa.DstFactor // GLenum
		return
		_ = glErrorInvalidEnum_550_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_552_msg := "No context bound" // string
		return
		_ = error_552_msg
	}
	GetContext_551_result := context // Contextʳ
	ctx := GetContext_551_result     // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_548_major, minRequiredVersion_548_minor, context, GetContext_551_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_553_major := uint32(2) // u32
	minRequiredVersion_553_minor := uint32(0) // u32
	switch ϟa.SrcFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_554_param := ϟa.SrcFactorRgb // GLenum
		return
		_ = glErrorInvalidEnum_554_param
	}
	switch ϟa.DstFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_555_param := ϟa.DstFactorRgb // GLenum
		return
		_ = glErrorInvalidEnum_555_param
	}
	switch ϟa.SrcFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_556_param := ϟa.SrcFactorAlpha // GLenum
		return
		_ = glErrorInvalidEnum_556_param
	}
	switch ϟa.DstFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_557_param := ϟa.DstFactorAlpha // GLenum
		return
		_ = glErrorInvalidEnum_557_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_559_msg := "No context bound" // string
		return
		_ = error_559_msg
	}
	GetContext_558_result := context // Contextʳ
	ctx := GetContext_558_result     // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_553_major, minRequiredVersion_553_minor, context, GetContext_558_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparatei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_560_major := uint32(3) // u32
	minRequiredVersion_560_minor := uint32(2) // u32
	switch ϟa.SrcRGB {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_561_param := ϟa.SrcRGB // GLenum
		return
		_ = glErrorInvalidEnum_561_param
	}
	switch ϟa.DstRGB {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_562_param := ϟa.DstRGB // GLenum
		return
		_ = glErrorInvalidEnum_562_param
	}
	switch ϟa.SrcAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_563_param := ϟa.SrcAlpha // GLenum
		return
		_ = glErrorInvalidEnum_563_param
	}
	switch ϟa.DstAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_564_param := ϟa.DstAlpha // GLenum
		return
		_ = glErrorInvalidEnum_564_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_560_major, minRequiredVersion_560_minor
	return nil
}
func (ϟa *GlBlendFunci) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_565_major := uint32(3) // u32
	minRequiredVersion_565_minor := uint32(2) // u32
	switch ϟa.Src {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_566_param := ϟa.Src // GLenum
		return
		_ = glErrorInvalidEnum_566_param
	}
	switch ϟa.Dst {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_567_param := ϟa.Dst // GLenum
		return
		_ = glErrorInvalidEnum_567_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_565_major, minRequiredVersion_565_minor
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_568_major := uint32(2) // u32
	minRequiredVersion_568_minor := uint32(0) // u32
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		glErrorInvalidEnum_569_param := ϟa.Function // GLenum
		return
		_ = glErrorInvalidEnum_569_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_571_msg := "No context bound" // string
		return
		_ = error_571_msg
	}
	GetContext_570_result := context // Contextʳ
	ctx := GetContext_570_result     // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_568_major, minRequiredVersion_568_minor, context, GetContext_570_result, ctx
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_572_major := uint32(2)    // u32
	minRequiredVersion_572_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_574_msg := "No context bound" // string
		return
		_ = error_574_msg
	}
	GetContext_573_result := context // Contextʳ
	ctx := GetContext_573_result     // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_572_major, minRequiredVersion_572_minor, context, GetContext_573_result, ctx
	return nil
}
func (ϟa *GlSampleMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_575_major := uint32(3) // u32
	minRequiredVersion_575_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_575_major, minRequiredVersion_575_minor
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_576_major := uint32(2)    // u32
	minRequiredVersion_576_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_578_msg := "No context bound" // string
		return
		_ = error_578_msg
	}
	GetContext_577_result := context // Contextʳ
	ctx := GetContext_577_result     // Contextʳ
	ctx.Rasterizing.Scissor = Rect{X: ϟa.X, Y: ϟa.Y, Width: ϟa.Width, Height: ϟa.Height}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_576_major, minRequiredVersion_576_minor, context, GetContext_577_result, ctx
	return nil
}
func (ϟa *GlStencilFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_579_major := uint32(2) // u32
	minRequiredVersion_579_minor := uint32(0) // u32
	switch ϟa.Func {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		glErrorInvalidEnum_580_param := ϟa.Func // GLenum
		return
		_ = glErrorInvalidEnum_580_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_579_major, minRequiredVersion_579_minor
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_581_major := uint32(2) // u32
	minRequiredVersion_581_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_582_param := ϟa.Face // GLenum
		return
		_ = glErrorInvalidEnum_582_param
	}
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		glErrorInvalidEnum_583_param := ϟa.Function // GLenum
		return
		_ = glErrorInvalidEnum_583_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_581_major, minRequiredVersion_581_minor
	return nil
}
func (ϟa *GlStencilOp) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_584_major := uint32(2) // u32
	minRequiredVersion_584_minor := uint32(0) // u32
	switch ϟa.Fail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_585_param := ϟa.Fail // GLenum
		return
		_ = glErrorInvalidEnum_585_param
	}
	switch ϟa.Zfail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_586_param := ϟa.Zfail // GLenum
		return
		_ = glErrorInvalidEnum_586_param
	}
	switch ϟa.Zpass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_587_param := ϟa.Zpass // GLenum
		return
		_ = glErrorInvalidEnum_587_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_584_major, minRequiredVersion_584_minor
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_588_major := uint32(2) // u32
	minRequiredVersion_588_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_589_param := ϟa.Face // GLenum
		return
		_ = glErrorInvalidEnum_589_param
	}
	switch ϟa.StencilFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_590_param := ϟa.StencilFail // GLenum
		return
		_ = glErrorInvalidEnum_590_param
	}
	switch ϟa.StencilPassDepthFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_591_param := ϟa.StencilPassDepthFail // GLenum
		return
		_ = glErrorInvalidEnum_591_param
	}
	switch ϟa.StencilPassDepthPass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_592_param := ϟa.StencilPassDepthPass // GLenum
		return
		_ = glErrorInvalidEnum_592_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_588_major, minRequiredVersion_588_minor
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_593_major := uint32(2) // u32
	minRequiredVersion_593_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_594_major := uint32(3) // u32
		minRequiredVersion_594_minor := uint32(0) // u32
		_, _ = minRequiredVersion_594_major, minRequiredVersion_594_minor
	default:
		glErrorInvalidEnum_595_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_595_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_597_msg := "No context bound" // string
		return
		_ = error_597_msg
	}
	GetContext_596_result := context // Contextʳ
	ctx := GetContext_596_result     // Contextʳ
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
	_, _, _, _, _ = minRequiredVersion_593_major, minRequiredVersion_593_minor, context, GetContext_596_result, ctx
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_598_major := uint32(2) // u32
	minRequiredVersion_598_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_599_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_599_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_601_msg := "No context bound" // string
		return
		_ = error_601_msg
	}
	GetContext_600_result := context // Contextʳ
	ctx := GetContext_600_result     // Contextʳ
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = &Renderbuffer{}
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_598_major, minRequiredVersion_598_minor, context, GetContext_600_result, ctx
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_602_major := uint32(3)                                                                                                // u32
	minRequiredVersion_602_minor := uint32(0)                                                                                                // u32
	supportsBits_603_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_603_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	switch ϟa.Filter {
	case GLenum_GL_LINEAR, GLenum_GL_NEAREST:
	default:
		glErrorInvalidEnum_604_param := ϟa.Filter // GLenum
		return
		_ = glErrorInvalidEnum_604_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_602_major, minRequiredVersion_602_minor, supportsBits_603_seenBits, supportsBits_603_validBits
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_605_major := uint32(2) // u32
	minRequiredVersion_605_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_606_major := uint32(3) // u32
		minRequiredVersion_606_minor := uint32(0) // u32
		_, _ = minRequiredVersion_606_major, minRequiredVersion_606_minor
	default:
		glErrorInvalidEnum_607_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_607_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_605_major, minRequiredVersion_605_minor
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_608_major := uint32(2)                                                                                                // u32
	minRequiredVersion_608_minor := uint32(0)                                                                                                // u32
	supportsBits_609_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_609_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_608_major, minRequiredVersion_608_minor, supportsBits_609_seenBits, supportsBits_609_validBits
	return nil
}
func (ϟa *GlClearBufferfi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_610_major := uint32(3) // u32
	minRequiredVersion_610_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_DEPTH_STENCIL:
	default:
		glErrorInvalidEnum_611_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_611_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_610_major, minRequiredVersion_610_minor
	return nil
}
func (ϟa *GlClearBufferfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_612_major := uint32(3) // u32
	minRequiredVersion_612_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_DEPTH:
	default:
		glErrorInvalidEnum_613_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_613_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_612_major, minRequiredVersion_612_minor
	return nil
}
func (ϟa *GlClearBufferiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_614_major := uint32(3) // u32
	minRequiredVersion_614_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_STENCIL:
	default:
		glErrorInvalidEnum_615_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_615_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_614_major, minRequiredVersion_614_minor
	return nil
}
func (ϟa *GlClearBufferuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_616_major := uint32(3) // u32
	minRequiredVersion_616_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
	default:
		glErrorInvalidEnum_617_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_617_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_616_major, minRequiredVersion_616_minor
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_618_major := uint32(2)    // u32
	minRequiredVersion_618_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_620_msg := "No context bound" // string
		return
		_ = error_620_msg
	}
	GetContext_619_result := context // Contextʳ
	ctx := GetContext_619_result     // Contextʳ
	ctx.Clearing.ClearColor = Color{Red: ϟa.R, Green: ϟa.G, Blue: ϟa.B, Alpha: ϟa.A}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_618_major, minRequiredVersion_618_minor, context, GetContext_619_result, ctx
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_621_major := uint32(2)    // u32
	minRequiredVersion_621_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_623_msg := "No context bound" // string
		return
		_ = error_623_msg
	}
	GetContext_622_result := context // Contextʳ
	ctx := GetContext_622_result     // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_621_major, minRequiredVersion_621_minor, context, GetContext_622_result, ctx
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_624_major := uint32(2)    // u32
	minRequiredVersion_624_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_626_msg := "No context bound" // string
		return
		_ = error_626_msg
	}
	GetContext_625_result := context // Contextʳ
	ctx := GetContext_625_result     // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_624_major, minRequiredVersion_624_minor, context, GetContext_625_result, ctx
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_627_major := uint32(2)    // u32
	minRequiredVersion_627_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_629_msg := "No context bound" // string
		return
		_ = error_629_msg
	}
	GetContext_628_result := context // Contextʳ
	ctx := GetContext_628_result     // Contextʳ
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_627_major, minRequiredVersion_627_minor, context, GetContext_628_result, ctx
	return nil
}
func (ϟa *GlColorMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_630_major := uint32(3) // u32
	minRequiredVersion_630_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_630_major, minRequiredVersion_630_minor
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_631_major := uint32(2)                                   // u32
	minRequiredVersion_631_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_633_msg := "No context bound" // string
		return
		_ = error_633_msg
	}
	GetContext_632_result := context // Contextʳ
	ctx := GetContext_632_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Framebuffers, f.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_631_major, minRequiredVersion_631_minor, f, context, GetContext_632_result, ctx
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_634_major := uint32(2)                                    // u32
	minRequiredVersion_634_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_636_msg := "No context bound" // string
		return
		_ = error_636_msg
	}
	GetContext_635_result := context // Contextʳ
	ctx := GetContext_635_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Renderbuffers, r.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_634_major, minRequiredVersion_634_minor, r, context, GetContext_635_result, ctx
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_637_major := uint32(2)    // u32
	minRequiredVersion_637_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_639_msg := "No context bound" // string
		return
		_ = error_639_msg
	}
	GetContext_638_result := context // Contextʳ
	ctx := GetContext_638_result     // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_637_major, minRequiredVersion_637_minor, context, GetContext_638_result, ctx
	return nil
}
func (ϟa *GlFramebufferParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_640_major := uint32(3) // u32
	minRequiredVersion_640_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_641_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_641_param
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	case GLenum_GL_FRAMEBUFFER_DEFAULT_LAYERS:
		minRequiredVersion_642_major := uint32(3) // u32
		minRequiredVersion_642_minor := uint32(2) // u32
		_, _ = minRequiredVersion_642_major, minRequiredVersion_642_minor
	default:
		glErrorInvalidEnum_643_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_643_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_640_major, minRequiredVersion_640_minor
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_644_major := uint32(2) // u32
	minRequiredVersion_644_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_645_major := uint32(3) // u32
		minRequiredVersion_645_minor := uint32(0) // u32
		_, _ = minRequiredVersion_645_major, minRequiredVersion_645_minor
	default:
		glErrorInvalidEnum_646_param := ϟa.FramebufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_646_param
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_647_major := uint32(3) // u32
		minRequiredVersion_647_minor := uint32(0) // u32
		_, _ = minRequiredVersion_647_major, minRequiredVersion_647_minor
	default:
		glErrorInvalidEnum_648_param := ϟa.FramebufferAttachment // GLenum
		return
		_ = glErrorInvalidEnum_648_param
	}
	switch ϟa.RenderbufferTarget {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_649_param := ϟa.RenderbufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_649_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_651_msg := "No context bound" // string
		return
		_ = error_651_msg
	}
	GetContext_650_result := context // Contextʳ
	ctx := GetContext_650_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_644_major, minRequiredVersion_644_minor, context, GetContext_650_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_652_major := uint32(3) // u32
	minRequiredVersion_652_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_653_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_653_param
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	default:
		glErrorInvalidEnum_654_param := ϟa.Attachment // GLenum
		return
		_ = glErrorInvalidEnum_654_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_652_major, minRequiredVersion_652_minor
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_655_major := uint32(2) // u32
	minRequiredVersion_655_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_656_major := uint32(3) // u32
		minRequiredVersion_656_minor := uint32(0) // u32
		_, _ = minRequiredVersion_656_major, minRequiredVersion_656_minor
	default:
		glErrorInvalidEnum_657_param := ϟa.FramebufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_657_param
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_658_major := uint32(3) // u32
		minRequiredVersion_658_minor := uint32(0) // u32
		_, _ = minRequiredVersion_658_major, minRequiredVersion_658_minor
	default:
		glErrorInvalidEnum_659_param := ϟa.FramebufferAttachment // GLenum
		return
		_ = glErrorInvalidEnum_659_param
	}
	switch ϟa.TextureTarget {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_660_major := uint32(3) // u32
		minRequiredVersion_660_minor := uint32(1) // u32
		_, _ = minRequiredVersion_660_major, minRequiredVersion_660_minor
	default:
		glErrorInvalidEnum_661_param := ϟa.TextureTarget // GLenum
		return
		_ = glErrorInvalidEnum_661_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_663_msg := "No context bound" // string
		return
		_ = error_663_msg
	}
	GetContext_662_result := context // Contextʳ
	ctx := GetContext_662_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_655_major, minRequiredVersion_655_minor, context, GetContext_662_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTextureLayer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_664_major := uint32(3) // u32
	minRequiredVersion_664_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_665_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_665_param
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	default:
		glErrorInvalidEnum_666_param := ϟa.Attachment // GLenum
		return
		_ = glErrorInvalidEnum_666_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_664_major, minRequiredVersion_664_minor
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_667_major := uint32(2)                                   // u32
	minRequiredVersion_667_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_669_msg := "No context bound" // string
		return
		_ = error_669_msg
	}
	GetContext_668_result := context // Contextʳ
	ctx := GetContext_668_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // FramebufferId
		ctx.Instances.Framebuffers[id] = &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}}
		f.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_667_major, minRequiredVersion_667_minor, f, context, GetContext_668_result, ctx
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_670_major := uint32(2)                                    // u32
	minRequiredVersion_670_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_672_msg := "No context bound" // string
		return
		_ = error_672_msg
	}
	GetContext_671_result := context // Contextʳ
	ctx := GetContext_671_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = &Renderbuffer{}
		r.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_670_major, minRequiredVersion_670_minor, r, context, GetContext_671_result, ctx
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_673_major := uint32(2) // u32
	minRequiredVersion_673_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_674_major := uint32(3) // u32
		minRequiredVersion_674_minor := uint32(0) // u32
		_, _ = minRequiredVersion_674_major, minRequiredVersion_674_minor
	default:
		glErrorInvalidEnum_675_param := ϟa.FramebufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_675_param
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL:
		minRequiredVersion_676_major := uint32(3) // u32
		minRequiredVersion_676_minor := uint32(0) // u32
		_, _ = minRequiredVersion_676_major, minRequiredVersion_676_minor
	default:
		glErrorInvalidEnum_677_param := ϟa.Attachment // GLenum
		return
		_ = glErrorInvalidEnum_677_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME, GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		minRequiredVersion_678_major := uint32(3) // u32
		minRequiredVersion_678_minor := uint32(0) // u32
		_, _ = minRequiredVersion_678_major, minRequiredVersion_678_minor
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_LAYERED:
		minRequiredVersion_679_major := uint32(3) // u32
		minRequiredVersion_679_minor := uint32(2) // u32
		_, _ = minRequiredVersion_679_major, minRequiredVersion_679_minor
	default:
		glErrorInvalidEnum_680_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_680_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_682_msg := "No context bound" // string
		return
		_ = error_682_msg
	}
	GetContext_681_result := context // Contextʳ
	ctx := GetContext_681_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_673_major, minRequiredVersion_673_minor, context, GetContext_681_result, ctx, target, framebufferId, framebuffer, a
	return nil
}
func (ϟa *GlGetFramebufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_683_major := uint32(3) // u32
	minRequiredVersion_683_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_684_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_684_param
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	case GLenum_GL_FRAMEBUFFER_DEFAULT_LAYERS:
		minRequiredVersion_685_major := uint32(3) // u32
		minRequiredVersion_685_minor := uint32(2) // u32
		_, _ = minRequiredVersion_685_major, minRequiredVersion_685_minor
	default:
		glErrorInvalidEnum_686_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_686_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_683_major, minRequiredVersion_683_minor
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_687_major := uint32(2) // u32
	minRequiredVersion_687_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_688_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_688_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_RENDERBUFFER_ALPHA_SIZE, GLenum_GL_RENDERBUFFER_BLUE_SIZE, GLenum_GL_RENDERBUFFER_DEPTH_SIZE, GLenum_GL_RENDERBUFFER_GREEN_SIZE, GLenum_GL_RENDERBUFFER_HEIGHT, GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT, GLenum_GL_RENDERBUFFER_RED_SIZE, GLenum_GL_RENDERBUFFER_STENCIL_SIZE, GLenum_GL_RENDERBUFFER_WIDTH:
	case GLenum_GL_RENDERBUFFER_SAMPLES:
		minRequiredVersion_689_major := uint32(3) // u32
		minRequiredVersion_689_minor := uint32(0) // u32
		_, _ = minRequiredVersion_689_major, minRequiredVersion_689_minor
	default:
		glErrorInvalidEnum_690_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_690_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_692_msg := "No context bound" // string
		return
		_ = error_692_msg
	}
	GetContext_691_result := context            // Contextʳ
	ctx := GetContext_691_result                // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target) // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)   // Renderbufferʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_687_major, minRequiredVersion_687_minor, context, GetContext_691_result, ctx, id, rb
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_693_major := uint32(3) // u32
	minRequiredVersion_693_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_694_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_694_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_693_major, minRequiredVersion_693_minor
	return nil
}
func (ϟa *GlInvalidateSubFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_695_major := uint32(3) // u32
	minRequiredVersion_695_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_696_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_696_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_695_major, minRequiredVersion_695_minor
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_697_major := uint32(2)    // u32
	minRequiredVersion_697_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_699_msg := "No context bound" // string
		return
		_ = error_699_msg
	}
	GetContext_698_result := context // Contextʳ
	ctx := GetContext_698_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_697_major, minRequiredVersion_697_minor, context, GetContext_698_result, ctx
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_700_major := uint32(2)    // u32
	minRequiredVersion_700_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_702_msg := "No context bound" // string
		return
		_ = error_702_msg
	}
	GetContext_701_result := context // Contextʳ
	ctx := GetContext_701_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_700_major, minRequiredVersion_700_minor, context, GetContext_701_result, ctx
	return nil
}
func (ϟa *GlReadBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_703_major := uint32(3) // u32
	minRequiredVersion_703_minor := uint32(0) // u32
	switch ϟa.Src {
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_NONE:
	default:
		glErrorInvalidEnum_704_param := ϟa.Src // GLenum
		return
		_ = glErrorInvalidEnum_704_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_703_major, minRequiredVersion_703_minor
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_705_major := uint32(2) // u32
	minRequiredVersion_705_minor := uint32(0) // u32
	switch ϟa.Format {
	case GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER:
	default:
		glErrorInvalidEnum_706_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_706_param
	}
	switch ϟa.Type {
	case GLenum_GL_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT:
	case GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_707_major := uint32(3) // u32
		minRequiredVersion_707_minor := uint32(1) // u32
		_, _ = minRequiredVersion_707_major, minRequiredVersion_707_minor
	default:
		glErrorInvalidEnum_708_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_708_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Data.Slice(uint64(uint32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_705_major, minRequiredVersion_705_minor
	return nil
}
func (ϟa *GlReadnPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_709_major := uint32(3) // u32
	minRequiredVersion_709_minor := uint32(2) // u32
	switch ϟa.Format {
	case GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER:
	default:
		glErrorInvalidEnum_710_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_710_param
	}
	switch ϟa.Type {
	case GLenum_GL_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
	default:
		glErrorInvalidEnum_711_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_711_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_709_major, minRequiredVersion_709_minor
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_712_major := uint32(2) // u32
	minRequiredVersion_712_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_713_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_713_param
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGBA4, GLenum_GL_STENCIL_INDEX8:
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_714_major := uint32(3) // u32
		minRequiredVersion_714_minor := uint32(0) // u32
		_, _ = minRequiredVersion_714_major, minRequiredVersion_714_minor
	default:
		glErrorInvalidEnum_715_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_715_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_717_msg := "No context bound" // string
		return
		_ = error_717_msg
	}
	GetContext_716_result := context            // Contextʳ
	ctx := GetContext_716_result                // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target) // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)   // Renderbufferʳ
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_712_major, minRequiredVersion_712_minor, context, GetContext_716_result, ctx, id, rb
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_718_major := uint32(3) // u32
	minRequiredVersion_718_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_719_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_719_param
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		glErrorInvalidEnum_720_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_720_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_718_major, minRequiredVersion_718_minor
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_721_major := uint32(2)    // u32
	minRequiredVersion_721_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_723_msg := "No context bound" // string
		return
		_ = error_723_msg
	}
	GetContext_722_result := context // Contextʳ
	ctx := GetContext_722_result     // Contextʳ
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_721_major, minRequiredVersion_721_minor, context, GetContext_722_result, ctx
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_724_major := uint32(2) // u32
	minRequiredVersion_724_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_725_param := ϟa.Face // GLenum
		return
		_ = glErrorInvalidEnum_725_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_727_msg := "No context bound" // string
		return
		_ = error_727_msg
	}
	GetContext_726_result := context // Contextʳ
	ctx := GetContext_726_result     // Contextʳ
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
	_, _, _, _, _ = minRequiredVersion_724_major, minRequiredVersion_724_minor, context, GetContext_726_result, ctx
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_728_major := uint32(2) // u32
	minRequiredVersion_728_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_729_major := uint32(3) // u32
		minRequiredVersion_729_minor := uint32(0) // u32
		_, _ = minRequiredVersion_729_major, minRequiredVersion_729_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_730_major := uint32(3) // u32
		minRequiredVersion_730_minor := uint32(1) // u32
		_, _ = minRequiredVersion_730_major, minRequiredVersion_730_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS:
		minRequiredVersion_731_major := uint32(3) // u32
		minRequiredVersion_731_minor := uint32(2) // u32
		_, _ = minRequiredVersion_731_major, minRequiredVersion_731_minor
	default:
		glErrorInvalidEnum_732_param := ϟa.Capability // GLenum
		return
		_ = glErrorInvalidEnum_732_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_734_msg := "No context bound" // string
		return
		_ = error_734_msg
	}
	GetContext_733_result := context // Contextʳ
	ctx := GetContext_733_result     // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_728_major, minRequiredVersion_728_minor, context, GetContext_733_result, ctx
	return nil
}
func (ϟa *GlDisablei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_735_major := uint32(3) // u32
	minRequiredVersion_735_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		glErrorInvalidEnum_736_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_736_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_735_major, minRequiredVersion_735_minor
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_737_major := uint32(2) // u32
	minRequiredVersion_737_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_738_major := uint32(3) // u32
		minRequiredVersion_738_minor := uint32(0) // u32
		_, _ = minRequiredVersion_738_major, minRequiredVersion_738_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_739_major := uint32(3) // u32
		minRequiredVersion_739_minor := uint32(1) // u32
		_, _ = minRequiredVersion_739_major, minRequiredVersion_739_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS:
		minRequiredVersion_740_major := uint32(3) // u32
		minRequiredVersion_740_minor := uint32(2) // u32
		_, _ = minRequiredVersion_740_major, minRequiredVersion_740_minor
	default:
		glErrorInvalidEnum_741_param := ϟa.Capability // GLenum
		return
		_ = glErrorInvalidEnum_741_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_743_msg := "No context bound" // string
		return
		_ = error_743_msg
	}
	GetContext_742_result := context // Contextʳ
	ctx := GetContext_742_result     // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_737_major, minRequiredVersion_737_minor, context, GetContext_742_result, ctx
	return nil
}
func (ϟa *GlEnablei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_744_major := uint32(3) // u32
	minRequiredVersion_744_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		glErrorInvalidEnum_745_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_745_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_744_major, minRequiredVersion_744_minor
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_746_major := uint32(2) // u32
	minRequiredVersion_746_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_746_major, minRequiredVersion_746_minor
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_747_major := uint32(2) // u32
	minRequiredVersion_747_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_747_major, minRequiredVersion_747_minor
	return nil
}
func (ϟa *GlFlushMappedBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_748_major := uint32(3) // u32
	minRequiredVersion_748_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_749_major := uint32(3) // u32
		minRequiredVersion_749_minor := uint32(2) // u32
		_, _ = minRequiredVersion_749_major, minRequiredVersion_749_minor
	default:
		glErrorInvalidEnum_750_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_750_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_748_major, minRequiredVersion_748_minor
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_751_major := uint32(2) // u32
	minRequiredVersion_751_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_751_major, minRequiredVersion_751_minor
	return nil
}
func (ϟa *GlGetGraphicsResetStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_752_major := uint32(3) // u32
	minRequiredVersion_752_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_752_major, minRequiredVersion_752_minor
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_753_major := uint32(2) // u32
	minRequiredVersion_753_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_GENERATE_MIPMAP_HINT:
	case GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT:
		minRequiredVersion_754_major := uint32(3) // u32
		minRequiredVersion_754_minor := uint32(0) // u32
		_, _ = minRequiredVersion_754_major, minRequiredVersion_754_minor
	default:
		glErrorInvalidEnum_755_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_755_param
	}
	switch ϟa.Mode {
	case GLenum_GL_DONT_CARE, GLenum_GL_FASTEST, GLenum_GL_NICEST:
	default:
		glErrorInvalidEnum_756_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_756_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_758_msg := "No context bound" // string
		return
		_ = error_758_msg
	}
	GetContext_757_result := context // Contextʳ
	ctx := GetContext_757_result     // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_753_major, minRequiredVersion_753_minor, context, GetContext_757_result, ctx
	return nil
}
func (ϟa *GlActiveShaderProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_759_major := uint32(3) // u32
	minRequiredVersion_759_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_759_major, minRequiredVersion_759_minor
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_760_major := uint32(2)    // u32
	minRequiredVersion_760_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_762_msg := "No context bound" // string
		return
		_ = error_762_msg
	}
	GetContext_761_result := context            // Contextʳ
	ctx := GetContext_761_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)   // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_760_major, minRequiredVersion_760_minor, context, GetContext_761_result, ctx, p, s
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_763_major := uint32(2)    // u32
	minRequiredVersion_763_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_765_msg := "No context bound" // string
		return
		_ = error_765_msg
	}
	GetContext_764_result := context            // Contextʳ
	ctx := GetContext_764_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_763_major, minRequiredVersion_763_minor, context, GetContext_764_result, ctx, p
	return nil
}
func (ϟa *GlBindProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_766_major := uint32(3) // u32
	minRequiredVersion_766_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_766_major, minRequiredVersion_766_minor
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_767_major := uint32(2) // u32
	minRequiredVersion_767_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_767_major, minRequiredVersion_767_minor
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_768_major := uint32(2)    // u32
	minRequiredVersion_768_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_770_msg := "No context bound" // string
		return
		_ = error_770_msg
	}
	GetContext_769_result := context // Contextʳ
	ctx := GetContext_769_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = &Program{Shaders: GLenumːShaderIdᵐ{}, AttributeBindings: StringːAttributeLocationᵐ{}, Attributes: S32ːVertexAttributeᵐ{}, Uniforms: UniformLocationːUniformᵐ{}}
	ϟa.Result = id
	_, _, _, _, _, _ = minRequiredVersion_768_major, minRequiredVersion_768_minor, context, GetContext_769_result, ctx, id
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_771_major := uint32(2) // u32
	minRequiredVersion_771_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_COMPUTE_SHADER:
		minRequiredVersion_772_major := uint32(3) // u32
		minRequiredVersion_772_minor := uint32(1) // u32
		_, _ = minRequiredVersion_772_major, minRequiredVersion_772_minor
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_773_major := uint32(3) // u32
		minRequiredVersion_773_minor := uint32(2) // u32
		_, _ = minRequiredVersion_773_major, minRequiredVersion_773_minor
	default:
		glErrorInvalidEnum_774_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_774_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_776_msg := "No context bound" // string
		return
		_ = error_776_msg
	}
	GetContext_775_result := context // Contextʳ
	ctx := GetContext_775_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = &Shader{Compiled: false, Deletable: false}
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _, _, _, _, _, _ = minRequiredVersion_771_major, minRequiredVersion_771_minor, context, GetContext_775_result, ctx, id, s
	return nil
}
func (ϟa *GlCreateShaderProgramv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_777_major := uint32(3) // u32
	minRequiredVersion_777_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_778_major := uint32(3) // u32
		minRequiredVersion_778_minor := uint32(2) // u32
		_, _ = minRequiredVersion_778_major, minRequiredVersion_778_minor
	default:
		glErrorInvalidEnum_779_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_779_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_777_major, minRequiredVersion_777_minor
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_780_major := uint32(2)    // u32
	minRequiredVersion_780_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_782_msg := "No context bound" // string
		return
		_ = error_782_msg
	}
	GetContext_781_result := context // Contextʳ
	ctx := GetContext_781_result     // Contextʳ
	delete(ctx.Instances.Programs, ϟa.Program)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_780_major, minRequiredVersion_780_minor, context, GetContext_781_result, ctx
	return nil
}
func (ϟa *GlDeleteProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_783_major := uint32(3) // u32
	minRequiredVersion_783_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_783_major, minRequiredVersion_783_minor
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_784_major := uint32(2)    // u32
	minRequiredVersion_784_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_786_msg := "No context bound" // string
		return
		_ = error_786_msg
	}
	GetContext_785_result := context          // Contextʳ
	ctx := GetContext_785_result              // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
	s.Deletable = true
	delete(ctx.Instances.Shaders, ϟa.Shader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_784_major, minRequiredVersion_784_minor, context, GetContext_785_result, ctx, s
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_787_major := uint32(2)    // u32
	minRequiredVersion_787_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_789_msg := "No context bound" // string
		return
		_ = error_789_msg
	}
	GetContext_788_result := context            // Contextʳ
	ctx := GetContext_788_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)   // Shaderʳ
	delete(p.Shaders, s.Type)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_787_major, minRequiredVersion_787_minor, context, GetContext_788_result, ctx, p, s
	return nil
}
func (ϟa *GlDispatchCompute) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_790_major := uint32(3) // u32
	minRequiredVersion_790_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_790_major, minRequiredVersion_790_minor
	return nil
}
func (ϟa *GlDispatchComputeIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_791_major := uint32(3) // u32
	minRequiredVersion_791_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_791_major, minRequiredVersion_791_minor
	return nil
}
func (ϟa *GlGenProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_792_major := uint32(3) // u32
	minRequiredVersion_792_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_792_major, minRequiredVersion_792_minor
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_793_major := uint32(2) // u32
	minRequiredVersion_793_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_793_major, minRequiredVersion_793_minor
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_794_major := uint32(2) // u32
	minRequiredVersion_794_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_794_major, minRequiredVersion_794_minor
	return nil
}
func (ϟa *GlGetActiveUniformBlockName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_795_major := uint32(3) // u32
	minRequiredVersion_795_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _ = minRequiredVersion_795_major, minRequiredVersion_795_minor, l
	return nil
}
func (ϟa *GlGetActiveUniformBlockiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_796_major := uint32(3) // u32
	minRequiredVersion_796_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS, GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES, GLenum_GL_UNIFORM_BLOCK_BINDING, GLenum_GL_UNIFORM_BLOCK_DATA_SIZE, GLenum_GL_UNIFORM_BLOCK_NAME_LENGTH, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
	default:
		glErrorInvalidEnum_797_param := ϟa.ParameterName // GLenum
		return
		_ = glErrorInvalidEnum_797_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_796_major, minRequiredVersion_796_minor
	return nil
}
func (ϟa *GlGetActiveUniformsiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_798_major := uint32(3) // u32
	minRequiredVersion_798_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_ARRAY_STRIDE, GLenum_GL_UNIFORM_BLOCK_INDEX, GLenum_GL_UNIFORM_IS_ROW_MAJOR, GLenum_GL_UNIFORM_MATRIX_STRIDE, GLenum_GL_UNIFORM_NAME_LENGTH, GLenum_GL_UNIFORM_OFFSET, GLenum_GL_UNIFORM_SIZE, GLenum_GL_UNIFORM_TYPE:
	default:
		glErrorInvalidEnum_799_param := ϟa.ParameterName // GLenum
		return
		_ = glErrorInvalidEnum_799_param
	}
	ϟa.UniformIndices.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_798_major, minRequiredVersion_798_minor
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_800_major := uint32(2)    // u32
	minRequiredVersion_800_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_802_msg := "No context bound" // string
		return
		_ = error_802_msg
	}
	GetContext_801_result := context            // Contextʳ
	ctx := GetContext_801_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	min_803_a := int32(ϟa.BufferLength)         // s32
	min_803_b := int32(len(p.Shaders))          // s32
	min_803_result := func() (result int32) {
		switch (min_803_a) < (min_803_b) {
		case true:
			return min_803_a
		case false:
			return min_803_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_803_a) < (min_803_b), ϟa))
			return result
		}
	}() // s32
	l := min_803_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_800_major, minRequiredVersion_800_minor, context, GetContext_801_result, ctx, p, min_803_a, min_803_b, min_803_result, l
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_804_major := uint32(2) // u32
	minRequiredVersion_804_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_804_major, minRequiredVersion_804_minor
	return nil
}
func (ϟa *GlGetFragDataLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_805_major := uint32(3) // u32
	minRequiredVersion_805_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_805_major, minRequiredVersion_805_minor
	return nil
}
func (ϟa *GlGetProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_806_major := uint32(3) // u32
	minRequiredVersion_806_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_806_major, minRequiredVersion_806_minor
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_807_major := uint32(2)    // u32
	minRequiredVersion_807_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_809_msg := "No context bound" // string
		return
		_ = error_809_msg
	}
	GetContext_808_result := context            // Contextʳ
	ctx := GetContext_808_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	min_810_a := int32(ϟa.BufferLength)         // s32
	min_810_b := int32(p.InfoLog.Count)         // s32
	min_810_result := func() (result int32) {
		switch (min_810_a) < (min_810_b) {
		case true:
			return min_810_a
		case false:
			return min_810_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_810_a) < (min_810_b), ϟa))
			return result
		}
	}() // s32
	l := min_810_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(p.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_807_major, minRequiredVersion_807_minor, context, GetContext_808_result, ctx, p, min_810_a, min_810_b, min_810_result, l
	return nil
}
func (ϟa *GlGetProgramInterfaceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_811_major := uint32(3) // u32
	minRequiredVersion_811_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_812_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_812_param
	}
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_RESOURCES, GLenum_GL_MAX_NAME_LENGTH, GLenum_GL_MAX_NUM_ACTIVE_VARIABLES:
	default:
		glErrorInvalidEnum_813_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_813_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_811_major, minRequiredVersion_811_minor
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_814_major := uint32(3) // u32
	minRequiredVersion_814_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_814_major, minRequiredVersion_814_minor
	return nil
}
func (ϟa *GlGetProgramPipelineiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_815_major := uint32(3) // u32
	minRequiredVersion_815_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_PROGRAM, GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_VALIDATE_STATUS, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_816_major := uint32(3) // u32
		minRequiredVersion_816_minor := uint32(2) // u32
		_, _ = minRequiredVersion_816_major, minRequiredVersion_816_minor
	default:
		glErrorInvalidEnum_817_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_817_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_815_major, minRequiredVersion_815_minor
	return nil
}
func (ϟa *GlGetProgramResourceIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_818_major := uint32(3) // u32
	minRequiredVersion_818_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_819_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_819_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_818_major, minRequiredVersion_818_minor
	return nil
}
func (ϟa *GlGetProgramResourceLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_820_major := uint32(3) // u32
	minRequiredVersion_820_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM:
	default:
		glErrorInvalidEnum_821_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_821_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_820_major, minRequiredVersion_820_minor
	return nil
}
func (ϟa *GlGetProgramResourceName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_822_major := uint32(3) // u32
	minRequiredVersion_822_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_823_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_823_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_822_major, minRequiredVersion_822_minor
	return nil
}
func (ϟa *GlGetProgramResourceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_824_major := uint32(3) // u32
	minRequiredVersion_824_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_825_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_825_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_824_major, minRequiredVersion_824_minor
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_826_major := uint32(2) // u32
	minRequiredVersion_826_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_ACTIVE_ATTRIBUTES, GLenum_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH, GLenum_GL_ACTIVE_UNIFORMS, GLenum_GL_ACTIVE_UNIFORM_MAX_LENGTH, GLenum_GL_ATTACHED_SHADERS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_LINK_STATUS, GLenum_GL_VALIDATE_STATUS:
	case GLenum_GL_ACTIVE_UNIFORM_BLOCKS, GLenum_GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH, GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_MODE, GLenum_GL_TRANSFORM_FEEDBACK_VARYINGS, GLenum_GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		minRequiredVersion_827_major := uint32(3) // u32
		minRequiredVersion_827_minor := uint32(0) // u32
		_, _ = minRequiredVersion_827_major, minRequiredVersion_827_minor
	case GLenum_GL_ACTIVE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_828_major := uint32(3) // u32
		minRequiredVersion_828_minor := uint32(1) // u32
		_, _ = minRequiredVersion_828_major, minRequiredVersion_828_minor
	case GLenum_GL_GEOMETRY_INPUT_TYPE, GLenum_GL_GEOMETRY_OUTPUT_TYPE, GLenum_GL_GEOMETRY_VERTICES_OUT, GLenum_GL_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_TESS_CONTROL_OUTPUT_VERTICES, GLenum_GL_TESS_GEN_MODE, GLenum_GL_TESS_GEN_POINT_MODE, GLenum_GL_TESS_GEN_SPACING, GLenum_GL_TESS_GEN_VERTEX_ORDER:
		minRequiredVersion_829_major := uint32(3) // u32
		minRequiredVersion_829_minor := uint32(2) // u32
		_, _ = minRequiredVersion_829_major, minRequiredVersion_829_minor
	default:
		glErrorInvalidEnum_830_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_830_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_826_major, minRequiredVersion_826_minor
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_831_major := uint32(2)    // u32
	minRequiredVersion_831_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_833_msg := "No context bound" // string
		return
		_ = error_833_msg
	}
	GetContext_832_result := context          // Contextʳ
	ctx := GetContext_832_result              // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
	min_834_a := int32(ϟa.BufferLength)       // s32
	min_834_b := int32(s.InfoLog.Count)       // s32
	min_834_result := func() (result int32) {
		switch (min_834_a) < (min_834_b) {
		case true:
			return min_834_a
		case false:
			return min_834_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_834_a) < (min_834_b), ϟa))
			return result
		}
	}() // s32
	l := min_834_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(s.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_831_major, minRequiredVersion_831_minor, context, GetContext_832_result, ctx, s, min_834_a, min_834_b, min_834_result, l
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_835_major := uint32(2) // u32
	minRequiredVersion_835_minor := uint32(0) // u32
	switch ϟa.ShaderType {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	default:
		glErrorInvalidEnum_836_param := ϟa.ShaderType // GLenum
		return
		_ = glErrorInvalidEnum_836_param
	}
	switch ϟa.PrecisionType {
	case GLenum_GL_HIGH_FLOAT, GLenum_GL_HIGH_INT, GLenum_GL_LOW_FLOAT, GLenum_GL_LOW_INT, GLenum_GL_MEDIUM_FLOAT, GLenum_GL_MEDIUM_INT:
	default:
		glErrorInvalidEnum_837_param := ϟa.PrecisionType // GLenum
		return
		_ = glErrorInvalidEnum_837_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Range.Slice(uint64(0), uint64(2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_835_major, minRequiredVersion_835_minor
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_838_major := uint32(2)    // u32
	minRequiredVersion_838_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_840_msg := "No context bound" // string
		return
		_ = error_840_msg
	}
	GetContext_839_result := context          // Contextʳ
	ctx := GetContext_839_result              // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
	min_841_a := int32(ϟa.BufferLength)       // s32
	min_841_b := int32(len(s.Source))         // s32
	min_841_result := func() (result int32) {
		switch (min_841_a) < (min_841_b) {
		case true:
			return min_841_a
		case false:
			return min_841_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_841_a) < (min_841_b), ϟa))
			return result
		}
	}() // s32
	l := min_841_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	Charᵖ(ϟa.Source).Slice(uint64(int32(0)), uint64(l), ϟs).Copy(MakeCharˢFromString(s.Source, ϟs).Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_838_major, minRequiredVersion_838_minor, context, GetContext_839_result, ctx, s, min_841_a, min_841_b, min_841_result, l
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_842_major := uint32(2) // u32
	minRequiredVersion_842_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_COMPILE_STATUS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_SHADER_SOURCE_LENGTH, GLenum_GL_SHADER_TYPE:
	default:
		glErrorInvalidEnum_843_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_843_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_845_msg := "No context bound" // string
		return
		_ = error_845_msg
	}
	GetContext_844_result := context          // Contextʳ
	ctx := GetContext_844_result              // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
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
	_, _, _, _, _, _ = minRequiredVersion_842_major, minRequiredVersion_842_minor, context, GetContext_844_result, ctx, s
	return nil
}
func (ϟa *GlGetUniformBlockIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_846_major := uint32(3) // u32
	minRequiredVersion_846_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_846_major, minRequiredVersion_846_minor
	return nil
}
func (ϟa *GlGetUniformIndices) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_847_major := uint32(3) // u32
	minRequiredVersion_847_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_847_major, minRequiredVersion_847_minor
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_848_major := uint32(2) // u32
	minRequiredVersion_848_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_848_major, minRequiredVersion_848_minor
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_849_major := uint32(2) // u32
	minRequiredVersion_849_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_849_major, minRequiredVersion_849_minor
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_850_major := uint32(2) // u32
	minRequiredVersion_850_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_850_major, minRequiredVersion_850_minor
	return nil
}
func (ϟa *GlGetUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_851_major := uint32(3) // u32
	minRequiredVersion_851_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_851_major, minRequiredVersion_851_minor
	return nil
}
func (ϟa *GlGetnUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_852_major := uint32(3) // u32
	minRequiredVersion_852_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_852_major, minRequiredVersion_852_minor
	return nil
}
func (ϟa *GlGetnUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_853_major := uint32(3) // u32
	minRequiredVersion_853_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_853_major, minRequiredVersion_853_minor
	return nil
}
func (ϟa *GlGetnUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_854_major := uint32(3) // u32
	minRequiredVersion_854_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_854_major, minRequiredVersion_854_minor
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_855_major := uint32(2)    // u32
	minRequiredVersion_855_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_857_msg := "No context bound" // string
		return
		_ = error_857_msg
	}
	GetContext_856_result := context // Contextʳ
	ctx := GetContext_856_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Programs.Contains(ϟa.Program) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_855_major, minRequiredVersion_855_minor, context, GetContext_856_result, ctx
	return nil
}
func (ϟa *GlIsProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_858_major := uint32(3) // u32
	minRequiredVersion_858_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_858_major, minRequiredVersion_858_minor
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_859_major := uint32(2)    // u32
	minRequiredVersion_859_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_861_msg := "No context bound" // string
		return
		_ = error_861_msg
	}
	GetContext_860_result := context // Contextʳ
	ctx := GetContext_860_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Shaders.Contains(ϟa.Shader) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_859_major, minRequiredVersion_859_minor, context, GetContext_860_result, ctx
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_862_major := uint32(2) // u32
	minRequiredVersion_862_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_862_major, minRequiredVersion_862_minor
	return nil
}
func (ϟa *GlMemoryBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_863_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_863_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_864_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_864_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_863_major, minRequiredVersion_863_minor, supportsBits_864_seenBits, supportsBits_864_validBits
	return nil
}
func (ϟa *GlMemoryBarrierByRegion) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_865_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_865_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_866_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_866_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_865_major, minRequiredVersion_865_minor, supportsBits_866_seenBits, supportsBits_866_validBits
	return nil
}
func (ϟa *GlProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_867_major := uint32(3) // u32
	minRequiredVersion_867_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		glErrorInvalidEnum_868_param := ϟa.BinaryFormat // GLenum
		return
		_ = glErrorInvalidEnum_868_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_867_major, minRequiredVersion_867_minor
	return nil
}
func (ϟa *GlProgramParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_869_major := uint32(3) // u32
	minRequiredVersion_869_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT:
	case GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_870_major := uint32(3) // u32
		minRequiredVersion_870_minor := uint32(1) // u32
		_, _ = minRequiredVersion_870_major, minRequiredVersion_870_minor
	default:
		glErrorInvalidEnum_871_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_871_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_869_major, minRequiredVersion_869_minor
	return nil
}
func (ϟa *GlProgramUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_872_major := uint32(3) // u32
	minRequiredVersion_872_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_872_major, minRequiredVersion_872_minor
	return nil
}
func (ϟa *GlProgramUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_873_major := uint32(3) // u32
	minRequiredVersion_873_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_873_major, minRequiredVersion_873_minor
	return nil
}
func (ϟa *GlProgramUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_874_major := uint32(3) // u32
	minRequiredVersion_874_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_874_major, minRequiredVersion_874_minor
	return nil
}
func (ϟa *GlProgramUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_875_major := uint32(3) // u32
	minRequiredVersion_875_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_875_major, minRequiredVersion_875_minor
	return nil
}
func (ϟa *GlProgramUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_876_major := uint32(3) // u32
	minRequiredVersion_876_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_876_major, minRequiredVersion_876_minor
	return nil
}
func (ϟa *GlProgramUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_877_major := uint32(3) // u32
	minRequiredVersion_877_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_877_major, minRequiredVersion_877_minor
	return nil
}
func (ϟa *GlProgramUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_878_major := uint32(3) // u32
	minRequiredVersion_878_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_878_major, minRequiredVersion_878_minor
	return nil
}
func (ϟa *GlProgramUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_879_major := uint32(3) // u32
	minRequiredVersion_879_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_879_major, minRequiredVersion_879_minor
	return nil
}
func (ϟa *GlProgramUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_880_major := uint32(3) // u32
	minRequiredVersion_880_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_880_major, minRequiredVersion_880_minor
	return nil
}
func (ϟa *GlProgramUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_881_major := uint32(3) // u32
	minRequiredVersion_881_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_881_major, minRequiredVersion_881_minor
	return nil
}
func (ϟa *GlProgramUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_882_major := uint32(3) // u32
	minRequiredVersion_882_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_882_major, minRequiredVersion_882_minor
	return nil
}
func (ϟa *GlProgramUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_883_major := uint32(3) // u32
	minRequiredVersion_883_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_883_major, minRequiredVersion_883_minor
	return nil
}
func (ϟa *GlProgramUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_884_major := uint32(3) // u32
	minRequiredVersion_884_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_884_major, minRequiredVersion_884_minor
	return nil
}
func (ϟa *GlProgramUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_885_major := uint32(3) // u32
	minRequiredVersion_885_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_885_major, minRequiredVersion_885_minor
	return nil
}
func (ϟa *GlProgramUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_886_major := uint32(3) // u32
	minRequiredVersion_886_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_886_major, minRequiredVersion_886_minor
	return nil
}
func (ϟa *GlProgramUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_887_major := uint32(3) // u32
	minRequiredVersion_887_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_887_major, minRequiredVersion_887_minor
	return nil
}
func (ϟa *GlProgramUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_888_major := uint32(3) // u32
	minRequiredVersion_888_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_888_major, minRequiredVersion_888_minor
	return nil
}
func (ϟa *GlProgramUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_889_major := uint32(3) // u32
	minRequiredVersion_889_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_889_major, minRequiredVersion_889_minor
	return nil
}
func (ϟa *GlProgramUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_890_major := uint32(3) // u32
	minRequiredVersion_890_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_890_major, minRequiredVersion_890_minor
	return nil
}
func (ϟa *GlProgramUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_891_major := uint32(3) // u32
	minRequiredVersion_891_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_891_major, minRequiredVersion_891_minor
	return nil
}
func (ϟa *GlProgramUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_892_major := uint32(3) // u32
	minRequiredVersion_892_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_892_major, minRequiredVersion_892_minor
	return nil
}
func (ϟa *GlProgramUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_893_major := uint32(3) // u32
	minRequiredVersion_893_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_893_major, minRequiredVersion_893_minor
	return nil
}
func (ϟa *GlProgramUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_894_major := uint32(3) // u32
	minRequiredVersion_894_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_894_major, minRequiredVersion_894_minor
	return nil
}
func (ϟa *GlProgramUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_895_major := uint32(3) // u32
	minRequiredVersion_895_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_895_major, minRequiredVersion_895_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_896_major := uint32(3) // u32
	minRequiredVersion_896_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_896_major, minRequiredVersion_896_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_897_major := uint32(3) // u32
	minRequiredVersion_897_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_897_major, minRequiredVersion_897_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_898_major := uint32(3) // u32
	minRequiredVersion_898_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_898_major, minRequiredVersion_898_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_899_major := uint32(3) // u32
	minRequiredVersion_899_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_899_major, minRequiredVersion_899_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_900_major := uint32(3) // u32
	minRequiredVersion_900_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_900_major, minRequiredVersion_900_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_901_major := uint32(3) // u32
	minRequiredVersion_901_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_901_major, minRequiredVersion_901_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_902_major := uint32(3) // u32
	minRequiredVersion_902_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_902_major, minRequiredVersion_902_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_903_major := uint32(3) // u32
	minRequiredVersion_903_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_903_major, minRequiredVersion_903_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_904_major := uint32(3) // u32
	minRequiredVersion_904_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_904_major, minRequiredVersion_904_minor
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_905_major := uint32(2) // u32
	minRequiredVersion_905_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_905_major, minRequiredVersion_905_minor
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_906_major := uint32(2) // u32
	minRequiredVersion_906_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		glErrorInvalidEnum_907_param := ϟa.BinaryFormat // GLenum
		return
		_ = glErrorInvalidEnum_907_param
	}
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_906_major, minRequiredVersion_906_minor
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_908_major := uint32(2)                                   // u32
	minRequiredVersion_908_minor := uint32(0)                                   // u32
	sources := ϟa.Source.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	lengths := ϟa.Length.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_910_msg := "No context bound" // string
		return
		_ = error_910_msg
	}
	GetContext_909_result := context          // Contextʳ
	ctx := GetContext_909_result              // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_908_major, minRequiredVersion_908_minor, sources, lengths, context, GetContext_909_result, ctx, s
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_911_major := uint32(2)    // u32
	minRequiredVersion_911_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_913_msg := "No context bound" // string
		return
		_ = error_913_msg
	}
	GetContext_912_result := context // Contextʳ
	ctx := GetContext_912_result     // Contextʳ
	v := MakeGLfloatˢ(uint64(1), ϟs) // GLfloatˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_911_major, minRequiredVersion_911_minor, context, GetContext_912_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_914_major := uint32(2)    // u32
	minRequiredVersion_914_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_916_msg := "No context bound" // string
		return
		_ = error_916_msg
	}
	GetContext_915_result := context                                      // Contextʳ
	ctx := GetContext_915_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLfloatˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_914_major, minRequiredVersion_914_minor, context, GetContext_915_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_917_major := uint32(2)    // u32
	minRequiredVersion_917_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_919_msg := "No context bound" // string
		return
		_ = error_919_msg
	}
	GetContext_918_result := context // Contextʳ
	ctx := GetContext_918_result     // Contextʳ
	v := MakeGLintˢ(uint64(1), ϟs)   // GLintˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_917_major, minRequiredVersion_917_minor, context, GetContext_918_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_920_major := uint32(2)    // u32
	minRequiredVersion_920_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_922_msg := "No context bound" // string
		return
		_ = error_922_msg
	}
	GetContext_921_result := context                                      // Contextʳ
	ctx := GetContext_921_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_920_major, minRequiredVersion_920_minor, context, GetContext_921_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_923_major := uint32(3) // u32
	minRequiredVersion_923_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_923_major, minRequiredVersion_923_minor
	return nil
}
func (ϟa *GlUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_924_major := uint32(3) // u32
	minRequiredVersion_924_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_924_major, minRequiredVersion_924_minor
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_925_major := uint32(2)    // u32
	minRequiredVersion_925_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_927_msg := "No context bound" // string
		return
		_ = error_927_msg
	}
	GetContext_926_result := context // Contextʳ
	ctx := GetContext_926_result     // Contextʳ
	v := MakeVec2fˢ(uint64(1), ϟs)   // Vec2fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2f{Elements: [2]GLfloat{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_925_major, minRequiredVersion_925_minor, context, GetContext_926_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_928_major := uint32(2)    // u32
	minRequiredVersion_928_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_930_msg := "No context bound" // string
		return
		_ = error_930_msg
	}
	GetContext_929_result := context                                              // Contextʳ
	ctx := GetContext_929_result                                                  // Contextʳ
	v := Vec2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_928_major, minRequiredVersion_928_minor, context, GetContext_929_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_931_major := uint32(2)    // u32
	minRequiredVersion_931_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_933_msg := "No context bound" // string
		return
		_ = error_933_msg
	}
	GetContext_932_result := context // Contextʳ
	ctx := GetContext_932_result     // Contextʳ
	v := MakeVec2iˢ(uint64(1), ϟs)   // Vec2iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2i{Elements: [2]GLint{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_931_major, minRequiredVersion_931_minor, context, GetContext_932_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_934_major := uint32(2)    // u32
	minRequiredVersion_934_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_936_msg := "No context bound" // string
		return
		_ = error_936_msg
	}
	GetContext_935_result := context                                              // Contextʳ
	ctx := GetContext_935_result                                                  // Contextʳ
	v := Vec2iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_934_major, minRequiredVersion_934_minor, context, GetContext_935_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_937_major := uint32(3) // u32
	minRequiredVersion_937_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_937_major, minRequiredVersion_937_minor
	return nil
}
func (ϟa *GlUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_938_major := uint32(3) // u32
	minRequiredVersion_938_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_938_major, minRequiredVersion_938_minor
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_939_major := uint32(2)    // u32
	minRequiredVersion_939_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_941_msg := "No context bound" // string
		return
		_ = error_941_msg
	}
	GetContext_940_result := context // Contextʳ
	ctx := GetContext_940_result     // Contextʳ
	v := MakeVec3fˢ(uint64(1), ϟs)   // Vec3fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3f{Elements: [3]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_939_major, minRequiredVersion_939_minor, context, GetContext_940_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_942_major := uint32(2)    // u32
	minRequiredVersion_942_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_944_msg := "No context bound" // string
		return
		_ = error_944_msg
	}
	GetContext_943_result := context                                              // Contextʳ
	ctx := GetContext_943_result                                                  // Contextʳ
	v := Vec3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_942_major, minRequiredVersion_942_minor, context, GetContext_943_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_945_major := uint32(2)    // u32
	minRequiredVersion_945_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_947_msg := "No context bound" // string
		return
		_ = error_947_msg
	}
	GetContext_946_result := context // Contextʳ
	ctx := GetContext_946_result     // Contextʳ
	v := MakeVec3iˢ(uint64(1), ϟs)   // Vec3iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3i{Elements: [3]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_945_major, minRequiredVersion_945_minor, context, GetContext_946_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_948_major := uint32(2)    // u32
	minRequiredVersion_948_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_950_msg := "No context bound" // string
		return
		_ = error_950_msg
	}
	GetContext_949_result := context                                              // Contextʳ
	ctx := GetContext_949_result                                                  // Contextʳ
	v := Vec3iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_948_major, minRequiredVersion_948_minor, context, GetContext_949_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_951_major := uint32(3) // u32
	minRequiredVersion_951_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_951_major, minRequiredVersion_951_minor
	return nil
}
func (ϟa *GlUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_952_major := uint32(3) // u32
	minRequiredVersion_952_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_952_major, minRequiredVersion_952_minor
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_953_major := uint32(2)    // u32
	minRequiredVersion_953_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_955_msg := "No context bound" // string
		return
		_ = error_955_msg
	}
	GetContext_954_result := context // Contextʳ
	ctx := GetContext_954_result     // Contextʳ
	v := MakeVec4fˢ(uint64(1), ϟs)   // Vec4fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4f{Elements: [4]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_953_major, minRequiredVersion_953_minor, context, GetContext_954_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_956_major := uint32(2)    // u32
	minRequiredVersion_956_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_958_msg := "No context bound" // string
		return
		_ = error_958_msg
	}
	GetContext_957_result := context                                              // Contextʳ
	ctx := GetContext_957_result                                                  // Contextʳ
	v := Vec4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_956_major, minRequiredVersion_956_minor, context, GetContext_957_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_959_major := uint32(2)    // u32
	minRequiredVersion_959_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_961_msg := "No context bound" // string
		return
		_ = error_961_msg
	}
	GetContext_960_result := context // Contextʳ
	ctx := GetContext_960_result     // Contextʳ
	v := MakeVec4iˢ(uint64(1), ϟs)   // Vec4iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4i{Elements: [4]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_959_major, minRequiredVersion_959_minor, context, GetContext_960_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_962_major := uint32(2)    // u32
	minRequiredVersion_962_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_964_msg := "No context bound" // string
		return
		_ = error_964_msg
	}
	GetContext_963_result := context                                              // Contextʳ
	ctx := GetContext_963_result                                                  // Contextʳ
	v := Vec4iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_962_major, minRequiredVersion_962_minor, context, GetContext_963_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_965_major := uint32(3) // u32
	minRequiredVersion_965_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_965_major, minRequiredVersion_965_minor
	return nil
}
func (ϟa *GlUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_966_major := uint32(3) // u32
	minRequiredVersion_966_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_966_major, minRequiredVersion_966_minor
	return nil
}
func (ϟa *GlUniformBlockBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_967_major := uint32(3) // u32
	minRequiredVersion_967_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_967_major, minRequiredVersion_967_minor
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_968_major := uint32(2)    // u32
	minRequiredVersion_968_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_970_msg := "No context bound" // string
		return
		_ = error_970_msg
	}
	GetContext_969_result := context                                              // Contextʳ
	ctx := GetContext_969_result                                                  // Contextʳ
	v := Mat2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_968_major, minRequiredVersion_968_minor, context, GetContext_969_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_971_major := uint32(3) // u32
	minRequiredVersion_971_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_971_major, minRequiredVersion_971_minor
	return nil
}
func (ϟa *GlUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_972_major := uint32(3) // u32
	minRequiredVersion_972_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_972_major, minRequiredVersion_972_minor
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_973_major := uint32(2)    // u32
	minRequiredVersion_973_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_975_msg := "No context bound" // string
		return
		_ = error_975_msg
	}
	GetContext_974_result := context                                              // Contextʳ
	ctx := GetContext_974_result                                                  // Contextʳ
	v := Mat3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_973_major, minRequiredVersion_973_minor, context, GetContext_974_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_976_major := uint32(3) // u32
	minRequiredVersion_976_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_976_major, minRequiredVersion_976_minor
	return nil
}
func (ϟa *GlUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_977_major := uint32(3) // u32
	minRequiredVersion_977_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_977_major, minRequiredVersion_977_minor
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_978_major := uint32(2)    // u32
	minRequiredVersion_978_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_980_msg := "No context bound" // string
		return
		_ = error_980_msg
	}
	GetContext_979_result := context                                              // Contextʳ
	ctx := GetContext_979_result                                                  // Contextʳ
	v := Mat4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_978_major, minRequiredVersion_978_minor, context, GetContext_979_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_981_major := uint32(3) // u32
	minRequiredVersion_981_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_981_major, minRequiredVersion_981_minor
	return nil
}
func (ϟa *GlUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_982_major := uint32(3) // u32
	minRequiredVersion_982_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_982_major, minRequiredVersion_982_minor
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_983_major := uint32(2)    // u32
	minRequiredVersion_983_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_985_msg := "No context bound" // string
		return
		_ = error_985_msg
	}
	GetContext_984_result := context // Contextʳ
	ctx := GetContext_984_result     // Contextʳ
	ctx.BoundProgram = ϟa.Program
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_983_major, minRequiredVersion_983_minor, context, GetContext_984_result, ctx
	return nil
}
func (ϟa *GlUseProgramStages) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_986_major := uint32(3)                                                                                                                                        // u32
	minRequiredVersion_986_minor := uint32(1)                                                                                                                                        // u32
	supportsBits_987_seenBits := ϟa.Stages                                                                                                                                           // GLbitfield
	supportsBits_987_validBits := (GLbitfield_GL_ALL_SHADER_BITS) | ((GLbitfield_GL_COMPUTE_SHADER_BIT) | ((GLbitfield_GL_FRAGMENT_SHADER_BIT) | (GLbitfield_GL_VERTEX_SHADER_BIT))) // GLbitfield
	if (GLbitfield_GL_ALL_SHADER_BITS)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_COMPUTE_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_FRAGMENT_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_VERTEX_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_986_major, minRequiredVersion_986_minor, supportsBits_987_seenBits, supportsBits_987_validBits
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_988_major := uint32(2) // u32
	minRequiredVersion_988_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_988_major, minRequiredVersion_988_minor
	return nil
}
func (ϟa *GlValidateProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_989_major := uint32(3) // u32
	minRequiredVersion_989_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_989_major, minRequiredVersion_989_minor
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_990_major := uint32(2) // u32
	minRequiredVersion_990_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_991_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_991_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_993_msg := "No context bound" // string
		return
		_ = error_993_msg
	}
	GetContext_992_result := context // Contextʳ
	ctx := GetContext_992_result     // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_990_major, minRequiredVersion_990_minor, context, GetContext_992_result, ctx
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_994_major := uint32(2)    // u32
	minRequiredVersion_994_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_996_msg := "No context bound" // string
		return
		_ = error_996_msg
	}
	GetContext_995_result := context // Contextʳ
	ctx := GetContext_995_result     // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_994_major, minRequiredVersion_994_minor, context, GetContext_995_result, ctx
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_997_major := uint32(2) // u32
	minRequiredVersion_997_minor := uint32(0) // u32
	switch ϟa.Orientation {
	case GLenum_GL_CCW, GLenum_GL_CW:
	default:
		glErrorInvalidEnum_998_param := ϟa.Orientation // GLenum
		return
		_ = glErrorInvalidEnum_998_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1000_msg := "No context bound" // string
		return
		_ = error_1000_msg
	}
	GetContext_999_result := context // Contextʳ
	ctx := GetContext_999_result     // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_997_major, minRequiredVersion_997_minor, context, GetContext_999_result, ctx
	return nil
}
func (ϟa *GlGetMultisamplefv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1001_major := uint32(3) // u32
	minRequiredVersion_1001_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_SAMPLE_POSITION:
	default:
		glErrorInvalidEnum_1002_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1002_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1001_major, minRequiredVersion_1001_minor
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1003_major := uint32(2)   // u32
	minRequiredVersion_1003_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1005_msg := "No context bound" // string
		return
		_ = error_1005_msg
	}
	GetContext_1004_result := context // Contextʳ
	ctx := GetContext_1004_result     // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1003_major, minRequiredVersion_1003_minor, context, GetContext_1004_result, ctx
	return nil
}
func (ϟa *GlMinSampleShading) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1006_major := uint32(3) // u32
	minRequiredVersion_1006_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1006_major, minRequiredVersion_1006_minor
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1007_major := uint32(2)   // u32
	minRequiredVersion_1007_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1009_msg := "No context bound" // string
		return
		_ = error_1009_msg
	}
	GetContext_1008_result := context // Contextʳ
	ctx := GetContext_1008_result     // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1007_major, minRequiredVersion_1007_minor, context, GetContext_1008_result, ctx
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1010_major := uint32(2)   // u32
	minRequiredVersion_1010_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1012_msg := "No context bound" // string
		return
		_ = error_1012_msg
	}
	GetContext_1011_result := context // Contextʳ
	ctx := GetContext_1011_result     // Contextʳ
	ctx.Rasterizing.Viewport = Rect{X: ϟa.X, Y: ϟa.Y, Width: ϟa.Width, Height: ϟa.Height}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1010_major, minRequiredVersion_1010_minor, context, GetContext_1011_result, ctx
	return nil
}
func (ϟa *GlGetBooleani_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1013_major := uint32(3) // u32
	minRequiredVersion_1013_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE, GLenum_GL_VIEWPORT:
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1014_major := uint32(3) // u32
		minRequiredVersion_1014_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1014_major, minRequiredVersion_1014_minor
	default:
		glErrorInvalidEnum_1015_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1015_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1013_major, minRequiredVersion_1013_minor
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1016_major := uint32(2) // u32
	minRequiredVersion_1016_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_1017_major := uint32(3) // u32
		minRequiredVersion_1017_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1017_major, minRequiredVersion_1017_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1018_major := uint32(3) // u32
		minRequiredVersion_1018_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1018_major, minRequiredVersion_1018_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1019_major := uint32(3) // u32
		minRequiredVersion_1019_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1019_major, minRequiredVersion_1019_minor
	default:
		glErrorInvalidEnum_1020_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1020_param
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLbooleanˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1022_msg := "No context bound" // string
		return
		_ = error_1022_msg
	}
	GetContext_1021_result := context // Contextʳ
	ctx := GetContext_1021_result     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1016_major, minRequiredVersion_1016_minor, v, context, GetContext_1021_result, ctx
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1023_major := uint32(2) // u32
	minRequiredVersion_1023_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_1024_major := uint32(3) // u32
		minRequiredVersion_1024_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1024_major, minRequiredVersion_1024_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1025_major := uint32(3) // u32
		minRequiredVersion_1025_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1025_major, minRequiredVersion_1025_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1026_major := uint32(3) // u32
		minRequiredVersion_1026_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1026_major, minRequiredVersion_1026_minor
	default:
		glErrorInvalidEnum_1027_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1027_param
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLfloatˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1029_msg := "No context bound" // string
		return
		_ = error_1029_msg
	}
	GetContext_1028_result := context // Contextʳ
	ctx := GetContext_1028_result     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1023_major, minRequiredVersion_1023_minor, v, context, GetContext_1028_result, ctx
	return nil
}
func (ϟa *GlGetInteger64i_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1030_major := uint32(3) // u32
	minRequiredVersion_1030_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1031_major := uint32(3) // u32
		minRequiredVersion_1031_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1031_major, minRequiredVersion_1031_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1032_major := uint32(3) // u32
		minRequiredVersion_1032_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1032_major, minRequiredVersion_1032_minor
	default:
		glErrorInvalidEnum_1033_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1033_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1030_major, minRequiredVersion_1030_minor
	return nil
}
func (ϟa *GlGetInteger64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1034_major := uint32(3) // u32
	minRequiredVersion_1034_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1035_major := uint32(3) // u32
		minRequiredVersion_1035_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1035_major, minRequiredVersion_1035_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1036_major := uint32(3) // u32
		minRequiredVersion_1036_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1036_major, minRequiredVersion_1036_minor
	default:
		glErrorInvalidEnum_1037_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1037_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1034_major, minRequiredVersion_1034_minor
	return nil
}
func (ϟa *GlGetIntegeri_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1038_major := uint32(3) // u32
	minRequiredVersion_1038_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1039_major := uint32(3) // u32
		minRequiredVersion_1039_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1039_major, minRequiredVersion_1039_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1040_major := uint32(3) // u32
		minRequiredVersion_1040_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1040_major, minRequiredVersion_1040_minor
	default:
		glErrorInvalidEnum_1041_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1041_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1038_major, minRequiredVersion_1038_minor
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1042_major := uint32(2) // u32
	minRequiredVersion_1042_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_1043_major := uint32(3) // u32
		minRequiredVersion_1043_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1043_major, minRequiredVersion_1043_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1044_major := uint32(3) // u32
		minRequiredVersion_1044_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1044_major, minRequiredVersion_1044_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1045_major := uint32(3) // u32
		minRequiredVersion_1045_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1045_major, minRequiredVersion_1045_minor
	default:
		glErrorInvalidEnum_1046_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1046_param
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1048_msg := "No context bound" // string
		return
		_ = error_1048_msg
	}
	GetContext_1047_result := context // Contextʳ
	ctx := GetContext_1047_result     // Contextʳ
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
	case GLenum_GL_MAJOR_VERSION:
		v.Index(uint64(0), ϟs).Write(ctx.Info.VersionMajor, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MINOR_VERSION:
		v.Index(uint64(0), ϟs).Write(ctx.Info.VersionMinor, ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_1042_major, minRequiredVersion_1042_minor, v, context, GetContext_1047_result, ctx
	return nil
}
func (ϟa *GlGetInternalformativ) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1049_major := uint32(3) // u32
	minRequiredVersion_1049_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1050_major := uint32(3) // u32
		minRequiredVersion_1050_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1050_major, minRequiredVersion_1050_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY:
		minRequiredVersion_1051_major := uint32(3) // u32
		minRequiredVersion_1051_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1051_major, minRequiredVersion_1051_minor
	default:
		glErrorInvalidEnum_1052_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1052_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1053_major := uint32(3) // u32
		minRequiredVersion_1053_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1053_major, minRequiredVersion_1053_minor
	default:
		glErrorInvalidEnum_1054_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1054_param
	}
	switch ϟa.Pname {
	case GLenum_GL_NUM_SAMPLE_COUNTS, GLenum_GL_SAMPLES:
	default:
		glErrorInvalidEnum_1055_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1055_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1049_major, minRequiredVersion_1049_minor
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1056_major := uint32(2) // u32
	minRequiredVersion_1056_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_EXTENSIONS, GLenum_GL_RENDERER, GLenum_GL_SHADING_LANGUAGE_VERSION, GLenum_GL_VENDOR, GLenum_GL_VERSION:
	default:
		glErrorInvalidEnum_1057_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1057_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = GLubyteᶜᵖ{}
	_, _ = minRequiredVersion_1056_major, minRequiredVersion_1056_minor
	return nil
}
func (ϟa *GlGetStringi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1058_major := uint32(3) // u32
	minRequiredVersion_1058_minor := uint32(0) // u32
	switch ϟa.Name {
	case GLenum_GL_EXTENSIONS:
	default:
		glErrorInvalidEnum_1059_param := ϟa.Name // GLenum
		return
		_ = glErrorInvalidEnum_1059_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1058_major, minRequiredVersion_1058_minor
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1060_major := uint32(2) // u32
	minRequiredVersion_1060_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_1061_major := uint32(3) // u32
		minRequiredVersion_1061_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1061_major, minRequiredVersion_1061_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_1062_major := uint32(3) // u32
		minRequiredVersion_1062_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1062_major, minRequiredVersion_1062_minor
	default:
		glErrorInvalidEnum_1063_param := ϟa.Capability // GLenum
		return
		_ = glErrorInvalidEnum_1063_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1065_msg := "No context bound" // string
		return
		_ = error_1065_msg
	}
	GetContext_1064_result := context // Contextʳ
	ctx := GetContext_1064_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Capabilities.Get(ϟa.Capability) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_1060_major, minRequiredVersion_1060_minor, context, GetContext_1064_result, ctx
	return nil
}
func (ϟa *GlIsEnabledi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1066_major := uint32(3) // u32
	minRequiredVersion_1066_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		glErrorInvalidEnum_1067_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1067_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1066_major, minRequiredVersion_1066_minor
	return nil
}
func (ϟa *GlClientWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1068_major := uint32(3)                           // u32
	minRequiredVersion_1068_minor := uint32(0)                           // u32
	supportsBits_1069_seenBits := ϟa.SyncFlags                           // GLbitfield
	supportsBits_1069_validBits := GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT // GLbitfield
	if (GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT)&(ϟa.SyncFlags) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _, _, _ = minRequiredVersion_1068_major, minRequiredVersion_1068_minor, supportsBits_1069_seenBits, supportsBits_1069_validBits
	return nil
}
func (ϟa *GlDeleteSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1070_major := uint32(3) // u32
	minRequiredVersion_1070_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1070_major, minRequiredVersion_1070_minor
	return nil
}
func (ϟa *GlFenceSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1071_major := uint32(3) // u32
	minRequiredVersion_1071_minor := uint32(0) // u32
	switch ϟa.Condition {
	case GLenum_GL_SYNC_GPU_COMMANDS_COMPLETE:
	default:
		glErrorInvalidEnum_1072_param := ϟa.Condition // GLenum
		return
		_ = glErrorInvalidEnum_1072_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1071_major, minRequiredVersion_1071_minor
	return nil
}
func (ϟa *GlGetSynciv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1073_major := uint32(3) // u32
	minRequiredVersion_1073_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_OBJECT_TYPE, GLenum_GL_SYNC_CONDITION, GLenum_GL_SYNC_FLAGS, GLenum_GL_SYNC_STATUS:
	default:
		glErrorInvalidEnum_1074_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1074_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1073_major, minRequiredVersion_1073_minor
	return nil
}
func (ϟa *GlIsSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1075_major := uint32(3) // u32
	minRequiredVersion_1075_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1075_major, minRequiredVersion_1075_minor
	return nil
}
func (ϟa *GlWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1076_major := uint32(3) // u32
	minRequiredVersion_1076_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1076_major, minRequiredVersion_1076_minor
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1077_major := uint32(2) // u32
	minRequiredVersion_1077_minor := uint32(0) // u32
	switch ϟa.Unit {
	case GLenum_GL_TEXTURE0, GLenum_GL_TEXTURE1, GLenum_GL_TEXTURE10, GLenum_GL_TEXTURE11, GLenum_GL_TEXTURE12, GLenum_GL_TEXTURE13, GLenum_GL_TEXTURE14, GLenum_GL_TEXTURE15, GLenum_GL_TEXTURE16, GLenum_GL_TEXTURE17, GLenum_GL_TEXTURE18, GLenum_GL_TEXTURE19, GLenum_GL_TEXTURE2, GLenum_GL_TEXTURE20, GLenum_GL_TEXTURE21, GLenum_GL_TEXTURE22, GLenum_GL_TEXTURE23, GLenum_GL_TEXTURE24, GLenum_GL_TEXTURE25, GLenum_GL_TEXTURE26, GLenum_GL_TEXTURE27, GLenum_GL_TEXTURE28, GLenum_GL_TEXTURE29, GLenum_GL_TEXTURE3, GLenum_GL_TEXTURE30, GLenum_GL_TEXTURE31, GLenum_GL_TEXTURE4, GLenum_GL_TEXTURE5, GLenum_GL_TEXTURE6, GLenum_GL_TEXTURE7, GLenum_GL_TEXTURE8, GLenum_GL_TEXTURE9:
	default:
		glErrorInvalidEnum_1078_param := ϟa.Unit // GLenum
		return
		_ = glErrorInvalidEnum_1078_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1080_msg := "No context bound" // string
		return
		_ = error_1080_msg
	}
	GetContext_1079_result := context // Contextʳ
	ctx := GetContext_1079_result     // Contextʳ
	ctx.ActiveTextureUnit = ϟa.Unit
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1077_major, minRequiredVersion_1077_minor, context, GetContext_1079_result, ctx
	return nil
}
func (ϟa *GlBindImageTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1081_major := uint32(3) // u32
	minRequiredVersion_1081_minor := uint32(1) // u32
	switch ϟa.Access {
	case GLenum_GL_READ_ONLY, GLenum_GL_READ_WRITE, GLenum_GL_WRITE_ONLY:
	default:
		glErrorInvalidEnum_1082_param := ϟa.Access // GLenum
		return
		_ = glErrorInvalidEnum_1082_param
	}
	switch ϟa.Format {
	case GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM:
	default:
		glErrorInvalidEnum_1083_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1083_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1081_major, minRequiredVersion_1081_minor
	return nil
}
func (ϟa *GlBindSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1084_major := uint32(3) // u32
	minRequiredVersion_1084_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1084_major, minRequiredVersion_1084_minor
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1085_major := uint32(2) // u32
	minRequiredVersion_1085_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1086_major := uint32(3) // u32
		minRequiredVersion_1086_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1086_major, minRequiredVersion_1086_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1087_major := uint32(3) // u32
		minRequiredVersion_1087_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1087_major, minRequiredVersion_1087_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1088_major := uint32(3) // u32
		minRequiredVersion_1088_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1088_major, minRequiredVersion_1088_minor
	default:
		glErrorInvalidEnum_1089_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1089_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1091_msg := "No context bound" // string
		return
		_ = error_1091_msg
	}
	GetContext_1090_result := context // Contextʳ
	ctx := GetContext_1090_result     // Contextʳ
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = (&Texture{ID: ϟa.Texture, Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	}
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	tu.Bindings[ϟa.Target] = ϟa.Texture
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1085_major, minRequiredVersion_1085_minor, context, GetContext_1090_result, ctx, tu
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1092_major := uint32(2) // u32
	minRequiredVersion_1092_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1093_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1093_param
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_1094_major := uint32(3) // u32
		minRequiredVersion_1094_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1094_major, minRequiredVersion_1094_minor
	case GLenum_GL_ATC_RGB_AMD, GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD, GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD:
		requiresExtension_1095_ext := ExtensionId_GL_AMD_compressed_ATC_texture // ExtensionId
		_ = requiresExtension_1095_ext
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1096_major := uint32(3) // u32
		minRequiredVersion_1096_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1096_major, minRequiredVersion_1096_minor
	default:
		glErrorInvalidEnum_1097_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1097_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1099_msg := "No context bound" // string
		return
		_ = error_1099_msg
	}
	GetContext_1098_result := context                 // Contextʳ
	ctx := GetContext_1098_result                     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1092_major, minRequiredVersion_1092_minor, context, GetContext_1098_result, ctx, tu
	return nil
}
func (ϟa *GlCompressedTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1100_major := uint32(3) // u32
	minRequiredVersion_1100_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1101_major := uint32(3) // u32
		minRequiredVersion_1101_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1101_major, minRequiredVersion_1101_minor
	default:
		glErrorInvalidEnum_1102_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1102_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1103_major := uint32(3) // u32
		minRequiredVersion_1103_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1103_major, minRequiredVersion_1103_minor
	default:
		glErrorInvalidEnum_1104_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1104_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1100_major, minRequiredVersion_1100_minor
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1105_major := uint32(2) // u32
	minRequiredVersion_1105_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1106_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1106_param
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_1107_major := uint32(3) // u32
		minRequiredVersion_1107_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1107_major, minRequiredVersion_1107_minor
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1108_major := uint32(3) // u32
		minRequiredVersion_1108_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1108_major, minRequiredVersion_1108_minor
	default:
		glErrorInvalidEnum_1109_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1109_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1105_major, minRequiredVersion_1105_minor
	return nil
}
func (ϟa *GlCompressedTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1110_major := uint32(3) // u32
	minRequiredVersion_1110_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1111_major := uint32(3) // u32
		minRequiredVersion_1111_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1111_major, minRequiredVersion_1111_minor
	default:
		glErrorInvalidEnum_1112_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1112_param
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1113_major := uint32(3) // u32
		minRequiredVersion_1113_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1113_major, minRequiredVersion_1113_minor
	default:
		glErrorInvalidEnum_1114_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1114_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1110_major, minRequiredVersion_1110_minor
	return nil
}
func (ϟa *GlCopyImageSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1115_major := uint32(3) // u32
	minRequiredVersion_1115_minor := uint32(2) // u32
	switch ϟa.SrcTarget {
	case GLenum_GL_RENDERBUFFER, GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1116_param := ϟa.SrcTarget // GLenum
		return
		_ = glErrorInvalidEnum_1116_param
	}
	switch ϟa.DstTarget {
	case GLenum_GL_RENDERBUFFER, GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1117_param := ϟa.DstTarget // GLenum
		return
		_ = glErrorInvalidEnum_1117_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1115_major, minRequiredVersion_1115_minor
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1118_major := uint32(2) // u32
	minRequiredVersion_1118_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1119_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1119_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_1120_major := uint32(3) // u32
		minRequiredVersion_1120_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1120_major, minRequiredVersion_1120_minor
	default:
		glErrorInvalidEnum_1121_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1121_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1118_major, minRequiredVersion_1118_minor
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1122_major := uint32(2) // u32
	minRequiredVersion_1122_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1123_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1123_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1122_major, minRequiredVersion_1122_minor
	return nil
}
func (ϟa *GlCopyTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1124_major := uint32(3) // u32
	minRequiredVersion_1124_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1125_major := uint32(3) // u32
		minRequiredVersion_1125_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1125_major, minRequiredVersion_1125_minor
	default:
		glErrorInvalidEnum_1126_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1126_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1124_major, minRequiredVersion_1124_minor
	return nil
}
func (ϟa *GlDeleteSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1127_major := uint32(3) // u32
	minRequiredVersion_1127_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1127_major, minRequiredVersion_1127_minor
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1128_major := uint32(2)                              // u32
	minRequiredVersion_1128_minor := uint32(0)                              // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1130_msg := "No context bound" // string
		return
		_ = error_1130_msg
	}
	GetContext_1129_result := context // Contextʳ
	ctx := GetContext_1129_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Textures, t.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1128_major, minRequiredVersion_1128_minor, t, context, GetContext_1129_result, ctx
	return nil
}
func (ϟa *GlGenSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1131_major := uint32(3) // u32
	minRequiredVersion_1131_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1131_major, minRequiredVersion_1131_minor
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1132_major := uint32(2)                              // u32
	minRequiredVersion_1132_minor := uint32(0)                              // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1134_msg := "No context bound" // string
		return
		_ = error_1134_msg
	}
	GetContext_1133_result := context // Contextʳ
	ctx := GetContext_1133_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // TextureId
		ctx.Instances.Textures[id] = (&Texture{ID: id, Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
		t.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1132_major, minRequiredVersion_1132_minor, t, context, GetContext_1133_result, ctx
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1135_major := uint32(2) // u32
	minRequiredVersion_1135_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1136_major := uint32(3) // u32
		minRequiredVersion_1136_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1136_major, minRequiredVersion_1136_minor
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1137_major := uint32(3) // u32
		minRequiredVersion_1137_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1137_major, minRequiredVersion_1137_minor
	default:
		glErrorInvalidEnum_1138_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1138_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1135_major, minRequiredVersion_1135_minor
	return nil
}
func (ϟa *GlGetSamplerParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1139_major := uint32(3) // u32
	minRequiredVersion_1139_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1140_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1140_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1139_major, minRequiredVersion_1139_minor
	return nil
}
func (ϟa *GlGetSamplerParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1141_major := uint32(3) // u32
	minRequiredVersion_1141_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1142_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1142_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1141_major, minRequiredVersion_1141_minor
	return nil
}
func (ϟa *GlGetSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1143_major := uint32(3) // u32
	minRequiredVersion_1143_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1144_major := uint32(3) // u32
		minRequiredVersion_1144_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1144_major, minRequiredVersion_1144_minor
	default:
		glErrorInvalidEnum_1145_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1145_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1143_major, minRequiredVersion_1143_minor
	return nil
}
func (ϟa *GlGetSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1146_major := uint32(3) // u32
	minRequiredVersion_1146_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1147_major := uint32(3) // u32
		minRequiredVersion_1147_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1147_major, minRequiredVersion_1147_minor
	default:
		glErrorInvalidEnum_1148_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1148_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1146_major, minRequiredVersion_1146_minor
	return nil
}
func (ϟa *GlGetTexLevelParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1149_major := uint32(3) // u32
	minRequiredVersion_1149_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1150_major := uint32(3) // u32
		minRequiredVersion_1150_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1150_major, minRequiredVersion_1150_minor
	default:
		glErrorInvalidEnum_1151_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1151_param
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	case GLenum_GL_TEXTURE_BUFFER_DATA_STORE_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET, GLenum_GL_TEXTURE_BUFFER_SIZE:
		minRequiredVersion_1152_major := uint32(3) // u32
		minRequiredVersion_1152_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1152_major, minRequiredVersion_1152_minor
	default:
		glErrorInvalidEnum_1153_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1153_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1149_major, minRequiredVersion_1149_minor
	return nil
}
func (ϟa *GlGetTexLevelParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1154_major := uint32(3) // u32
	minRequiredVersion_1154_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1155_major := uint32(3) // u32
		minRequiredVersion_1155_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1155_major, minRequiredVersion_1155_minor
	default:
		glErrorInvalidEnum_1156_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1156_param
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	case GLenum_GL_TEXTURE_BUFFER_DATA_STORE_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET, GLenum_GL_TEXTURE_BUFFER_SIZE:
		minRequiredVersion_1157_major := uint32(3) // u32
		minRequiredVersion_1157_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1157_major, minRequiredVersion_1157_minor
	default:
		glErrorInvalidEnum_1158_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1158_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1154_major, minRequiredVersion_1154_minor
	return nil
}
func (ϟa *GlGetTexParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1159_major := uint32(3) // u32
	minRequiredVersion_1159_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1160_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1160_param
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1161_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1161_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1159_major, minRequiredVersion_1159_minor
	return nil
}
func (ϟa *GlGetTexParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1162_major := uint32(3) // u32
	minRequiredVersion_1162_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1163_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1163_param
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1164_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1164_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1162_major, minRequiredVersion_1162_minor
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1165_major := uint32(2) // u32
	minRequiredVersion_1165_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1166_major := uint32(3) // u32
		minRequiredVersion_1166_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1166_major, minRequiredVersion_1166_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1167_major := uint32(3) // u32
		minRequiredVersion_1167_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1167_major, minRequiredVersion_1167_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1168_major := uint32(3) // u32
		minRequiredVersion_1168_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1168_major, minRequiredVersion_1168_minor
	default:
		glErrorInvalidEnum_1169_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1169_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1170_major := uint32(3) // u32
		minRequiredVersion_1170_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1170_major, minRequiredVersion_1170_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1171_major := uint32(3) // u32
		minRequiredVersion_1171_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1171_major, minRequiredVersion_1171_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1172_major := uint32(3) // u32
		minRequiredVersion_1172_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1172_major, minRequiredVersion_1172_minor
	default:
		glErrorInvalidEnum_1173_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1173_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1175_msg := "No context bound" // string
		return
		_ = error_1175_msg
	}
	GetContext_1174_result := context                 // Contextʳ
	ctx := GetContext_1174_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1165_major, minRequiredVersion_1165_minor, context, GetContext_1174_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1176_major := uint32(2) // u32
	minRequiredVersion_1176_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1177_major := uint32(3) // u32
		minRequiredVersion_1177_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1177_major, minRequiredVersion_1177_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1178_major := uint32(3) // u32
		minRequiredVersion_1178_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1178_major, minRequiredVersion_1178_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1179_major := uint32(3) // u32
		minRequiredVersion_1179_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1179_major, minRequiredVersion_1179_minor
	default:
		glErrorInvalidEnum_1180_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1180_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1181_major := uint32(3) // u32
		minRequiredVersion_1181_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1181_major, minRequiredVersion_1181_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1182_major := uint32(3) // u32
		minRequiredVersion_1182_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1182_major, minRequiredVersion_1182_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1183_major := uint32(3) // u32
		minRequiredVersion_1183_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1183_major, minRequiredVersion_1183_minor
	default:
		glErrorInvalidEnum_1184_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1184_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1186_msg := "No context bound" // string
		return
		_ = error_1186_msg
	}
	GetContext_1185_result := context                 // Contextʳ
	ctx := GetContext_1185_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1176_major, minRequiredVersion_1176_minor, context, GetContext_1185_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlIsSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1187_major := uint32(3) // u32
	minRequiredVersion_1187_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1187_major, minRequiredVersion_1187_minor
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1188_major := uint32(2)   // u32
	minRequiredVersion_1188_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1190_msg := "No context bound" // string
		return
		_ = error_1190_msg
	}
	GetContext_1189_result := context // Contextʳ
	ctx := GetContext_1189_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Textures.Contains(ϟa.Texture) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_1188_major, minRequiredVersion_1188_minor, context, GetContext_1189_result, ctx
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1191_major := uint32(2) // u32
	minRequiredVersion_1191_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_PACK_ALIGNMENT, GLenum_GL_UNPACK_ALIGNMENT:
	case GLenum_GL_PACK_IMAGE_HEIGHT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_IMAGES, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS:
		minRequiredVersion_1192_major := uint32(3) // u32
		minRequiredVersion_1192_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1192_major, minRequiredVersion_1192_minor
	default:
		glErrorInvalidEnum_1193_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1193_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1195_msg := "No context bound" // string
		return
		_ = error_1195_msg
	}
	GetContext_1194_result := context // Contextʳ
	ctx := GetContext_1194_result     // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1191_major, minRequiredVersion_1191_minor, context, GetContext_1194_result, ctx
	return nil
}
func (ϟa *GlSamplerParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1196_major := uint32(3) // u32
	minRequiredVersion_1196_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1197_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1197_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1196_major, minRequiredVersion_1196_minor
	return nil
}
func (ϟa *GlSamplerParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1198_major := uint32(3) // u32
	minRequiredVersion_1198_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1199_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1199_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1198_major, minRequiredVersion_1198_minor
	return nil
}
func (ϟa *GlSamplerParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1200_major := uint32(3) // u32
	minRequiredVersion_1200_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1201_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1201_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1200_major, minRequiredVersion_1200_minor
	return nil
}
func (ϟa *GlSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1202_major := uint32(3) // u32
	minRequiredVersion_1202_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1203_major := uint32(3) // u32
		minRequiredVersion_1203_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1203_major, minRequiredVersion_1203_minor
	default:
		glErrorInvalidEnum_1204_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1204_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1202_major, minRequiredVersion_1202_minor
	return nil
}
func (ϟa *GlSamplerParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1205_major := uint32(3) // u32
	minRequiredVersion_1205_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1206_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1206_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1205_major, minRequiredVersion_1205_minor
	return nil
}
func (ϟa *GlSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1207_major := uint32(3) // u32
	minRequiredVersion_1207_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1208_major := uint32(3) // u32
		minRequiredVersion_1208_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1208_major, minRequiredVersion_1208_minor
	default:
		glErrorInvalidEnum_1209_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1209_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1207_major, minRequiredVersion_1207_minor
	return nil
}
func (ϟa *GlTexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1210_major := uint32(3) // u32
	minRequiredVersion_1210_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_BUFFER:
	default:
		glErrorInvalidEnum_1211_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1211_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_R16, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGBA16, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI:
	default:
		glErrorInvalidEnum_1212_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1212_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1210_major, minRequiredVersion_1210_minor
	return nil
}
func (ϟa *GlTexBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1213_major := uint32(3) // u32
	minRequiredVersion_1213_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_BUFFER:
	default:
		glErrorInvalidEnum_1214_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1214_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_R16, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGBA16, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI:
	default:
		glErrorInvalidEnum_1215_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1215_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1213_major, minRequiredVersion_1213_minor
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1216_major := uint32(2) // u32
	minRequiredVersion_1216_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1217_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1217_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1218_major := uint32(3) // u32
		minRequiredVersion_1218_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1218_major, minRequiredVersion_1218_minor
	case GLenum_GL_STENCIL_INDEX:
		minRequiredVersion_1219_major := uint32(3) // u32
		minRequiredVersion_1219_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1219_major, minRequiredVersion_1219_minor
	default:
		glErrorInvalidEnum_1220_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1220_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1221_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1221_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1222_major := uint32(3) // u32
		minRequiredVersion_1222_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1222_major, minRequiredVersion_1222_minor
	default:
		glErrorInvalidEnum_1223_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1223_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1225_msg := "No context bound" // string
		return
		_ = error_1225_msg
	}
	GetContext_1224_result := context                 // Contextʳ
	ctx := GetContext_1224_result                     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1216_major, minRequiredVersion_1216_minor, context, GetContext_1224_result, ctx, tu
	return nil
}
func (ϟa *GlTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1226_major := uint32(3) // u32
	minRequiredVersion_1226_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1227_major := uint32(3) // u32
		minRequiredVersion_1227_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1227_major, minRequiredVersion_1227_minor
	default:
		glErrorInvalidEnum_1228_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1228_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	case GLenum_GL_STENCIL_INDEX:
		minRequiredVersion_1229_major := uint32(3) // u32
		minRequiredVersion_1229_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1229_major, minRequiredVersion_1229_minor
	default:
		glErrorInvalidEnum_1230_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1230_param
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1231_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1231_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		glErrorInvalidEnum_1232_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1232_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1226_major, minRequiredVersion_1226_minor
	return nil
}
func (ϟa *GlTexParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1233_major := uint32(3) // u32
	minRequiredVersion_1233_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1234_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1234_param
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1235_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1235_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1233_major, minRequiredVersion_1233_minor
	return nil
}
func (ϟa *GlTexParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1236_major := uint32(3) // u32
	minRequiredVersion_1236_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1237_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1237_param
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1238_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1238_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1236_major, minRequiredVersion_1236_minor
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1239_major := uint32(2) // u32
	minRequiredVersion_1239_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1240_major := uint32(3) // u32
		minRequiredVersion_1240_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1240_major, minRequiredVersion_1240_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1241_major := uint32(3) // u32
		minRequiredVersion_1241_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1241_major, minRequiredVersion_1241_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1242_major := uint32(3) // u32
		minRequiredVersion_1242_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1242_major, minRequiredVersion_1242_minor
	default:
		glErrorInvalidEnum_1243_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1243_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1244_major := uint32(3) // u32
		minRequiredVersion_1244_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1244_major, minRequiredVersion_1244_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1245_major := uint32(3) // u32
		minRequiredVersion_1245_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1245_major, minRequiredVersion_1245_minor
	default:
		glErrorInvalidEnum_1246_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1246_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1248_msg := "No context bound" // string
		return
		_ = error_1248_msg
	}
	GetContext_1247_result := context                 // Contextʳ
	ctx := GetContext_1247_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1239_major, minRequiredVersion_1239_minor, context, GetContext_1247_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1249_major := uint32(2) // u32
	minRequiredVersion_1249_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1250_major := uint32(3) // u32
		minRequiredVersion_1250_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1250_major, minRequiredVersion_1250_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1251_major := uint32(3) // u32
		minRequiredVersion_1251_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1251_major, minRequiredVersion_1251_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1252_major := uint32(3) // u32
		minRequiredVersion_1252_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1252_major, minRequiredVersion_1252_minor
	default:
		glErrorInvalidEnum_1253_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1253_param
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1254_major := uint32(3) // u32
		minRequiredVersion_1254_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1254_major, minRequiredVersion_1254_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1255_major := uint32(3) // u32
		minRequiredVersion_1255_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1255_major, minRequiredVersion_1255_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1256_major := uint32(3) // u32
		minRequiredVersion_1256_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1256_major, minRequiredVersion_1256_minor
	default:
		glErrorInvalidEnum_1257_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1257_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1249_major, minRequiredVersion_1249_minor
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1258_major := uint32(2) // u32
	minRequiredVersion_1258_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1259_major := uint32(3) // u32
		minRequiredVersion_1259_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1259_major, minRequiredVersion_1259_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1260_major := uint32(3) // u32
		minRequiredVersion_1260_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1260_major, minRequiredVersion_1260_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1261_major := uint32(3) // u32
		minRequiredVersion_1261_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1261_major, minRequiredVersion_1261_minor
	default:
		glErrorInvalidEnum_1262_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1262_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1263_major := uint32(3) // u32
		minRequiredVersion_1263_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1263_major, minRequiredVersion_1263_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1264_major := uint32(3) // u32
		minRequiredVersion_1264_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1264_major, minRequiredVersion_1264_minor
	default:
		glErrorInvalidEnum_1265_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1265_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1267_msg := "No context bound" // string
		return
		_ = error_1267_msg
	}
	GetContext_1266_result := context                 // Contextʳ
	ctx := GetContext_1266_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1258_major, minRequiredVersion_1258_minor, context, GetContext_1266_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1268_major := uint32(2) // u32
	minRequiredVersion_1268_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1269_major := uint32(3) // u32
		minRequiredVersion_1269_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1269_major, minRequiredVersion_1269_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1270_major := uint32(3) // u32
		minRequiredVersion_1270_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1270_major, minRequiredVersion_1270_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1271_major := uint32(3) // u32
		minRequiredVersion_1271_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1271_major, minRequiredVersion_1271_minor
	default:
		glErrorInvalidEnum_1272_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1272_param
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1273_major := uint32(3) // u32
		minRequiredVersion_1273_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1273_major, minRequiredVersion_1273_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1274_major := uint32(3) // u32
		minRequiredVersion_1274_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1274_major, minRequiredVersion_1274_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1275_major := uint32(3) // u32
		minRequiredVersion_1275_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1275_major, minRequiredVersion_1275_minor
	default:
		glErrorInvalidEnum_1276_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1276_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1268_major, minRequiredVersion_1268_minor
	return nil
}
func (ϟa *GlTexStorage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1277_major := uint32(3) // u32
	minRequiredVersion_1277_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	default:
		glErrorInvalidEnum_1278_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1278_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8, GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1279_major := uint32(3) // u32
		minRequiredVersion_1279_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1279_major, minRequiredVersion_1279_minor
	default:
		glErrorInvalidEnum_1280_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1280_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1277_major, minRequiredVersion_1277_minor
	return nil
}
func (ϟa *GlTexStorage2DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1281_major := uint32(3) // u32
	minRequiredVersion_1281_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
	default:
		glErrorInvalidEnum_1282_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1282_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1283_major := uint32(3) // u32
		minRequiredVersion_1283_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1283_major, minRequiredVersion_1283_minor
	default:
		glErrorInvalidEnum_1284_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1284_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1281_major, minRequiredVersion_1281_minor
	return nil
}
func (ϟa *GlTexStorage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1285_major := uint32(3) // u32
	minRequiredVersion_1285_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1286_major := uint32(3) // u32
		minRequiredVersion_1286_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1286_major, minRequiredVersion_1286_minor
	default:
		glErrorInvalidEnum_1287_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1287_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8, GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1288_major := uint32(3) // u32
		minRequiredVersion_1288_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1288_major, minRequiredVersion_1288_minor
	default:
		glErrorInvalidEnum_1289_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1289_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1285_major, minRequiredVersion_1285_minor
	return nil
}
func (ϟa *GlTexStorage3DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1290_major := uint32(3) // u32
	minRequiredVersion_1290_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY:
	default:
		glErrorInvalidEnum_1291_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1291_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		glErrorInvalidEnum_1292_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1292_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1290_major, minRequiredVersion_1290_minor
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1293_major := uint32(2) // u32
	minRequiredVersion_1293_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1294_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1294_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1295_major := uint32(3) // u32
		minRequiredVersion_1295_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1295_major, minRequiredVersion_1295_minor
	default:
		glErrorInvalidEnum_1296_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1296_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1297_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1297_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1298_major := uint32(3) // u32
		minRequiredVersion_1298_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1298_major, minRequiredVersion_1298_minor
	default:
		glErrorInvalidEnum_1299_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1299_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1301_msg := "No context bound" // string
		return
		_ = error_1301_msg
	}
	GetContext_1300_result := context                 // Contextʳ
	ctx := GetContext_1300_result                     // Contextʳ
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
	for y := uint32(uint32(0)); y < uint32(ϟa.Height); y++ {
		src := (src_stride) * (y)                  // u32
		dst := ((dst_stride) * (y)) + (dst_offset) // u32
		image.Data.Slice(uint64(dst), uint64((dst)+(line_bytes)), ϟs).Copy(src_data.Slice(uint64(src), uint64((src)+(line_bytes)), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = src, dst
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = minRequiredVersion_1293_major, minRequiredVersion_1293_minor, context, GetContext_1300_result, ctx, tu, image, pbo, url, src_width, src_stride, src_size, dst_stride, dst_offset, src_data, line_bytes
	return nil
}
func (ϟa *GlTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1302_major := uint32(3) // u32
	minRequiredVersion_1302_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1303_major := uint32(3) // u32
		minRequiredVersion_1303_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1303_major, minRequiredVersion_1303_minor
	default:
		glErrorInvalidEnum_1304_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1304_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	default:
		glErrorInvalidEnum_1305_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1305_param
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1306_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1306_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		glErrorInvalidEnum_1307_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1307_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1302_major, minRequiredVersion_1302_minor
	return nil
}
func (ϟa *GlBeginTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1308_major := uint32(3) // u32
	minRequiredVersion_1308_minor := uint32(0) // u32
	switch ϟa.PrimitiveMode {
	case GLenum_GL_LINES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES:
	default:
		glErrorInvalidEnum_1309_param := ϟa.PrimitiveMode // GLenum
		return
		_ = glErrorInvalidEnum_1309_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1308_major, minRequiredVersion_1308_minor
	return nil
}
func (ϟa *GlBindTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1310_major := uint32(3) // u32
	minRequiredVersion_1310_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK:
	default:
		glErrorInvalidEnum_1311_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1311_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1310_major, minRequiredVersion_1310_minor
	return nil
}
func (ϟa *GlDeleteTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1312_major := uint32(3) // u32
	minRequiredVersion_1312_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1312_major, minRequiredVersion_1312_minor
	return nil
}
func (ϟa *GlEndTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1313_major := uint32(3) // u32
	minRequiredVersion_1313_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1313_major, minRequiredVersion_1313_minor
	return nil
}
func (ϟa *GlGenTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1314_major := uint32(3) // u32
	minRequiredVersion_1314_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1314_major, minRequiredVersion_1314_minor
	return nil
}
func (ϟa *GlGetTransformFeedbackVarying) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1315_major := uint32(3) // u32
	minRequiredVersion_1315_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1315_major, minRequiredVersion_1315_minor
	return nil
}
func (ϟa *GlIsTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1316_major := uint32(3) // u32
	minRequiredVersion_1316_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1316_major, minRequiredVersion_1316_minor
	return nil
}
func (ϟa *GlPauseTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1317_major := uint32(3) // u32
	minRequiredVersion_1317_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1317_major, minRequiredVersion_1317_minor
	return nil
}
func (ϟa *GlResumeTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1318_major := uint32(3) // u32
	minRequiredVersion_1318_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1318_major, minRequiredVersion_1318_minor
	return nil
}
func (ϟa *GlTransformFeedbackVaryings) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1319_major := uint32(3) // u32
	minRequiredVersion_1319_minor := uint32(0) // u32
	switch ϟa.BufferMode {
	case GLenum_GL_INTERLEAVED_ATTRIBS, GLenum_GL_SEPARATE_ATTRIBS:
	default:
		glErrorInvalidEnum_1320_param := ϟa.BufferMode // GLenum
		return
		_ = glErrorInvalidEnum_1320_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1319_major, minRequiredVersion_1319_minor
	return nil
}
func (ϟa *GlBindVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1321_major := uint32(3)   // u32
	minRequiredVersion_1321_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1323_msg := "No context bound" // string
		return
		_ = error_1323_msg
	}
	GetContext_1322_result := context // Contextʳ
	ctx := GetContext_1322_result     // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = &VertexArray{}
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1321_major, minRequiredVersion_1321_minor, context, GetContext_1322_result, ctx
	return nil
}
func (ϟa *GlBindVertexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1324_major := uint32(3) // u32
	minRequiredVersion_1324_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1324_major, minRequiredVersion_1324_minor
	return nil
}
func (ϟa *GlDeleteVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1325_major := uint32(3)   // u32
	minRequiredVersion_1325_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1327_msg := "No context bound" // string
		return
		_ = error_1327_msg
	}
	GetContext_1326_result := context                                     // Contextʳ
	ctx := GetContext_1326_result                                         // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1325_major, minRequiredVersion_1325_minor, context, GetContext_1326_result, ctx, a
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1328_major := uint32(2)   // u32
	minRequiredVersion_1328_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1330_msg := "No context bound" // string
		return
		_ = error_1330_msg
	}
	GetContext_1329_result := context // Contextʳ
	ctx := GetContext_1329_result     // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1328_major, minRequiredVersion_1328_minor, context, GetContext_1329_result, ctx
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1331_major := uint32(2)   // u32
	minRequiredVersion_1331_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1333_msg := "No context bound" // string
		return
		_ = error_1333_msg
	}
	GetContext_1332_result := context // Contextʳ
	ctx := GetContext_1332_result     // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1331_major, minRequiredVersion_1331_minor, context, GetContext_1332_result, ctx
	return nil
}
func (ϟa *GlGenVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1334_major := uint32(3)                            // u32
	minRequiredVersion_1334_minor := uint32(0)                            // u32
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1336_msg := "No context bound" // string
		return
		_ = error_1336_msg
	}
	GetContext_1335_result := context // Contextʳ
	ctx := GetContext_1335_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = &VertexArray{}
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1334_major, minRequiredVersion_1334_minor, a, context, GetContext_1335_result, ctx
	return nil
}
func (ϟa *GlGetVertexAttribIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1337_major := uint32(3) // u32
	minRequiredVersion_1337_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1338_major := uint32(3) // u32
		minRequiredVersion_1338_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1338_major, minRequiredVersion_1338_minor
	default:
		glErrorInvalidEnum_1339_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1339_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1337_major, minRequiredVersion_1337_minor
	return nil
}
func (ϟa *GlGetVertexAttribIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1340_major := uint32(3) // u32
	minRequiredVersion_1340_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1341_major := uint32(3) // u32
		minRequiredVersion_1341_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1341_major, minRequiredVersion_1341_minor
	default:
		glErrorInvalidEnum_1342_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1342_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1340_major, minRequiredVersion_1340_minor
	return nil
}
func (ϟa *GlGetVertexAttribPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1343_major := uint32(2) // u32
	minRequiredVersion_1343_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_POINTER:
	default:
		glErrorInvalidEnum_1344_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1344_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1343_major, minRequiredVersion_1343_minor
	return nil
}
func (ϟa *GlGetVertexAttribfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1345_major := uint32(2) // u32
	minRequiredVersion_1345_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1346_major := uint32(3) // u32
		minRequiredVersion_1346_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1346_major, minRequiredVersion_1346_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1347_major := uint32(3) // u32
		minRequiredVersion_1347_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1347_major, minRequiredVersion_1347_minor
	default:
		glErrorInvalidEnum_1348_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1348_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1345_major, minRequiredVersion_1345_minor
	return nil
}
func (ϟa *GlGetVertexAttribiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1349_major := uint32(2) // u32
	minRequiredVersion_1349_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1350_major := uint32(3) // u32
		minRequiredVersion_1350_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1350_major, minRequiredVersion_1350_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1351_major := uint32(3) // u32
		minRequiredVersion_1351_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1351_major, minRequiredVersion_1351_minor
	default:
		glErrorInvalidEnum_1352_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1352_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1349_major, minRequiredVersion_1349_minor
	return nil
}
func (ϟa *GlIsVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1353_major := uint32(3) // u32
	minRequiredVersion_1353_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1353_major, minRequiredVersion_1353_minor
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1354_major := uint32(2) // u32
	minRequiredVersion_1354_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1354_major, minRequiredVersion_1354_minor
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1355_major := uint32(2) // u32
	minRequiredVersion_1355_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1355_major, minRequiredVersion_1355_minor
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1356_major := uint32(2) // u32
	minRequiredVersion_1356_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1356_major, minRequiredVersion_1356_minor
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1357_major := uint32(2) // u32
	minRequiredVersion_1357_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(2), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1357_major, minRequiredVersion_1357_minor
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1358_major := uint32(2) // u32
	minRequiredVersion_1358_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1358_major, minRequiredVersion_1358_minor
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1359_major := uint32(2) // u32
	minRequiredVersion_1359_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(3), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1359_major, minRequiredVersion_1359_minor
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1360_major := uint32(2) // u32
	minRequiredVersion_1360_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1360_major, minRequiredVersion_1360_minor
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1361_major := uint32(2) // u32
	minRequiredVersion_1361_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1361_major, minRequiredVersion_1361_minor
	return nil
}
func (ϟa *GlVertexAttribBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1362_major := uint32(3) // u32
	minRequiredVersion_1362_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1362_major, minRequiredVersion_1362_minor
	return nil
}
func (ϟa *GlVertexAttribDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1363_major := uint32(3) // u32
	minRequiredVersion_1363_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1363_major, minRequiredVersion_1363_minor
	return nil
}
func (ϟa *GlVertexAttribFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1364_major := uint32(3) // u32
	minRequiredVersion_1364_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_1365_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1365_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1364_major, minRequiredVersion_1364_minor
	return nil
}
func (ϟa *GlVertexAttribI4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1366_major := uint32(3) // u32
	minRequiredVersion_1366_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1366_major, minRequiredVersion_1366_minor
	return nil
}
func (ϟa *GlVertexAttribI4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1367_major := uint32(3) // u32
	minRequiredVersion_1367_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1367_major, minRequiredVersion_1367_minor
	return nil
}
func (ϟa *GlVertexAttribI4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1368_major := uint32(3) // u32
	minRequiredVersion_1368_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1368_major, minRequiredVersion_1368_minor
	return nil
}
func (ϟa *GlVertexAttribI4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1369_major := uint32(3) // u32
	minRequiredVersion_1369_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1369_major, minRequiredVersion_1369_minor
	return nil
}
func (ϟa *GlVertexAttribIFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1370_major := uint32(3) // u32
	minRequiredVersion_1370_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_1371_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1371_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1370_major, minRequiredVersion_1370_minor
	return nil
}
func (ϟa *GlVertexAttribIPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1372_major := uint32(3) // u32
	minRequiredVersion_1372_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1373_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1373_ext
	case GLenum_GL_BYTE, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_1374_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1374_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1372_major, minRequiredVersion_1372_minor
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1375_major := uint32(2) // u32
	minRequiredVersion_1375_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1376_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1376_ext
	case GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_1377_major := uint32(3) // u32
		minRequiredVersion_1377_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1377_major, minRequiredVersion_1377_minor
	default:
		glErrorInvalidEnum_1378_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1378_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1380_msg := "No context bound" // string
		return
		_ = error_1380_msg
	}
	GetContext_1379_result := context               // Contextʳ
	ctx := GetContext_1379_result                   // Contextʳ
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayʳ
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = ϟa.Data
	a.Buffer = ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1375_major, minRequiredVersion_1375_minor, context, GetContext_1379_result, ctx, a
	return nil
}
func (ϟa *GlVertexBindingDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1381_major := uint32(3) // u32
	minRequiredVersion_1381_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1381_major, minRequiredVersion_1381_minor
	return nil
}
func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
func (ϟa *EglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	CreateContext_1382_result := ctx // Contextʳ
	ϟc.EGLContexts[context] = CreateContext_1382_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1382_result
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1383_context := ϟc.EGLContexts.Get(ϟa.Context) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1383_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1383_context
	return nil
}
func (ϟa *EglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *EglQuerySurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlXCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	CreateContext_1384_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1384_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1384_result
	return nil
}
func (ϟa *GlXCreateNewContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	CreateContext_1385_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1385_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1385_result
	return nil
}
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1386_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1386_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1386_context
	return nil
}
func (ϟa *GlXMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1387_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1387_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1387_context
	return nil
}
func (ϟa *GlXSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlXQueryDrawable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *WglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	CreateContext_1388_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1388_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1388_result
	return nil
}
func (ϟa *WglCreateContextAttribsARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	CreateContext_1389_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1389_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1389_result
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1390_context := ϟc.WGLContexts.Get(ϟa.Hglrc) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1390_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1390_context
	return nil
}
func (ϟa *WglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CGLCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	CreateContext_1391_result := ctx // Contextʳ
	ϟc.CGLContexts[context] = CreateContext_1391_result
	ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(context, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1391_result
	return nil
}
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1392_context := ϟc.CGLContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1392_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1392_context
	return nil
}
func (ϟa *CGLGetSurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
func (ϟa *CGSGetSurfaceBounds) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Bounds.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *CGLFlushDrawable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlGetQueryObjecti64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}
func (ϟa *GlGetQueryObjectui64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}
func (ϟa *ReplayCreateRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *ReplayBindRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *SwitchThread) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.CurrentThread = ϟa.ThreadID
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *ContextInfo) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1394_msg := "No context bound" // string
		return
		_ = error_1394_msg
	}
	GetContext_1393_result := context // Contextʳ
	ctx := GetContext_1393_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_1393_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
	return nil
}
func (ϟa *StartTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *StopTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *FlushPostBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
