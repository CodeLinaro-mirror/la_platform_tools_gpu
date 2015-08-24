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
	ϟa.Ids.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_8_ext
	return nil
}
func (ϟa *GlDebugMessageInsertKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_9_ext := ExtensionId_GL_KHR_debug // ExtensionId
	readString_10_length := ϟa.Length                   // GLsizei
	readString_10_buffer := ϟa.Message                  // GLcharᶜᵖ
	if (readString_10_buffer) != (GLcharᶜᵖ{}) {
		if (readString_10_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_10_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_10_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_10_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_9_ext, readString_10_length, readString_10_buffer
	return nil
}
func (ϟa *GlDisableiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_11_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_11_ext
	return nil
}
func (ϟa *GlEnableiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_12_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_12_ext
	return nil
}
func (ϟa *GlFramebufferTextureEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_13_ext := ExtensionId_GL_EXT_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_13_ext
	return nil
}
func (ϟa *GlGetDebugMessageLogKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_14_ext := ExtensionId_GL_KHR_debug // ExtensionId
	GetDebugMessageLog_15_count := ϟa.Count              // GLuint
	GetDebugMessageLog_15_bufSize := ϟa.BufSize          // GLsizei
	GetDebugMessageLog_15_sources := ϟa.Sources          // GLenumᵖ
	GetDebugMessageLog_15_types := ϟa.Types              // GLenumᵖ
	GetDebugMessageLog_15_ids := ϟa.Ids                  // GLuintᵖ
	GetDebugMessageLog_15_severities := ϟa.Severities    // GLenumᵖ
	GetDebugMessageLog_15_lengths := ϟa.Lengths          // GLsizeiᵖ
	GetDebugMessageLog_15_messageLog := ϟa.MessageLog    // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLuint(ϟa.Result) // GLuint
	if (GetDebugMessageLog_15_sources) != (GLenumᵖ{}) {
		GetDebugMessageLog_15_sources.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_15_types) != (GLenumᵖ{}) {
		GetDebugMessageLog_15_types.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_15_ids) != (GLuintᵖ{}) {
		GetDebugMessageLog_15_ids.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_15_severities) != (GLenumᵖ{}) {
		GetDebugMessageLog_15_severities.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_15_lengths) != (GLsizeiᵖ{}) {
		GetDebugMessageLog_15_lengths.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (l) > (GLuint(uint32(0))) {
		GetDebugMessageLog_15_messageLog.Slice(uint64(GLsizei(int32(0))), uint64(GetDebugMessageLog_15_bufSize), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	GetDebugMessageLog_15_result := l // GLuint
	ϟa.Result = GetDebugMessageLog_15_result
	_, _, _, _, _, _, _, _, _, _, _ = requiresExtension_14_ext, GetDebugMessageLog_15_count, GetDebugMessageLog_15_bufSize, GetDebugMessageLog_15_sources, GetDebugMessageLog_15_types, GetDebugMessageLog_15_ids, GetDebugMessageLog_15_severities, GetDebugMessageLog_15_lengths, GetDebugMessageLog_15_messageLog, l, GetDebugMessageLog_15_result
	return nil
}
func (ϟa *GlGetObjectLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_16_ext := ExtensionId_GL_KHR_debug // ExtensionId
	writeString_17_buffer_size := ϟa.BufSize             // GLsizei
	writeString_17_buffer_bytes_written := ϟa.Length     // GLsizeiᵖ
	writeString_17_buffer := ϟa.Label                    // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_17_buffer) != (GLcharᵖ{})) && ((writeString_17_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_17_buffer_size // GLsizei
		if (writeString_17_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_17_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_17_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_17_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _ = requiresExtension_16_ext, writeString_17_buffer_size, writeString_17_buffer_bytes_written, writeString_17_buffer
	return nil
}
func (ϟa *GlGetObjectPtrLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_18_ext := ExtensionId_GL_KHR_debug // ExtensionId
	writeString_19_buffer_size := ϟa.BufSize             // GLsizei
	writeString_19_buffer_bytes_written := ϟa.Length     // GLsizeiᵖ
	writeString_19_buffer := ϟa.Label                    // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_19_buffer) != (GLcharᵖ{})) && ((writeString_19_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_19_buffer_size // GLsizei
		if (writeString_19_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_19_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_19_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_19_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _ = requiresExtension_18_ext, writeString_19_buffer_size, writeString_19_buffer_bytes_written, writeString_19_buffer
	return nil
}
func (ϟa *GlGetPointervKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_20_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_20_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_21_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_22_major := uint32(3)                            // u32
	minRequiredVersion_22_minor := uint32(0)                            // u32
	GetSamplerParameter_23_sampler := ϟa.Sampler                        // SamplerId
	GetSamplerParameter_23_pname := ϟa.Pname                            // GLenum
	GetSamplerParameter_23_params := ϟa.Params                          // GLintᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetSamplerParameter_23_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetSamplerParameter_23_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		GetSamplerParameter_23_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		glErrorInvalidEnum_24_param := GetSamplerParameter_23_pname // GLenum
		return
		_ = glErrorInvalidEnum_24_param
	}
	_, _, _, _, _, _ = requiresExtension_21_ext, minRequiredVersion_22_major, minRequiredVersion_22_minor, GetSamplerParameter_23_sampler, GetSamplerParameter_23_pname, GetSamplerParameter_23_params
	return nil
}
func (ϟa *GlGetSamplerParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_25_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_26_major := uint32(3)                            // u32
	minRequiredVersion_26_minor := uint32(0)                            // u32
	GetSamplerParameter_27_sampler := ϟa.Sampler                        // SamplerId
	GetSamplerParameter_27_pname := ϟa.Pname                            // GLenum
	GetSamplerParameter_27_params := ϟa.Params                          // GLuintᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetSamplerParameter_27_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetSamplerParameter_27_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		GetSamplerParameter_27_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		glErrorInvalidEnum_28_param := GetSamplerParameter_27_pname // GLenum
		return
		_ = glErrorInvalidEnum_28_param
	}
	_, _, _, _, _, _ = requiresExtension_25_ext, minRequiredVersion_26_major, minRequiredVersion_26_minor, GetSamplerParameter_27_sampler, GetSamplerParameter_27_pname, GetSamplerParameter_27_params
	return nil
}
func (ϟa *GlGetTexParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_29_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_30_major := uint32(3)                            // u32
	minRequiredVersion_30_minor := uint32(0)                            // u32
	GetTexParameter_31_target := ϟa.Target                              // GLenum
	GetTexParameter_31_parameter := ϟa.Pname                            // GLenum
	GetTexParameter_31_params := ϟa.Params                              // GLintᵖ
	minRequiredVersion_32_major := uint32(2)                            // u32
	minRequiredVersion_32_minor := uint32(0)                            // u32
	switch GetTexParameter_31_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_33_major := uint32(3) // u32
		minRequiredVersion_33_minor := uint32(0) // u32
		_, _ = minRequiredVersion_33_major, minRequiredVersion_33_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_34_major := uint32(3) // u32
		minRequiredVersion_34_minor := uint32(1) // u32
		_, _ = minRequiredVersion_34_major, minRequiredVersion_34_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_35_major := uint32(3) // u32
		minRequiredVersion_35_minor := uint32(2) // u32
		_, _ = minRequiredVersion_35_major, minRequiredVersion_35_minor
	default:
		glErrorInvalidEnum_36_param := GetTexParameter_31_target // GLenum
		return
		_ = glErrorInvalidEnum_36_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetTexParameter_31_parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetTexParameter_31_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_37_major := uint32(3) // u32
		minRequiredVersion_37_minor := uint32(0) // u32
		GetTexParameter_31_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_37_major, minRequiredVersion_37_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_38_major := uint32(3) // u32
		minRequiredVersion_38_minor := uint32(1) // u32
		GetTexParameter_31_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_38_major, minRequiredVersion_38_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_39_major := uint32(3) // u32
		minRequiredVersion_39_minor := uint32(2) // u32
		GetTexParameter_31_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_39_major, minRequiredVersion_39_minor
	default:
		glErrorInvalidEnum_40_param := GetTexParameter_31_parameter // GLenum
		return
		_ = glErrorInvalidEnum_40_param
	}
	_, _, _, _, _, _, _, _ = requiresExtension_29_ext, minRequiredVersion_30_major, minRequiredVersion_30_minor, GetTexParameter_31_target, GetTexParameter_31_parameter, GetTexParameter_31_params, minRequiredVersion_32_major, minRequiredVersion_32_minor
	return nil
}
func (ϟa *GlGetTexParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_41_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_42_major := uint32(3)                            // u32
	minRequiredVersion_42_minor := uint32(0)                            // u32
	GetTexParameter_43_target := ϟa.Target                              // GLenum
	GetTexParameter_43_parameter := ϟa.Pname                            // GLenum
	GetTexParameter_43_params := ϟa.Params                              // GLuintᵖ
	minRequiredVersion_44_major := uint32(2)                            // u32
	minRequiredVersion_44_minor := uint32(0)                            // u32
	switch GetTexParameter_43_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_45_major := uint32(3) // u32
		minRequiredVersion_45_minor := uint32(0) // u32
		_, _ = minRequiredVersion_45_major, minRequiredVersion_45_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_46_major := uint32(3) // u32
		minRequiredVersion_46_minor := uint32(1) // u32
		_, _ = minRequiredVersion_46_major, minRequiredVersion_46_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_47_major := uint32(3) // u32
		minRequiredVersion_47_minor := uint32(2) // u32
		_, _ = minRequiredVersion_47_major, minRequiredVersion_47_minor
	default:
		glErrorInvalidEnum_48_param := GetTexParameter_43_target // GLenum
		return
		_ = glErrorInvalidEnum_48_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetTexParameter_43_parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetTexParameter_43_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_49_major := uint32(3) // u32
		minRequiredVersion_49_minor := uint32(0) // u32
		GetTexParameter_43_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_49_major, minRequiredVersion_49_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_50_major := uint32(3) // u32
		minRequiredVersion_50_minor := uint32(1) // u32
		GetTexParameter_43_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_50_major, minRequiredVersion_50_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_51_major := uint32(3) // u32
		minRequiredVersion_51_minor := uint32(2) // u32
		GetTexParameter_43_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_51_major, minRequiredVersion_51_minor
	default:
		glErrorInvalidEnum_52_param := GetTexParameter_43_parameter // GLenum
		return
		_ = glErrorInvalidEnum_52_param
	}
	_, _, _, _, _, _, _, _ = requiresExtension_41_ext, minRequiredVersion_42_major, minRequiredVersion_42_minor, GetTexParameter_43_target, GetTexParameter_43_parameter, GetTexParameter_43_params, minRequiredVersion_44_major, minRequiredVersion_44_minor
	return nil
}
func (ϟa *GlIsEnablediEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_53_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_53_ext
	return nil
}
func (ϟa *GlMinSampleShadingOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_54_ext := ExtensionId_GL_OES_sample_shading // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_54_ext
	return nil
}
func (ϟa *GlObjectLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_55_ext := ExtensionId_GL_KHR_debug // ExtensionId
	readString_56_length := ϟa.Length                    // GLsizei
	readString_56_buffer := ϟa.Label                     // GLcharᶜᵖ
	if (readString_56_buffer) != (GLcharᶜᵖ{}) {
		if (readString_56_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_56_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_56_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_56_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_55_ext, readString_56_length, readString_56_buffer
	return nil
}
func (ϟa *GlObjectPtrLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_57_ext := ExtensionId_GL_KHR_debug // ExtensionId
	readString_58_length := ϟa.Length                    // GLsizei
	readString_58_buffer := ϟa.Label                     // GLcharᶜᵖ
	if (readString_58_buffer) != (GLcharᶜᵖ{}) {
		if (readString_58_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_58_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_58_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_58_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_57_ext, readString_58_length, readString_58_buffer
	return nil
}
func (ϟa *GlPatchParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_59_ext := ExtensionId_GL_EXT_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_59_ext
	return nil
}
func (ϟa *GlPopDebugGroupKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_60_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_60_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_61_ext := ExtensionId_GL_EXT_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_61_ext
	return nil
}
func (ϟa *GlPushDebugGroupKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_62_ext := ExtensionId_GL_KHR_debug // ExtensionId
	readString_63_length := ϟa.Length                    // GLsizei
	readString_63_buffer := ϟa.Message                   // GLcharᶜᵖ
	if (readString_63_buffer) != (GLcharᶜᵖ{}) {
		if (readString_63_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_63_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_63_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_63_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_62_ext, readString_63_length, readString_63_buffer
	return nil
}
func (ϟa *GlSamplerParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_64_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_65_major := uint32(3)                            // u32
	minRequiredVersion_65_minor := uint32(0)                            // u32
	SamplerParameterv_66_sampler := ϟa.Sampler                          // SamplerId
	SamplerParameterv_66_pname := ϟa.Pname                              // GLenum
	SamplerParameterv_66_params := ϟa.Param                             // GLintᶜᵖ
	minRequiredVersion_67_major := uint32(3)                            // u32
	minRequiredVersion_67_minor := uint32(0)                            // u32
	switch SamplerParameterv_66_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		SamplerParameterv_66_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_68_major := uint32(3) // u32
		minRequiredVersion_68_minor := uint32(2) // u32
		SamplerParameterv_66_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_68_major, minRequiredVersion_68_minor
	default:
		glErrorInvalidEnum_69_param := SamplerParameterv_66_pname // GLenum
		return
		_ = glErrorInvalidEnum_69_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = requiresExtension_64_ext, minRequiredVersion_65_major, minRequiredVersion_65_minor, SamplerParameterv_66_sampler, SamplerParameterv_66_pname, SamplerParameterv_66_params, minRequiredVersion_67_major, minRequiredVersion_67_minor
	return nil
}
func (ϟa *GlSamplerParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_70_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_71_major := uint32(3)                            // u32
	minRequiredVersion_71_minor := uint32(0)                            // u32
	SamplerParameterv_72_sampler := ϟa.Sampler                          // SamplerId
	SamplerParameterv_72_pname := ϟa.Pname                              // GLenum
	SamplerParameterv_72_params := ϟa.Param                             // GLuintᶜᵖ
	minRequiredVersion_73_major := uint32(3)                            // u32
	minRequiredVersion_73_minor := uint32(0)                            // u32
	switch SamplerParameterv_72_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		SamplerParameterv_72_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_74_major := uint32(3) // u32
		minRequiredVersion_74_minor := uint32(2) // u32
		SamplerParameterv_72_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_74_major, minRequiredVersion_74_minor
	default:
		glErrorInvalidEnum_75_param := SamplerParameterv_72_pname // GLenum
		return
		_ = glErrorInvalidEnum_75_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = requiresExtension_70_ext, minRequiredVersion_71_major, minRequiredVersion_71_minor, SamplerParameterv_72_sampler, SamplerParameterv_72_pname, SamplerParameterv_72_params, minRequiredVersion_73_major, minRequiredVersion_73_minor
	return nil
}
func (ϟa *GlTexBufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_76_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_76_ext
	return nil
}
func (ϟa *GlTexBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_77_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_77_ext
	return nil
}
func (ϟa *GlTexParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_78_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_79_major := uint32(3)                            // u32
	minRequiredVersion_79_minor := uint32(0)                            // u32
	TexParameterv_80_target := ϟa.Target                                // GLenum
	TexParameterv_80_pname := ϟa.Pname                                  // GLenum
	TexParameterv_80_params := ϟa.Params                                // GLintᶜᵖ
	switch TexParameterv_80_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_81_major := uint32(3) // u32
		minRequiredVersion_81_minor := uint32(0) // u32
		_, _ = minRequiredVersion_81_major, minRequiredVersion_81_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_82_major := uint32(3) // u32
		minRequiredVersion_82_minor := uint32(1) // u32
		_, _ = minRequiredVersion_82_major, minRequiredVersion_82_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_83_major := uint32(3) // u32
		minRequiredVersion_83_minor := uint32(2) // u32
		_, _ = minRequiredVersion_83_major, minRequiredVersion_83_minor
	default:
		glErrorInvalidEnum_84_param := TexParameterv_80_target // GLenum
		return
		_ = glErrorInvalidEnum_84_param
	}
	switch TexParameterv_80_pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		TexParameterv_80_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_85_major := uint32(3) // u32
		minRequiredVersion_85_minor := uint32(0) // u32
		TexParameterv_80_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_85_major, minRequiredVersion_85_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_86_major := uint32(3) // u32
		minRequiredVersion_86_minor := uint32(1) // u32
		TexParameterv_80_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_86_major, minRequiredVersion_86_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_87_major := uint32(3) // u32
		minRequiredVersion_87_minor := uint32(2) // u32
		TexParameterv_80_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_87_major, minRequiredVersion_87_minor
	default:
		glErrorInvalidEnum_88_param := TexParameterv_80_pname // GLenum
		return
		_ = glErrorInvalidEnum_88_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_78_ext, minRequiredVersion_79_major, minRequiredVersion_79_minor, TexParameterv_80_target, TexParameterv_80_pname, TexParameterv_80_params
	return nil
}
func (ϟa *GlTexParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_89_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	minRequiredVersion_90_major := uint32(3)                            // u32
	minRequiredVersion_90_minor := uint32(0)                            // u32
	TexParameterv_91_target := ϟa.Target                                // GLenum
	TexParameterv_91_pname := ϟa.Pname                                  // GLenum
	TexParameterv_91_params := ϟa.Params                                // GLuintᶜᵖ
	switch TexParameterv_91_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_92_major := uint32(3) // u32
		minRequiredVersion_92_minor := uint32(0) // u32
		_, _ = minRequiredVersion_92_major, minRequiredVersion_92_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_93_major := uint32(3) // u32
		minRequiredVersion_93_minor := uint32(1) // u32
		_, _ = minRequiredVersion_93_major, minRequiredVersion_93_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_94_major := uint32(3) // u32
		minRequiredVersion_94_minor := uint32(2) // u32
		_, _ = minRequiredVersion_94_major, minRequiredVersion_94_minor
	default:
		glErrorInvalidEnum_95_param := TexParameterv_91_target // GLenum
		return
		_ = glErrorInvalidEnum_95_param
	}
	switch TexParameterv_91_pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		TexParameterv_91_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_96_major := uint32(3) // u32
		minRequiredVersion_96_minor := uint32(0) // u32
		TexParameterv_91_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_96_major, minRequiredVersion_96_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_97_major := uint32(3) // u32
		minRequiredVersion_97_minor := uint32(1) // u32
		TexParameterv_91_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_97_major, minRequiredVersion_97_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_98_major := uint32(3) // u32
		minRequiredVersion_98_minor := uint32(2) // u32
		TexParameterv_91_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_98_major, minRequiredVersion_98_minor
	default:
		glErrorInvalidEnum_99_param := TexParameterv_91_pname // GLenum
		return
		_ = glErrorInvalidEnum_99_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_89_ext, minRequiredVersion_90_major, minRequiredVersion_90_minor, TexParameterv_91_target, TexParameterv_91_pname, TexParameterv_91_params
	return nil
}
func (ϟa *GlTexStorage3DMultisampleOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_100_ext := ExtensionId_GL_OES_texture_storage_multisample_2d_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_100_ext
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_101_major := uint32(3) // u32
	minRequiredVersion_101_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_102_major := uint32(3) // u32
		minRequiredVersion_102_minor := uint32(2) // u32
		_, _ = minRequiredVersion_102_major, minRequiredVersion_102_minor
	default:
		glErrorInvalidEnum_103_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_103_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_101_major, minRequiredVersion_101_minor
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_104_major := uint32(3)                              // u32
	minRequiredVersion_104_minor := uint32(0)                              // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_106_msg := "No context bound" // string
		return
		_ = error_106_msg
	}
	GetContext_105_result := context // Contextʳ
	ctx := GetContext_105_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_104_major, minRequiredVersion_104_minor, q, context, GetContext_105_result, ctx
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_107_major := uint32(3) // u32
	minRequiredVersion_107_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_108_major := uint32(3) // u32
		minRequiredVersion_108_minor := uint32(2) // u32
		_, _ = minRequiredVersion_108_major, minRequiredVersion_108_minor
	default:
		glErrorInvalidEnum_109_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_109_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_107_major, minRequiredVersion_107_minor
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_110_major := uint32(3)                              // u32
	minRequiredVersion_110_minor := uint32(0)                              // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_112_msg := "No context bound" // string
		return
		_ = error_112_msg
	}
	GetContext_111_result := context // Contextʳ
	ctx := GetContext_111_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = &Query{}
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_110_major, minRequiredVersion_110_minor, q, context, GetContext_111_result, ctx
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_113_major := uint32(3) // u32
	minRequiredVersion_113_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_QUERY_RESULT, GLenum_GL_QUERY_RESULT_AVAILABLE:
	default:
		glErrorInvalidEnum_114_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_114_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_113_major, minRequiredVersion_113_minor
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_115_major := uint32(3) // u32
	minRequiredVersion_115_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_116_major := uint32(3) // u32
		minRequiredVersion_116_minor := uint32(2) // u32
		_, _ = minRequiredVersion_116_major, minRequiredVersion_116_minor
	default:
		glErrorInvalidEnum_117_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_117_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_CURRENT_QUERY:
	default:
		glErrorInvalidEnum_118_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_118_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_115_major, minRequiredVersion_115_minor
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_119_major := uint32(3)    // u32
	minRequiredVersion_119_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_121_msg := "No context bound" // string
		return
		_ = error_121_msg
	}
	GetContext_120_result := context // Contextʳ
	ctx := GetContext_120_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_119_major, minRequiredVersion_119_minor, context, GetContext_120_result, ctx
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_122_major := uint32(2) // u32
	minRequiredVersion_122_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_123_major := uint32(3) // u32
		minRequiredVersion_123_minor := uint32(0) // u32
		_, _ = minRequiredVersion_123_major, minRequiredVersion_123_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_124_major := uint32(3) // u32
		minRequiredVersion_124_minor := uint32(1) // u32
		_, _ = minRequiredVersion_124_major, minRequiredVersion_124_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_125_major := uint32(3) // u32
		minRequiredVersion_125_minor := uint32(2) // u32
		_, _ = minRequiredVersion_125_major, minRequiredVersion_125_minor
	default:
		glErrorInvalidEnum_126_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_126_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_128_msg := "No context bound" // string
		return
		_ = error_128_msg
	}
	GetContext_127_result := context // Contextʳ
	ctx := GetContext_127_result     // Contextʳ
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_122_major, minRequiredVersion_122_minor, context, GetContext_127_result, ctx
	return nil
}
func (ϟa *GlBindBufferBase) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_129_major := uint32(3) // u32
	minRequiredVersion_129_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_130_major := uint32(3) // u32
		minRequiredVersion_130_minor := uint32(1) // u32
		_, _ = minRequiredVersion_130_major, minRequiredVersion_130_minor
	default:
		glErrorInvalidEnum_131_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_131_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_129_major, minRequiredVersion_129_minor
	return nil
}
func (ϟa *GlBindBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_132_major := uint32(3) // u32
	minRequiredVersion_132_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_133_major := uint32(3) // u32
		minRequiredVersion_133_minor := uint32(1) // u32
		_, _ = minRequiredVersion_133_major, minRequiredVersion_133_minor
	default:
		glErrorInvalidEnum_134_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_134_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_132_major, minRequiredVersion_132_minor
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_135_major := uint32(2) // u32
	minRequiredVersion_135_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_136_major := uint32(3) // u32
		minRequiredVersion_136_minor := uint32(0) // u32
		_, _ = minRequiredVersion_136_major, minRequiredVersion_136_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_137_major := uint32(3) // u32
		minRequiredVersion_137_minor := uint32(1) // u32
		_, _ = minRequiredVersion_137_major, minRequiredVersion_137_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_138_major := uint32(3) // u32
		minRequiredVersion_138_minor := uint32(2) // u32
		_, _ = minRequiredVersion_138_major, minRequiredVersion_138_minor
	default:
		glErrorInvalidEnum_139_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_139_param
	}
	switch ϟa.Usage {
	case GLenum_GL_DYNAMIC_DRAW, GLenum_GL_STATIC_DRAW, GLenum_GL_STREAM_DRAW:
	case GLenum_GL_DYNAMIC_COPY, GLenum_GL_DYNAMIC_READ, GLenum_GL_STATIC_COPY, GLenum_GL_STATIC_READ, GLenum_GL_STREAM_COPY, GLenum_GL_STREAM_READ:
		minRequiredVersion_140_major := uint32(3) // u32
		minRequiredVersion_140_minor := uint32(0) // u32
		_, _ = minRequiredVersion_140_major, minRequiredVersion_140_minor
	default:
		glErrorInvalidEnum_141_param := ϟa.Usage // GLenum
		return
		_ = glErrorInvalidEnum_141_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_143_msg := "No context bound" // string
		return
		_ = error_143_msg
	}
	GetContext_142_result := context      // Contextʳ
	ctx := GetContext_142_result          // Contextʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_135_major, minRequiredVersion_135_minor, context, GetContext_142_result, ctx, id, b
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_144_major := uint32(2) // u32
	minRequiredVersion_144_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_145_major := uint32(3) // u32
		minRequiredVersion_145_minor := uint32(0) // u32
		_, _ = minRequiredVersion_145_major, minRequiredVersion_145_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_146_major := uint32(3) // u32
		minRequiredVersion_146_minor := uint32(1) // u32
		_, _ = minRequiredVersion_146_major, minRequiredVersion_146_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_147_major := uint32(3) // u32
		minRequiredVersion_147_minor := uint32(2) // u32
		_, _ = minRequiredVersion_147_major, minRequiredVersion_147_minor
	default:
		glErrorInvalidEnum_148_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_148_param
	}
	ϟa.Data.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_144_major, minRequiredVersion_144_minor
	return nil
}
func (ϟa *GlCopyBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_149_major := uint32(3) // u32
	minRequiredVersion_149_minor := uint32(0) // u32
	switch ϟa.ReadTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_150_major := uint32(3) // u32
		minRequiredVersion_150_minor := uint32(2) // u32
		_, _ = minRequiredVersion_150_major, minRequiredVersion_150_minor
	default:
		glErrorInvalidEnum_151_param := ϟa.ReadTarget // GLenum
		return
		_ = glErrorInvalidEnum_151_param
	}
	switch ϟa.WriteTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_152_major := uint32(3) // u32
		minRequiredVersion_152_minor := uint32(2) // u32
		_, _ = minRequiredVersion_152_major, minRequiredVersion_152_minor
	default:
		glErrorInvalidEnum_153_param := ϟa.WriteTarget // GLenum
		return
		_ = glErrorInvalidEnum_153_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_149_major, minRequiredVersion_149_minor
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_154_major := uint32(2)                              // u32
	minRequiredVersion_154_minor := uint32(0)                              // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_156_msg := "No context bound" // string
		return
		_ = error_156_msg
	}
	GetContext_155_result := context // Contextʳ
	ctx := GetContext_155_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Buffers, b.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_154_major, minRequiredVersion_154_minor, b, context, GetContext_155_result, ctx
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_157_major := uint32(2)                              // u32
	minRequiredVersion_157_minor := uint32(0)                              // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_159_msg := "No context bound" // string
		return
		_ = error_159_msg
	}
	GetContext_158_result := context // Contextʳ
	ctx := GetContext_158_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // BufferId
		ctx.Instances.Buffers[id] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
		b.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_157_major, minRequiredVersion_157_minor, b, context, GetContext_158_result, ctx
	return nil
}
func (ϟa *GlGetBufferParameteri64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_160_major := uint32(3) // u32
	minRequiredVersion_160_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_161_major := uint32(3) // u32
		minRequiredVersion_161_minor := uint32(2) // u32
		_, _ = minRequiredVersion_161_major, minRequiredVersion_161_minor
	default:
		glErrorInvalidEnum_162_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_162_param
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET, GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	default:
		glErrorInvalidEnum_163_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_163_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_160_major, minRequiredVersion_160_minor
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_164_major := uint32(2) // u32
	minRequiredVersion_164_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_165_major := uint32(3) // u32
		minRequiredVersion_165_minor := uint32(0) // u32
		_, _ = minRequiredVersion_165_major, minRequiredVersion_165_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_166_major := uint32(3) // u32
		minRequiredVersion_166_minor := uint32(2) // u32
		_, _ = minRequiredVersion_166_major, minRequiredVersion_166_minor
	default:
		glErrorInvalidEnum_167_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_167_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET:
		minRequiredVersion_168_major := uint32(3) // u32
		minRequiredVersion_168_minor := uint32(0) // u32
		_, _ = minRequiredVersion_168_major, minRequiredVersion_168_minor
	default:
		glErrorInvalidEnum_169_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_169_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_171_msg := "No context bound" // string
		return
		_ = error_171_msg
	}
	GetContext_170_result := context      // Contextʳ
	ctx := GetContext_170_result          // Contextʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_164_major, minRequiredVersion_164_minor, context, GetContext_170_result, ctx, id, b
	return nil
}
func (ϟa *GlGetBufferPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_172_major := uint32(3) // u32
	minRequiredVersion_172_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_173_major := uint32(3) // u32
		minRequiredVersion_173_minor := uint32(2) // u32
		_, _ = minRequiredVersion_173_major, minRequiredVersion_173_minor
	default:
		glErrorInvalidEnum_174_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_174_param
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_MAP_POINTER:
	default:
		glErrorInvalidEnum_175_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_175_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_172_major, minRequiredVersion_172_minor
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_176_major := uint32(2)    // u32
	minRequiredVersion_176_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_178_msg := "No context bound" // string
		return
		_ = error_178_msg
	}
	GetContext_177_result := context // Contextʳ
	ctx := GetContext_177_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Buffers.Contains(ϟa.Buffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_176_major, minRequiredVersion_176_minor, context, GetContext_177_result, ctx
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_179_major := uint32(3) // u32
	minRequiredVersion_179_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_180_major := uint32(3) // u32
		minRequiredVersion_180_minor := uint32(1) // u32
		_, _ = minRequiredVersion_180_major, minRequiredVersion_180_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_181_major := uint32(3) // u32
		minRequiredVersion_181_minor := uint32(2) // u32
		_, _ = minRequiredVersion_181_major, minRequiredVersion_181_minor
	default:
		glErrorInvalidEnum_182_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_182_param
	}
	supportsBits_183_seenBits := ϟa.Access                                                                                                                                                                                                                                      // GLbitfield
	supportsBits_183_validBits := (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT) | ((GLbitfield_GL_MAP_READ_BIT) | ((GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT) | (GLbitfield_GL_MAP_WRITE_BIT))))) // GLbitfield
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
		error_185_msg := "No context bound" // string
		return
		_ = error_185_msg
	}
	GetContext_184_result := context                                // Contextʳ
	ctx := GetContext_184_result                                    // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_179_major, minRequiredVersion_179_minor, supportsBits_183_seenBits, supportsBits_183_validBits, context, GetContext_184_result, ctx, b, ptr
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_186_major := uint32(3) // u32
	minRequiredVersion_186_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_187_major := uint32(3) // u32
		minRequiredVersion_187_minor := uint32(1) // u32
		_, _ = minRequiredVersion_187_major, minRequiredVersion_187_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_188_major := uint32(3) // u32
		minRequiredVersion_188_minor := uint32(2) // u32
		_, _ = minRequiredVersion_188_major, minRequiredVersion_188_minor
	default:
		glErrorInvalidEnum_189_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_189_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_191_msg := "No context bound" // string
		return
		_ = error_191_msg
	}
	GetContext_190_result := context                                // Contextʳ
	ctx := GetContext_190_result                                    // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	b.Data.Slice(uint64(b.MappingOffset), uint64((b.MappingOffset)+(int32(b.MappingData.Count))), ϟs).Copy(b.MappingData, ϟa, ϟs, ϟd, ϟl, ϟb)
	externs{ϟa, ϟs, ϟd, ϟl, ϟb}.unmapMemory(b.MappingData)
	b.MappingOffset = int32(0)
	b.MappingData = U8ˢ{}
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _ = minRequiredVersion_186_major, minRequiredVersion_186_minor, context, GetContext_190_result, ctx, b
	return nil
}
func (ϟa *GlDebugMessageCallback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_192_major := uint32(3) // u32
	minRequiredVersion_192_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_192_major, minRequiredVersion_192_minor
	return nil
}
func (ϟa *GlDebugMessageControl) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_193_major := uint32(3) // u32
	minRequiredVersion_193_minor := uint32(2) // u32
	switch ϟa.Source {
	case GLenum_GL_DEBUG_SOURCE_API, GLenum_GL_DEBUG_SOURCE_APPLICATION, GLenum_GL_DEBUG_SOURCE_OTHER, GLenum_GL_DEBUG_SOURCE_SHADER_COMPILER, GLenum_GL_DEBUG_SOURCE_THIRD_PARTY, GLenum_GL_DEBUG_SOURCE_WINDOW_SYSTEM, GLenum_GL_DONT_CARE:
	default:
		glErrorInvalidEnum_194_param := ϟa.Source // GLenum
		return
		_ = glErrorInvalidEnum_194_param
	}
	switch ϟa.Type {
	case GLenum_GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR, GLenum_GL_DEBUG_TYPE_ERROR, GLenum_GL_DEBUG_TYPE_MARKER, GLenum_GL_DEBUG_TYPE_OTHER, GLenum_GL_DEBUG_TYPE_PERFORMANCE, GLenum_GL_DEBUG_TYPE_POP_GROUP, GLenum_GL_DEBUG_TYPE_PORTABILITY, GLenum_GL_DEBUG_TYPE_PUSH_GROUP, GLenum_GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR, GLenum_GL_DONT_CARE:
	default:
		glErrorInvalidEnum_195_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_195_param
	}
	switch ϟa.Severity {
	case GLenum_GL_DEBUG_SEVERITY_HIGH, GLenum_GL_DEBUG_SEVERITY_LOW, GLenum_GL_DEBUG_SEVERITY_MEDIUM, GLenum_GL_DEBUG_SEVERITY_NOTIFICATION, GLenum_GL_DONT_CARE:
	default:
		glErrorInvalidEnum_196_param := ϟa.Severity // GLenum
		return
		_ = glErrorInvalidEnum_196_param
	}
	ϟa.Ids.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_193_major, minRequiredVersion_193_minor
	return nil
}
func (ϟa *GlDebugMessageInsert) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_197_major := uint32(3) // u32
	minRequiredVersion_197_minor := uint32(2) // u32
	switch ϟa.Source {
	case GLenum_GL_DEBUG_SOURCE_APPLICATION, GLenum_GL_DEBUG_SOURCE_THIRD_PARTY:
	default:
		glErrorInvalidEnum_198_param := ϟa.Source // GLenum
		return
		_ = glErrorInvalidEnum_198_param
	}
	switch ϟa.Type {
	case GLenum_GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR, GLenum_GL_DEBUG_TYPE_ERROR, GLenum_GL_DEBUG_TYPE_MARKER, GLenum_GL_DEBUG_TYPE_OTHER, GLenum_GL_DEBUG_TYPE_PERFORMANCE, GLenum_GL_DEBUG_TYPE_POP_GROUP, GLenum_GL_DEBUG_TYPE_PORTABILITY, GLenum_GL_DEBUG_TYPE_PUSH_GROUP, GLenum_GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR:
	default:
		glErrorInvalidEnum_199_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_199_param
	}
	switch ϟa.Severity {
	case GLenum_GL_DEBUG_SEVERITY_HIGH, GLenum_GL_DEBUG_SEVERITY_LOW, GLenum_GL_DEBUG_SEVERITY_MEDIUM, GLenum_GL_DEBUG_SEVERITY_NOTIFICATION:
	default:
		glErrorInvalidEnum_200_param := ϟa.Severity // GLenum
		return
		_ = glErrorInvalidEnum_200_param
	}
	readString_201_length := ϟa.Length  // GLsizei
	readString_201_buffer := ϟa.Message // GLcharᶜᵖ
	if (readString_201_buffer) != (GLcharᶜᵖ{}) {
		if (readString_201_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_201_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_201_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_201_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_197_major, minRequiredVersion_197_minor, readString_201_length, readString_201_buffer
	return nil
}
func (ϟa *GlGetDebugMessageLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_202_major := uint32(3)          // u32
	minRequiredVersion_202_minor := uint32(2)          // u32
	GetDebugMessageLog_203_count := ϟa.Count           // GLuint
	GetDebugMessageLog_203_bufSize := ϟa.BufSize       // GLsizei
	GetDebugMessageLog_203_sources := ϟa.Sources       // GLenumᵖ
	GetDebugMessageLog_203_types := ϟa.Types           // GLenumᵖ
	GetDebugMessageLog_203_ids := ϟa.Ids               // GLuintᵖ
	GetDebugMessageLog_203_severities := ϟa.Severities // GLenumᵖ
	GetDebugMessageLog_203_lengths := ϟa.Lengths       // GLsizeiᵖ
	GetDebugMessageLog_203_messageLog := ϟa.MessageLog // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLuint(ϟa.Result) // GLuint
	if (GetDebugMessageLog_203_sources) != (GLenumᵖ{}) {
		GetDebugMessageLog_203_sources.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_203_types) != (GLenumᵖ{}) {
		GetDebugMessageLog_203_types.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_203_ids) != (GLuintᵖ{}) {
		GetDebugMessageLog_203_ids.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_203_severities) != (GLenumᵖ{}) {
		GetDebugMessageLog_203_severities.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (GetDebugMessageLog_203_lengths) != (GLsizeiᵖ{}) {
		GetDebugMessageLog_203_lengths.Slice(uint64(GLuint(uint32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (l) > (GLuint(uint32(0))) {
		GetDebugMessageLog_203_messageLog.Slice(uint64(GLsizei(int32(0))), uint64(GetDebugMessageLog_203_bufSize), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	GetDebugMessageLog_203_result := l // GLuint
	ϟa.Result = GetDebugMessageLog_203_result
	_, _, _, _, _, _, _, _, _, _, _, _ = minRequiredVersion_202_major, minRequiredVersion_202_minor, GetDebugMessageLog_203_count, GetDebugMessageLog_203_bufSize, GetDebugMessageLog_203_sources, GetDebugMessageLog_203_types, GetDebugMessageLog_203_ids, GetDebugMessageLog_203_severities, GetDebugMessageLog_203_lengths, GetDebugMessageLog_203_messageLog, l, GetDebugMessageLog_203_result
	return nil
}
func (ϟa *GlGetObjectLabel) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_204_major := uint32(3) // u32
	minRequiredVersion_204_minor := uint32(2) // u32
	switch ϟa.Identifier {
	case GLenum_GL_BUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_PROGRAM, GLenum_GL_PROGRAM_PIPELINE, GLenum_GL_QUERY, GLenum_GL_RENDERBUFFER, GLenum_GL_SAMPLER, GLenum_GL_SHADER, GLenum_GL_TEXTURE, GLenum_GL_TRANSFORM_FEEDBACK, GLenum_GL_VERTEX_ARRAY:
	default:
		glErrorInvalidEnum_205_param := ϟa.Identifier // GLenum
		return
		_ = glErrorInvalidEnum_205_param
	}
	writeString_206_buffer_size := ϟa.BufSize         // GLsizei
	writeString_206_buffer_bytes_written := ϟa.Length // GLsizeiᵖ
	writeString_206_buffer := ϟa.Label                // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_206_buffer) != (GLcharᵖ{})) && ((writeString_206_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_206_buffer_size // GLsizei
		if (writeString_206_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_206_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_206_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_206_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _ = minRequiredVersion_204_major, minRequiredVersion_204_minor, writeString_206_buffer_size, writeString_206_buffer_bytes_written, writeString_206_buffer
	return nil
}
func (ϟa *GlGetObjectPtrLabel) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_207_major := uint32(3)         // u32
	minRequiredVersion_207_minor := uint32(2)         // u32
	writeString_208_buffer_size := ϟa.BufSize         // GLsizei
	writeString_208_buffer_bytes_written := ϟa.Length // GLsizeiᵖ
	writeString_208_buffer := ϟa.Label                // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_208_buffer) != (GLcharᵖ{})) && ((writeString_208_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_208_buffer_size // GLsizei
		if (writeString_208_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_208_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_208_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_208_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _ = minRequiredVersion_207_major, minRequiredVersion_207_minor, writeString_208_buffer_size, writeString_208_buffer_bytes_written, writeString_208_buffer
	return nil
}
func (ϟa *GlGetPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_209_major := uint32(3) // u32
	minRequiredVersion_209_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_DEBUG_CALLBACK_FUNCTION, GLenum_GL_DEBUG_CALLBACK_USER_PARAM:
	default:
		glErrorInvalidEnum_210_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_210_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_209_major, minRequiredVersion_209_minor
	return nil
}
func (ϟa *GlObjectLabel) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_211_major := uint32(3) // u32
	minRequiredVersion_211_minor := uint32(2) // u32
	switch ϟa.Identifier {
	case GLenum_GL_BUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_PROGRAM, GLenum_GL_PROGRAM_PIPELINE, GLenum_GL_QUERY, GLenum_GL_RENDERBUFFER, GLenum_GL_SAMPLER, GLenum_GL_SHADER, GLenum_GL_TEXTURE, GLenum_GL_TRANSFORM_FEEDBACK, GLenum_GL_VERTEX_ARRAY:
	default:
		glErrorInvalidEnum_212_param := ϟa.Identifier // GLenum
		return
		_ = glErrorInvalidEnum_212_param
	}
	readString_213_length := ϟa.Length // GLsizei
	readString_213_buffer := ϟa.Label  // GLcharᶜᵖ
	if (readString_213_buffer) != (GLcharᶜᵖ{}) {
		if (readString_213_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_213_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_213_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_213_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_211_major, minRequiredVersion_211_minor, readString_213_length, readString_213_buffer
	return nil
}
func (ϟa *GlObjectPtrLabel) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_214_major := uint32(3) // u32
	minRequiredVersion_214_minor := uint32(2) // u32
	readString_215_length := ϟa.Length        // GLsizei
	readString_215_buffer := ϟa.Label         // GLcharᶜᵖ
	if (readString_215_buffer) != (GLcharᶜᵖ{}) {
		if (readString_215_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_215_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_215_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_215_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_214_major, minRequiredVersion_214_minor, readString_215_length, readString_215_buffer
	return nil
}
func (ϟa *GlPopDebugGroup) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_216_major := uint32(3) // u32
	minRequiredVersion_216_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_216_major, minRequiredVersion_216_minor
	return nil
}
func (ϟa *GlPushDebugGroup) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_217_major := uint32(3) // u32
	minRequiredVersion_217_minor := uint32(2) // u32
	switch ϟa.Source {
	case GLenum_GL_DEBUG_SOURCE_APPLICATION, GLenum_GL_DEBUG_SOURCE_THIRD_PARTY:
	default:
		glErrorInvalidEnum_218_param := ϟa.Source // GLenum
		return
		_ = glErrorInvalidEnum_218_param
	}
	readString_219_length := ϟa.Length  // GLsizei
	readString_219_buffer := ϟa.Message // GLcharᶜᵖ
	if (readString_219_buffer) != (GLcharᶜᵖ{}) {
		if (readString_219_length) < (GLsizei(int32(0))) {
			msg := strings.TrimRight(string(Charᵖ(readString_219_buffer).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
			_ = msg
		} else {
			readString_219_buffer.Slice(uint64(GLsizei(int32(0))), uint64(readString_219_length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_217_major, minRequiredVersion_217_minor, readString_219_length, readString_219_buffer
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_220_major := uint32(2) // u32
	minRequiredVersion_220_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_221_major := uint32(3) // u32
		minRequiredVersion_221_minor := uint32(2) // u32
		_, _ = minRequiredVersion_221_major, minRequiredVersion_221_minor
	default:
		glErrorInvalidEnum_222_param := ϟa.DrawMode // GLenum
		return
		_ = glErrorInvalidEnum_222_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_224_msg := "No context bound" // string
		return
		_ = error_224_msg
	}
	GetContext_223_result := context                                                // Contextʳ
	ctx := GetContext_223_result                                                    // Contextʳ
	last_index := (uint32(ϟa.FirstIndex)) + ((uint32(ϟa.IndexCount)) - (uint32(1))) // u32
	ReadVertexArrays_225_ctx := ctx                                                 // Contextʳ
	ReadVertexArrays_225_first_index := uint32(ϟa.FirstIndex)                       // u32
	ReadVertexArrays_225_last_index := last_index                                   // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_225_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_225_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
		if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
			vertexAttribTypeSize_226_t := arr.Type // GLenum
			vertexAttribTypeSize_226_result := func() (result uint32) {
				switch vertexAttribTypeSize_226_t {
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
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_226_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_226_result) * (arr.Size) // u32
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
			for v := uint32(ReadVertexArrays_225_first_index); v < (ReadVertexArrays_225_last_index)+(uint32(1)); v++ {
				offset := (elstride) * (v) // u32
				arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_226_t, vertexAttribTypeSize_226_result, elsize, elstride
		}
		_ = arr
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_220_major, minRequiredVersion_220_minor, context, GetContext_223_result, ctx, last_index, ReadVertexArrays_225_ctx, ReadVertexArrays_225_first_index, ReadVertexArrays_225_last_index
	return nil
}
func (ϟa *GlDrawArraysIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_227_major := uint32(3) // u32
	minRequiredVersion_227_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_228_major := uint32(3) // u32
		minRequiredVersion_228_minor := uint32(2) // u32
		_, _ = minRequiredVersion_228_major, minRequiredVersion_228_minor
	default:
		glErrorInvalidEnum_229_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_229_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_227_major, minRequiredVersion_227_minor
	return nil
}
func (ϟa *GlDrawArraysInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_230_major := uint32(3) // u32
	minRequiredVersion_230_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_231_major := uint32(3) // u32
		minRequiredVersion_231_minor := uint32(2) // u32
		_, _ = minRequiredVersion_231_major, minRequiredVersion_231_minor
	default:
		glErrorInvalidEnum_232_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_232_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_230_major, minRequiredVersion_230_minor
	return nil
}
func (ϟa *GlDrawBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_233_major := uint32(3) // u32
	minRequiredVersion_233_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_233_major, minRequiredVersion_233_minor
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_234_major := uint32(2) // u32
	minRequiredVersion_234_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_235_major := uint32(3) // u32
		minRequiredVersion_235_minor := uint32(2) // u32
		_, _ = minRequiredVersion_235_major, minRequiredVersion_235_minor
	default:
		glErrorInvalidEnum_236_param := ϟa.DrawMode // GLenum
		return
		_ = glErrorInvalidEnum_236_param
	}
	switch ϟa.IndicesType {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_UNSIGNED_INT:
		minRequiredVersion_237_major := uint32(3) // u32
		minRequiredVersion_237_minor := uint32(0) // u32
		_, _ = minRequiredVersion_237_major, minRequiredVersion_237_minor
	default:
		glErrorInvalidEnum_238_param := ϟa.IndicesType // GLenum
		return
		_ = glErrorInvalidEnum_238_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_240_msg := "No context bound" // string
		return
		_ = error_240_msg
	}
	GetContext_239_result := context                           // Contextʳ
	ctx := GetContext_239_result                               // Contextʳ
	count := uint32(ϟa.ElementCount)                           // u32
	id := ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	if (id) != (BufferId(uint32(0))) {
		index_data := ctx.Instances.Buffers.Get(id).Data                                                           // U8ˢ
		offset := uint32(uint64(ϟa.Indices.Address))                                                               // u32
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count)  // u32
		ReadVertexArrays_241_ctx := ctx                                                                            // Contextʳ
		ReadVertexArrays_241_first_index := first                                                                  // u32
		ReadVertexArrays_241_last_index := last                                                                    // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_241_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_241_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_242_t := arr.Type // GLenum
				vertexAttribTypeSize_242_result := func() (result uint32) {
					switch vertexAttribTypeSize_242_t {
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
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_242_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_242_result) * (arr.Size) // u32
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
				for v := uint32(ReadVertexArrays_241_first_index); v < (ReadVertexArrays_241_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_242_t, vertexAttribTypeSize_242_result, elsize, elstride
			}
			_ = arr
		}
		_, _, _, _, _, _, _ = index_data, offset, first, last, ReadVertexArrays_241_ctx, ReadVertexArrays_241_first_index, ReadVertexArrays_241_last_index
	} else {
		index_data := U8ᵖ(ϟa.Indices)                                                               // U8ᵖ
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(index_data, ϟa.IndicesType, uint32(0), count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(index_data, ϟa.IndicesType, uint32(0), count)  // u32
		ReadVertexArrays_243_ctx := ctx                                                             // Contextʳ
		ReadVertexArrays_243_first_index := first                                                   // u32
		ReadVertexArrays_243_last_index := last                                                     // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_243_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_243_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_244_t := arr.Type // GLenum
				vertexAttribTypeSize_244_result := func() (result uint32) {
					switch vertexAttribTypeSize_244_t {
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
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_244_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_244_result) * (arr.Size) // u32
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
				for v := uint32(ReadVertexArrays_243_first_index); v < (ReadVertexArrays_243_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_244_t, vertexAttribTypeSize_244_result, elsize, elstride
			}
			_ = arr
		}
		IndexSize_245_indices_type := ϟa.IndicesType // GLenum
		IndexSize_245_result := func() (result uint32) {
			switch IndexSize_245_indices_type {
			case GLenum_GL_UNSIGNED_BYTE:
				return uint32(1)
			case GLenum_GL_UNSIGNED_SHORT:
				return uint32(2)
			case GLenum_GL_UNSIGNED_INT:
				return uint32(4)
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_245_indices_type, ϟa))
				return result
			}
		}() // u32
		index_data.Slice(uint64(uint32(0)), uint64((uint32(ϟa.ElementCount))*(IndexSize_245_result)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_243_ctx, ReadVertexArrays_243_first_index, ReadVertexArrays_243_last_index, IndexSize_245_indices_type, IndexSize_245_result
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_234_major, minRequiredVersion_234_minor, context, GetContext_239_result, ctx, count, id
	return nil
}
func (ϟa *GlDrawElementsBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_246_major := uint32(3) // u32
	minRequiredVersion_246_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
	default:
		glErrorInvalidEnum_247_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_247_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_248_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_248_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_246_major, minRequiredVersion_246_minor
	return nil
}
func (ϟa *GlDrawElementsIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_249_major := uint32(3) // u32
	minRequiredVersion_249_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_250_major := uint32(3) // u32
		minRequiredVersion_250_minor := uint32(2) // u32
		_, _ = minRequiredVersion_250_major, minRequiredVersion_250_minor
	default:
		glErrorInvalidEnum_251_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_251_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_252_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_252_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_249_major, minRequiredVersion_249_minor
	return nil
}
func (ϟa *GlDrawElementsInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_253_major := uint32(3) // u32
	minRequiredVersion_253_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_254_major := uint32(3) // u32
		minRequiredVersion_254_minor := uint32(2) // u32
		_, _ = minRequiredVersion_254_major, minRequiredVersion_254_minor
	default:
		glErrorInvalidEnum_255_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_255_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_256_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_256_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_253_major, minRequiredVersion_253_minor
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_257_major := uint32(3) // u32
	minRequiredVersion_257_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
	default:
		glErrorInvalidEnum_258_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_258_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_259_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_259_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_257_major, minRequiredVersion_257_minor
	return nil
}
func (ϟa *GlDrawRangeElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_260_major := uint32(3) // u32
	minRequiredVersion_260_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_261_major := uint32(3) // u32
		minRequiredVersion_261_minor := uint32(2) // u32
		_, _ = minRequiredVersion_261_major, minRequiredVersion_261_minor
	default:
		glErrorInvalidEnum_262_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_262_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_263_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_263_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_260_major, minRequiredVersion_260_minor
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_264_major := uint32(3) // u32
	minRequiredVersion_264_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
	default:
		glErrorInvalidEnum_265_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_265_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_266_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_266_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_264_major, minRequiredVersion_264_minor
	return nil
}
func (ϟa *GlPatchParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_267_major := uint32(3) // u32
	minRequiredVersion_267_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_PATCH_VERTICES:
	default:
		glErrorInvalidEnum_268_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_268_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_267_major, minRequiredVersion_267_minor
	return nil
}
func (ϟa *GlPrimitiveBoundingBox) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_269_major := uint32(3) // u32
	minRequiredVersion_269_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_269_major, minRequiredVersion_269_minor
	return nil
}
func (ϟa *GlActiveShaderProgramEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_270_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_270_ext
	return nil
}
func (ϟa *GlAlphaFuncQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_271_ext := ExtensionId_GL_QCOM_alpha_test // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_271_ext
	return nil
}
func (ϟa *GlApplyFramebufferAttachmentCMAAINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_272_ext := ExtensionId_GL_INTEL_framebuffer_CMAA // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_272_ext
	return nil
}
func (ϟa *GlBeginConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_273_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_273_ext
	return nil
}
func (ϟa *GlBeginPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_274_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_274_ext
	return nil
}
func (ϟa *GlBeginPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_275_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_275_ext
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_276_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_277_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_276_ext, requiresExtension_277_ext
	return nil
}
func (ϟa *GlBindProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_278_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_278_ext
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_279_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_281_msg := "No context bound" // string
		return
		_ = error_281_msg
	}
	GetContext_280_result := context // Contextʳ
	ctx := GetContext_280_result     // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = &VertexArray{}
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = requiresExtension_279_ext, context, GetContext_280_result, ctx
	return nil
}
func (ϟa *GlBlendBarrierNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_282_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_282_ext
	return nil
}
func (ϟa *GlBlendEquationSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_283_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_283_ext
	return nil
}
func (ϟa *GlBlendEquationiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_284_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_284_ext
	return nil
}
func (ϟa *GlBlendFuncSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_285_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_285_ext
	return nil
}
func (ϟa *GlBlendFunciOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_286_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_286_ext
	return nil
}
func (ϟa *GlBlendParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_287_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_287_ext
	return nil
}
func (ϟa *GlBlitFramebufferANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_288_ext := ExtensionId_GL_ANGLE_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_288_ext
	return nil
}
func (ϟa *GlBlitFramebufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_289_ext := ExtensionId_GL_NV_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_289_ext
	return nil
}
func (ϟa *GlBufferStorageEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_290_ext := ExtensionId_GL_EXT_buffer_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_290_ext
	return nil
}
func (ϟa *GlClientWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_291_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_291_ext
	return nil
}
func (ϟa *GlColorMaskiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_292_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_292_ext
	return nil
}
func (ϟa *GlCompressedTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_293_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_293_ext
	return nil
}
func (ϟa *GlCompressedTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_294_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_294_ext
	return nil
}
func (ϟa *GlCopyBufferSubDataNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_295_ext := ExtensionId_GL_NV_copy_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_295_ext
	return nil
}
func (ϟa *GlCopyImageSubDataOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_296_ext := ExtensionId_GL_OES_copy_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_296_ext
	return nil
}
func (ϟa *GlCopyPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_297_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_297_ext
	return nil
}
func (ϟa *GlCopyTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_298_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_298_ext
	return nil
}
func (ϟa *GlCopyTextureLevelsAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_299_ext := ExtensionId_GL_APPLE_copy_texture_levels // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_299_ext
	return nil
}
func (ϟa *GlCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_300_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_300_ext
	return nil
}
func (ϟa *GlCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_301_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_301_ext
	return nil
}
func (ϟa *GlCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_302_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_302_ext
	return nil
}
func (ϟa *GlCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_303_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_303_ext
	return nil
}
func (ϟa *GlCoverageMaskNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_304_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_304_ext
	return nil
}
func (ϟa *GlCoverageModulationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_305_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_305_ext
	return nil
}
func (ϟa *GlCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_306_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_306_ext
	return nil
}
func (ϟa *GlCoverageOperationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_307_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_307_ext
	return nil
}
func (ϟa *GlCreatePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_308_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_308_ext
	return nil
}
func (ϟa *GlCreateShaderProgramvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_309_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_309_ext
	return nil
}
func (ϟa *GlDeleteFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_310_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_310_ext
	return nil
}
func (ϟa *GlDeletePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_311_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_311_ext
	return nil
}
func (ϟa *GlDeletePerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_312_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_312_ext
	return nil
}
func (ϟa *GlDeletePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_313_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_313_ext
	return nil
}
func (ϟa *GlDeleteProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_314_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_314_ext
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_315_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_316_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_318_msg := "No context bound" // string
		return
		_ = error_318_msg
	}
	GetContext_317_result := context // Contextʳ
	ctx := GetContext_317_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_315_ext, requiresExtension_316_ext, q, context, GetContext_317_result, ctx
	return nil
}
func (ϟa *GlDeleteSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_319_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_319_ext
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_320_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_322_msg := "No context bound" // string
		return
		_ = error_322_msg
	}
	GetContext_321_result := context                                      // Contextʳ
	ctx := GetContext_321_result                                          // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = requiresExtension_320_ext, context, GetContext_321_result, ctx, a
	return nil
}
func (ϟa *GlDepthRangeArrayfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_323_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_323_ext
	return nil
}
func (ϟa *GlDepthRangeIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_324_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_324_ext
	return nil
}
func (ϟa *GlDisableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_325_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_325_ext
	return nil
}
func (ϟa *GlDisableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_326_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_326_ext
	return nil
}
func (ϟa *GlDisableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_327_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_327_ext
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_328_ext := ExtensionId_GL_EXT_discard_framebuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_328_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_329_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_329_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_330_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_330_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_331_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_332_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_331_ext, requiresExtension_332_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_333_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_333_ext
	return nil
}
func (ϟa *GlDrawBuffersEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_334_ext := ExtensionId_GL_EXT_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_334_ext
	return nil
}
func (ϟa *GlDrawBuffersIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_335_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_335_ext
	return nil
}
func (ϟa *GlDrawBuffersNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_336_ext := ExtensionId_GL_NV_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_336_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_337_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_337_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_338_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_338_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_339_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_339_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_340_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_340_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_341_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_341_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_342_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_342_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_343_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_343_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_344_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_345_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_344_ext, requiresExtension_345_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_346_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_346_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_347_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_347_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_348_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_348_ext
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_349_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_349_ext
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_350_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_350_ext
	return nil
}
func (ϟa *GlEnableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_351_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_351_ext
	return nil
}
func (ϟa *GlEnableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_352_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_352_ext
	return nil
}
func (ϟa *GlEnableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_353_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_353_ext
	return nil
}
func (ϟa *GlEndConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_354_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_354_ext
	return nil
}
func (ϟa *GlEndPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_355_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_355_ext
	return nil
}
func (ϟa *GlEndPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_356_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_356_ext
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_357_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_358_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_357_ext, requiresExtension_358_ext
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_359_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_359_ext
	return nil
}
func (ϟa *GlExtGetBufferPointervQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_360_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_360_ext
	return nil
}
func (ϟa *GlExtGetBuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_361_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_361_ext
	return nil
}
func (ϟa *GlExtGetFramebuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_362_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_362_ext
	return nil
}
func (ϟa *GlExtGetProgramBinarySourceQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_363_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_363_ext
	return nil
}
func (ϟa *GlExtGetProgramsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_364_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_364_ext
	return nil
}
func (ϟa *GlExtGetRenderbuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_365_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_365_ext
	return nil
}
func (ϟa *GlExtGetShadersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_366_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_366_ext
	return nil
}
func (ϟa *GlExtGetTexLevelParameterivQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_367_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_367_ext
	return nil
}
func (ϟa *GlExtGetTexSubImageQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_368_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_368_ext
	return nil
}
func (ϟa *GlExtGetTexturesQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_369_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_369_ext
	return nil
}
func (ϟa *GlExtIsProgramBinaryQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_370_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_370_ext
	return nil
}
func (ϟa *GlExtTexObjectStateOverrideiQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_371_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_371_ext
	return nil
}
func (ϟa *GlFenceSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_372_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_372_ext
	return nil
}
func (ϟa *GlFinishFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_373_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_373_ext
	return nil
}
func (ϟa *GlFlushMappedBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_374_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_374_ext
	return nil
}
func (ϟa *GlFragmentCoverageColorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_375_ext := ExtensionId_GL_NV_fragment_coverage_to_color // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_375_ext
	return nil
}
func (ϟa *GlFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_376_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_376_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_377_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_377_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_378_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_378_ext
	return nil
}
func (ϟa *GlFramebufferTexture3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_379_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_379_ext
	return nil
}
func (ϟa *GlFramebufferTextureMultiviewOVR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_380_ext := ExtensionId_GL_OVR_multiview // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_380_ext
	return nil
}
func (ϟa *GlFramebufferTextureOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_381_ext := ExtensionId_GL_OES_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_381_ext
	return nil
}
func (ϟa *GlGenFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_382_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_382_ext
	return nil
}
func (ϟa *GlGenPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_383_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_383_ext
	return nil
}
func (ϟa *GlGenPerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_384_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_384_ext
	return nil
}
func (ϟa *GlGenProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_385_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_385_ext
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_386_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_387_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_389_msg := "No context bound" // string
		return
		_ = error_389_msg
	}
	GetContext_388_result := context // Contextʳ
	ctx := GetContext_388_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = &Query{}
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = requiresExtension_386_ext, requiresExtension_387_ext, q, context, GetContext_388_result, ctx
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_390_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_392_msg := "No context bound" // string
		return
		_ = error_392_msg
	}
	GetContext_391_result := context // Contextʳ
	ctx := GetContext_391_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = &VertexArray{}
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _ = requiresExtension_390_ext, a, context, GetContext_391_result, ctx
	return nil
}
func (ϟa *GlGetBufferPointervOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_393_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_393_ext
	return nil
}
func (ϟa *GlGetCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_394_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_394_ext
	return nil
}
func (ϟa *GlGetDriverControlStringQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_395_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_395_ext
	return nil
}
func (ϟa *GlGetDriverControlsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_396_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_396_ext
	return nil
}
func (ϟa *GlGetFenceivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_397_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_397_ext
	return nil
}
func (ϟa *GlGetFirstPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_398_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_398_ext
	return nil
}
func (ϟa *GlGetFloati_vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_399_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_399_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_400_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_400_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_401_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_401_ext
	return nil
}
func (ϟa *GlGetImageHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_402_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_402_ext
	return nil
}
func (ϟa *GlGetInteger64vAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_403_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_403_ext
	return nil
}
func (ϟa *GlGetIntegeri_vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_404_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_404_ext
	return nil
}
func (ϟa *GlGetInternalformatSampleivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_405_ext := ExtensionId_GL_NV_internalformat_sample_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_405_ext
	return nil
}
func (ϟa *GlGetNextPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_406_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_406_ext
	return nil
}
func (ϟa *GlGetObjectLabelEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_407_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_407_ext
	return nil
}
func (ϟa *GlGetPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_408_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_408_ext
	return nil
}
func (ϟa *GlGetPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_409_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_409_ext
	return nil
}
func (ϟa *GlGetPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_410_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_410_ext
	return nil
}
func (ϟa *GlGetPathLengthNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_411_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_411_ext
	return nil
}
func (ϟa *GlGetPathMetricRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_412_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_412_ext
	return nil
}
func (ϟa *GlGetPathMetricsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_413_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_413_ext
	return nil
}
func (ϟa *GlGetPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_414_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_414_ext
	return nil
}
func (ϟa *GlGetPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_415_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_415_ext
	return nil
}
func (ϟa *GlGetPathSpacingNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_416_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_416_ext
	return nil
}
func (ϟa *GlGetPerfCounterInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_417_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_417_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterDataAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_418_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_418_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterInfoAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_419_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_419_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_420_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_420_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_421_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_421_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_422_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_422_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_423_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_423_ext
	return nil
}
func (ϟa *GlGetPerfQueryDataINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_424_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_424_ext
	return nil
}
func (ϟa *GlGetPerfQueryIdByNameINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_425_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_425_ext
	return nil
}
func (ϟa *GlGetPerfQueryInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_426_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_426_ext
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_427_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_427_ext, l
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLogEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_428_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_428_ext
	return nil
}
func (ϟa *GlGetProgramPipelineivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_429_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_429_ext
	return nil
}
func (ϟa *GlGetProgramResourcefvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_430_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_430_ext
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_431_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_431_ext
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_432_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_432_ext
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_433_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_433_ext
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_434_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_435_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_434_ext, requiresExtension_435_ext
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_436_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_437_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_436_ext, requiresExtension_437_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_438_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_438_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_439_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_439_ext
	return nil
}
func (ϟa *GlGetSyncivAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_440_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_440_ext
	return nil
}
func (ϟa *GlGetTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_441_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_441_ext
	return nil
}
func (ϟa *GlGetTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_442_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_442_ext
	return nil
}
func (ϟa *GlGetTextureHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_443_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_443_ext
	return nil
}
func (ϟa *GlGetTextureSamplerHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_444_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_444_ext
	return nil
}
func (ϟa *GlGetTranslatedShaderSourceANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_445_ext := ExtensionId_GL_ANGLE_translated_shader_source // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_445_ext
	return nil
}
func (ϟa *GlGetnUniformfvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_446_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_446_ext
	return nil
}
func (ϟa *GlGetnUniformfvKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_447_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_447_ext
	return nil
}
func (ϟa *GlGetnUniformivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_448_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_448_ext
	return nil
}
func (ϟa *GlGetnUniformivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_449_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_449_ext
	return nil
}
func (ϟa *GlGetnUniformuivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_450_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_450_ext
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_451_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_451_ext
	return nil
}
func (ϟa *GlInterpolatePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_452_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_452_ext
	return nil
}
func (ϟa *GlIsEnablediNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_453_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_453_ext
	return nil
}
func (ϟa *GlIsEnablediOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_454_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_454_ext
	return nil
}
func (ϟa *GlIsFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_455_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_455_ext
	return nil
}
func (ϟa *GlIsImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_456_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_456_ext
	return nil
}
func (ϟa *GlIsPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_457_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_457_ext
	return nil
}
func (ϟa *GlIsPointInFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_458_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_458_ext
	return nil
}
func (ϟa *GlIsPointInStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_459_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_459_ext
	return nil
}
func (ϟa *GlIsProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_460_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_460_ext
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_461_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_462_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_464_msg := "No context bound" // string
		return
		_ = error_464_msg
	}
	GetContext_463_result := context // Contextʳ
	ctx := GetContext_463_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = requiresExtension_461_ext, requiresExtension_462_ext, context, GetContext_463_result, ctx
	return nil
}
func (ϟa *GlIsSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_465_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_465_ext
	return nil
}
func (ϟa *GlIsTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_466_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_466_ext
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_467_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_469_msg := "No context bound" // string
		return
		_ = error_469_msg
	}
	GetContext_468_result := context // Contextʳ
	ctx := GetContext_468_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.VertexArrays.Contains(ϟa.Array) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _ = requiresExtension_467_ext, context, GetContext_468_result, ctx
	return nil
}
func (ϟa *GlLabelObjectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_470_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_470_ext
	return nil
}
func (ϟa *GlMakeImageHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_471_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_471_ext
	return nil
}
func (ϟa *GlMakeImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_472_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_472_ext
	return nil
}
func (ϟa *GlMakeTextureHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_473_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_473_ext
	return nil
}
func (ϟa *GlMakeTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_474_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_474_ext
	return nil
}
func (ϟa *GlMapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_475_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_475_ext
	return nil
}
func (ϟa *GlMapBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_476_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_476_ext
	return nil
}
func (ϟa *GlMatrixLoad3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_477_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_477_ext
	return nil
}
func (ϟa *GlMatrixLoad3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_478_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_478_ext
	return nil
}
func (ϟa *GlMatrixLoadTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_479_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_479_ext
	return nil
}
func (ϟa *GlMatrixMult3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_480_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_480_ext
	return nil
}
func (ϟa *GlMatrixMult3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_481_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_481_ext
	return nil
}
func (ϟa *GlMatrixMultTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_482_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_482_ext
	return nil
}
func (ϟa *GlMultiDrawArraysEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_483_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_483_ext
	return nil
}
func (ϟa *GlMultiDrawArraysIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_484_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_484_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_485_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_485_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_486_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_486_ext
	return nil
}
func (ϟa *GlMultiDrawElementsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_487_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_487_ext
	return nil
}
func (ϟa *GlMultiDrawElementsIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_488_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_488_ext
	return nil
}
func (ϟa *GlNamedFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_489_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_489_ext
	return nil
}
func (ϟa *GlPatchParameteriOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_490_ext := ExtensionId_GL_OES_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_490_ext
	return nil
}
func (ϟa *GlPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_491_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_491_ext
	return nil
}
func (ϟa *GlPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_492_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_492_ext
	return nil
}
func (ϟa *GlPathCoverDepthFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_493_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_493_ext
	return nil
}
func (ϟa *GlPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_494_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_494_ext
	return nil
}
func (ϟa *GlPathGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_495_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_495_ext
	return nil
}
func (ϟa *GlPathGlyphIndexRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_496_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_496_ext
	return nil
}
func (ϟa *GlPathGlyphRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_497_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_497_ext
	return nil
}
func (ϟa *GlPathGlyphsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_498_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_498_ext
	return nil
}
func (ϟa *GlPathMemoryGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_499_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_499_ext
	return nil
}
func (ϟa *GlPathParameterfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_500_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_500_ext
	return nil
}
func (ϟa *GlPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_501_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_501_ext
	return nil
}
func (ϟa *GlPathParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_502_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_502_ext
	return nil
}
func (ϟa *GlPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_503_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_503_ext
	return nil
}
func (ϟa *GlPathStencilDepthOffsetNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_504_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_504_ext
	return nil
}
func (ϟa *GlPathStencilFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_505_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_505_ext
	return nil
}
func (ϟa *GlPathStringNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_506_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_506_ext
	return nil
}
func (ϟa *GlPathSubCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_507_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_507_ext
	return nil
}
func (ϟa *GlPathSubCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_508_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_508_ext
	return nil
}
func (ϟa *GlPointAlongPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_509_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_509_ext
	return nil
}
func (ϟa *GlPolygonModeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_510_ext := ExtensionId_GL_NV_polygon_mode // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_510_ext
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_511_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_511_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_512_ext := ExtensionId_GL_OES_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_512_ext
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_513_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_513_ext
	return nil
}
func (ϟa *GlProgramParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_514_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_514_ext
	return nil
}
func (ϟa *GlProgramPathFragmentInputGenNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_515_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_515_ext
	return nil
}
func (ϟa *GlProgramUniform1fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_516_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_516_ext
	return nil
}
func (ϟa *GlProgramUniform1fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_517_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_517_ext
	return nil
}
func (ϟa *GlProgramUniform1iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_518_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_518_ext
	return nil
}
func (ϟa *GlProgramUniform1ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_519_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_519_ext
	return nil
}
func (ϟa *GlProgramUniform1uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_520_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_520_ext
	return nil
}
func (ϟa *GlProgramUniform1uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_521_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_521_ext
	return nil
}
func (ϟa *GlProgramUniform2fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_522_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_522_ext
	return nil
}
func (ϟa *GlProgramUniform2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_523_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_523_ext
	return nil
}
func (ϟa *GlProgramUniform2iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_524_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_524_ext
	return nil
}
func (ϟa *GlProgramUniform2ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_525_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_525_ext
	return nil
}
func (ϟa *GlProgramUniform2uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_526_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_526_ext
	return nil
}
func (ϟa *GlProgramUniform2uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_527_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_527_ext
	return nil
}
func (ϟa *GlProgramUniform3fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_528_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_528_ext
	return nil
}
func (ϟa *GlProgramUniform3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_529_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_529_ext
	return nil
}
func (ϟa *GlProgramUniform3iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_530_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_530_ext
	return nil
}
func (ϟa *GlProgramUniform3ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_531_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_531_ext
	return nil
}
func (ϟa *GlProgramUniform3uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_532_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_532_ext
	return nil
}
func (ϟa *GlProgramUniform3uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_533_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_533_ext
	return nil
}
func (ϟa *GlProgramUniform4fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_534_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_534_ext
	return nil
}
func (ϟa *GlProgramUniform4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_535_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_535_ext
	return nil
}
func (ϟa *GlProgramUniform4iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_536_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_536_ext
	return nil
}
func (ϟa *GlProgramUniform4ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_537_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_537_ext
	return nil
}
func (ϟa *GlProgramUniform4uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_538_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_538_ext
	return nil
}
func (ϟa *GlProgramUniform4uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_539_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_539_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_540_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_540_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_541_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_541_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_542_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_542_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_543_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_543_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_544_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_544_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_545_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_545_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_546_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_546_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_547_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_547_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_548_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_548_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_549_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_549_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_550_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_550_ext
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_551_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_551_ext
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_552_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_552_ext
	return nil
}
func (ϟa *GlRasterSamplesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_553_ext := ExtensionId_GL_EXT_raster_multisample       // ExtensionId
	requiresExtension_554_ext := ExtensionId_GL_EXT_texture_filter_minmax    // ExtensionId
	requiresExtension_555_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_553_ext, requiresExtension_554_ext, requiresExtension_555_ext
	return nil
}
func (ϟa *GlReadBufferIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_556_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_556_ext
	return nil
}
func (ϟa *GlReadBufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_557_ext := ExtensionId_GL_NV_read_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_557_ext
	return nil
}
func (ϟa *GlReadnPixelsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_558_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_558_ext
	return nil
}
func (ϟa *GlReadnPixelsKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_559_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_559_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_560_ext := ExtensionId_GL_ANGLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_560_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_561_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_561_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_562_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_562_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_563_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_563_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_564_ext := ExtensionId_GL_NV_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_564_ext
	return nil
}
func (ϟa *GlResolveDepthValuesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_565_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_565_ext
	return nil
}
func (ϟa *GlResolveMultisampleFramebufferAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_566_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_566_ext
	return nil
}
func (ϟa *GlSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_567_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_567_ext
	return nil
}
func (ϟa *GlSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_568_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_568_ext
	return nil
}
func (ϟa *GlScissorArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_569_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_569_ext
	return nil
}
func (ϟa *GlScissorIndexedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_570_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_570_ext
	return nil
}
func (ϟa *GlScissorIndexedvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_571_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_571_ext
	return nil
}
func (ϟa *GlSelectPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_572_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_572_ext
	return nil
}
func (ϟa *GlSetFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_573_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_573_ext
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_574_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_574_ext
	return nil
}
func (ϟa *GlStencilFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_575_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_575_ext
	return nil
}
func (ϟa *GlStencilFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_576_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_576_ext
	return nil
}
func (ϟa *GlStencilStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_577_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_577_ext
	return nil
}
func (ϟa *GlStencilStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_578_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_578_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_579_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_579_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_580_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_580_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_581_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_581_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_582_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_582_ext
	return nil
}
func (ϟa *GlSubpixelPrecisionBiasNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_583_ext := ExtensionId_GL_NV_conservative_raster // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_583_ext
	return nil
}
func (ϟa *GlTestFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_584_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_584_ext
	return nil
}
func (ϟa *GlTexBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_585_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_585_ext
	return nil
}
func (ϟa *GlTexBufferRangeOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_586_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_586_ext
	return nil
}
func (ϟa *GlTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_587_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_587_ext
	return nil
}
func (ϟa *GlTexPageCommitmentARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_588_ext := ExtensionId_GL_EXT_sparse_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_588_ext
	return nil
}
func (ϟa *GlTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_589_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_589_ext
	return nil
}
func (ϟa *GlTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_590_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_590_ext
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_591_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_591_ext
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_592_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_592_ext
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_593_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_593_ext
	return nil
}
func (ϟa *GlTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_594_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_594_ext
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_595_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_595_ext
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_596_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_596_ext
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_597_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_597_ext
	return nil
}
func (ϟa *GlTextureViewEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_598_ext := ExtensionId_GL_EXT_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_598_ext
	return nil
}
func (ϟa *GlTextureViewOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_599_ext := ExtensionId_GL_OES_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_599_ext
	return nil
}
func (ϟa *GlTransformPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_600_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_600_ext
	return nil
}
func (ϟa *GlUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_601_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_601_ext
	return nil
}
func (ϟa *GlUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_602_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_602_ext
	return nil
}
func (ϟa *GlUniformMatrix2x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_603_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_603_ext
	return nil
}
func (ϟa *GlUniformMatrix2x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_604_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_604_ext
	return nil
}
func (ϟa *GlUniformMatrix3x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_605_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_605_ext
	return nil
}
func (ϟa *GlUniformMatrix3x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_606_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_606_ext
	return nil
}
func (ϟa *GlUniformMatrix4x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_607_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_607_ext
	return nil
}
func (ϟa *GlUniformMatrix4x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_608_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_608_ext
	return nil
}
func (ϟa *GlUnmapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_609_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_609_ext
	return nil
}
func (ϟa *GlUseProgramStagesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_610_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_610_ext
	return nil
}
func (ϟa *GlValidateProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_611_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_611_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_612_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_612_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_613_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_613_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_614_ext := ExtensionId_GL_NV_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_614_ext
	return nil
}
func (ϟa *GlViewportArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_615_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_615_ext
	return nil
}
func (ϟa *GlViewportIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_616_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_616_ext
	return nil
}
func (ϟa *GlViewportIndexedfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_617_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_617_ext
	return nil
}
func (ϟa *GlWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_618_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_618_ext
	return nil
}
func (ϟa *GlWeightPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_619_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_619_ext
	return nil
}
func (ϟa *GlBlendBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_620_major := uint32(3) // u32
	minRequiredVersion_620_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_620_major, minRequiredVersion_620_minor
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	ctx.Blending.BlendColor = Color{Red: ϟa.Red, Green: ϟa.Green, Blue: ϟa.Blue, Alpha: ϟa.Alpha}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_621_major, minRequiredVersion_621_minor, context, GetContext_622_result, ctx
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_624_major := uint32(2) // u32
	minRequiredVersion_624_minor := uint32(0) // u32
	switch ϟa.Equation {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_625_major := uint32(3) // u32
		minRequiredVersion_625_minor := uint32(0) // u32
		_, _ = minRequiredVersion_625_major, minRequiredVersion_625_minor
	default:
		glErrorInvalidEnum_626_param := ϟa.Equation // GLenum
		return
		_ = glErrorInvalidEnum_626_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_628_msg := "No context bound" // string
		return
		_ = error_628_msg
	}
	GetContext_627_result := context // Contextʳ
	ctx := GetContext_627_result     // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_624_major, minRequiredVersion_624_minor, context, GetContext_627_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_629_major := uint32(2) // u32
	minRequiredVersion_629_minor := uint32(0) // u32
	switch ϟa.Rgb {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_630_major := uint32(3) // u32
		minRequiredVersion_630_minor := uint32(0) // u32
		_, _ = minRequiredVersion_630_major, minRequiredVersion_630_minor
	default:
		glErrorInvalidEnum_631_param := ϟa.Rgb // GLenum
		return
		_ = glErrorInvalidEnum_631_param
	}
	switch ϟa.Alpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_632_major := uint32(3) // u32
		minRequiredVersion_632_minor := uint32(0) // u32
		_, _ = minRequiredVersion_632_major, minRequiredVersion_632_minor
	default:
		glErrorInvalidEnum_633_param := ϟa.Alpha // GLenum
		return
		_ = glErrorInvalidEnum_633_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_635_msg := "No context bound" // string
		return
		_ = error_635_msg
	}
	GetContext_634_result := context // Contextʳ
	ctx := GetContext_634_result     // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_629_major, minRequiredVersion_629_minor, context, GetContext_634_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparatei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_636_major := uint32(3) // u32
	minRequiredVersion_636_minor := uint32(2) // u32
	switch ϟa.ModeRGB {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		glErrorInvalidEnum_637_param := ϟa.ModeRGB // GLenum
		return
		_ = glErrorInvalidEnum_637_param
	}
	switch ϟa.ModeAlpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		glErrorInvalidEnum_638_param := ϟa.ModeAlpha // GLenum
		return
		_ = glErrorInvalidEnum_638_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_636_major, minRequiredVersion_636_minor
	return nil
}
func (ϟa *GlBlendEquationi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_639_major := uint32(3) // u32
	minRequiredVersion_639_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		glErrorInvalidEnum_640_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_640_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_639_major, minRequiredVersion_639_minor
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_641_major := uint32(2) // u32
	minRequiredVersion_641_minor := uint32(0) // u32
	switch ϟa.SrcFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_642_param := ϟa.SrcFactor // GLenum
		return
		_ = glErrorInvalidEnum_642_param
	}
	switch ϟa.DstFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_643_param := ϟa.DstFactor // GLenum
		return
		_ = glErrorInvalidEnum_643_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_645_msg := "No context bound" // string
		return
		_ = error_645_msg
	}
	GetContext_644_result := context // Contextʳ
	ctx := GetContext_644_result     // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_641_major, minRequiredVersion_641_minor, context, GetContext_644_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_646_major := uint32(2) // u32
	minRequiredVersion_646_minor := uint32(0) // u32
	switch ϟa.SrcFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_647_param := ϟa.SrcFactorRgb // GLenum
		return
		_ = glErrorInvalidEnum_647_param
	}
	switch ϟa.DstFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_648_param := ϟa.DstFactorRgb // GLenum
		return
		_ = glErrorInvalidEnum_648_param
	}
	switch ϟa.SrcFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_649_param := ϟa.SrcFactorAlpha // GLenum
		return
		_ = glErrorInvalidEnum_649_param
	}
	switch ϟa.DstFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_650_param := ϟa.DstFactorAlpha // GLenum
		return
		_ = glErrorInvalidEnum_650_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_652_msg := "No context bound" // string
		return
		_ = error_652_msg
	}
	GetContext_651_result := context // Contextʳ
	ctx := GetContext_651_result     // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_646_major, minRequiredVersion_646_minor, context, GetContext_651_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparatei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_653_major := uint32(3) // u32
	minRequiredVersion_653_minor := uint32(2) // u32
	switch ϟa.SrcRGB {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_654_param := ϟa.SrcRGB // GLenum
		return
		_ = glErrorInvalidEnum_654_param
	}
	switch ϟa.DstRGB {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_655_param := ϟa.DstRGB // GLenum
		return
		_ = glErrorInvalidEnum_655_param
	}
	switch ϟa.SrcAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_656_param := ϟa.SrcAlpha // GLenum
		return
		_ = glErrorInvalidEnum_656_param
	}
	switch ϟa.DstAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_657_param := ϟa.DstAlpha // GLenum
		return
		_ = glErrorInvalidEnum_657_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_653_major, minRequiredVersion_653_minor
	return nil
}
func (ϟa *GlBlendFunci) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_658_major := uint32(3) // u32
	minRequiredVersion_658_minor := uint32(2) // u32
	switch ϟa.Src {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_659_param := ϟa.Src // GLenum
		return
		_ = glErrorInvalidEnum_659_param
	}
	switch ϟa.Dst {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_660_param := ϟa.Dst // GLenum
		return
		_ = glErrorInvalidEnum_660_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_658_major, minRequiredVersion_658_minor
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_661_major := uint32(2) // u32
	minRequiredVersion_661_minor := uint32(0) // u32
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		glErrorInvalidEnum_662_param := ϟa.Function // GLenum
		return
		_ = glErrorInvalidEnum_662_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_664_msg := "No context bound" // string
		return
		_ = error_664_msg
	}
	GetContext_663_result := context // Contextʳ
	ctx := GetContext_663_result     // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_661_major, minRequiredVersion_661_minor, context, GetContext_663_result, ctx
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_665_major := uint32(2)    // u32
	minRequiredVersion_665_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_667_msg := "No context bound" // string
		return
		_ = error_667_msg
	}
	GetContext_666_result := context // Contextʳ
	ctx := GetContext_666_result     // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_665_major, minRequiredVersion_665_minor, context, GetContext_666_result, ctx
	return nil
}
func (ϟa *GlSampleMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_668_major := uint32(3) // u32
	minRequiredVersion_668_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_668_major, minRequiredVersion_668_minor
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_669_major := uint32(2)    // u32
	minRequiredVersion_669_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_671_msg := "No context bound" // string
		return
		_ = error_671_msg
	}
	GetContext_670_result := context // Contextʳ
	ctx := GetContext_670_result     // Contextʳ
	ctx.Rasterizing.Scissor = Rect{X: ϟa.X, Y: ϟa.Y, Width: ϟa.Width, Height: ϟa.Height}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_669_major, minRequiredVersion_669_minor, context, GetContext_670_result, ctx
	return nil
}
func (ϟa *GlStencilFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_672_major := uint32(2) // u32
	minRequiredVersion_672_minor := uint32(0) // u32
	switch ϟa.Func {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		glErrorInvalidEnum_673_param := ϟa.Func // GLenum
		return
		_ = glErrorInvalidEnum_673_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_672_major, minRequiredVersion_672_minor
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_674_major := uint32(2) // u32
	minRequiredVersion_674_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_675_param := ϟa.Face // GLenum
		return
		_ = glErrorInvalidEnum_675_param
	}
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		glErrorInvalidEnum_676_param := ϟa.Function // GLenum
		return
		_ = glErrorInvalidEnum_676_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_674_major, minRequiredVersion_674_minor
	return nil
}
func (ϟa *GlStencilOp) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_677_major := uint32(2) // u32
	minRequiredVersion_677_minor := uint32(0) // u32
	switch ϟa.Fail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_678_param := ϟa.Fail // GLenum
		return
		_ = glErrorInvalidEnum_678_param
	}
	switch ϟa.Zfail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_679_param := ϟa.Zfail // GLenum
		return
		_ = glErrorInvalidEnum_679_param
	}
	switch ϟa.Zpass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_680_param := ϟa.Zpass // GLenum
		return
		_ = glErrorInvalidEnum_680_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_677_major, minRequiredVersion_677_minor
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_681_major := uint32(2) // u32
	minRequiredVersion_681_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_682_param := ϟa.Face // GLenum
		return
		_ = glErrorInvalidEnum_682_param
	}
	switch ϟa.StencilFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_683_param := ϟa.StencilFail // GLenum
		return
		_ = glErrorInvalidEnum_683_param
	}
	switch ϟa.StencilPassDepthFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_684_param := ϟa.StencilPassDepthFail // GLenum
		return
		_ = glErrorInvalidEnum_684_param
	}
	switch ϟa.StencilPassDepthPass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		glErrorInvalidEnum_685_param := ϟa.StencilPassDepthPass // GLenum
		return
		_ = glErrorInvalidEnum_685_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_681_major, minRequiredVersion_681_minor
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_686_major := uint32(2) // u32
	minRequiredVersion_686_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_687_major := uint32(3) // u32
		minRequiredVersion_687_minor := uint32(0) // u32
		_, _ = minRequiredVersion_687_major, minRequiredVersion_687_minor
	default:
		glErrorInvalidEnum_688_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_688_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_690_msg := "No context bound" // string
		return
		_ = error_690_msg
	}
	GetContext_689_result := context // Contextʳ
	ctx := GetContext_689_result     // Contextʳ
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
	_, _, _, _, _ = minRequiredVersion_686_major, minRequiredVersion_686_minor, context, GetContext_689_result, ctx
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_691_major := uint32(2) // u32
	minRequiredVersion_691_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_692_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_692_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_694_msg := "No context bound" // string
		return
		_ = error_694_msg
	}
	GetContext_693_result := context // Contextʳ
	ctx := GetContext_693_result     // Contextʳ
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = &Renderbuffer{}
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_691_major, minRequiredVersion_691_minor, context, GetContext_693_result, ctx
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_695_major := uint32(3)                                                                                                // u32
	minRequiredVersion_695_minor := uint32(0)                                                                                                // u32
	supportsBits_696_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_696_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	switch ϟa.Filter {
	case GLenum_GL_LINEAR, GLenum_GL_NEAREST:
	default:
		glErrorInvalidEnum_697_param := ϟa.Filter // GLenum
		return
		_ = glErrorInvalidEnum_697_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_695_major, minRequiredVersion_695_minor, supportsBits_696_seenBits, supportsBits_696_validBits
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_698_major := uint32(2) // u32
	minRequiredVersion_698_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_699_major := uint32(3) // u32
		minRequiredVersion_699_minor := uint32(0) // u32
		_, _ = minRequiredVersion_699_major, minRequiredVersion_699_minor
	default:
		glErrorInvalidEnum_700_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_700_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_698_major, minRequiredVersion_698_minor
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_701_major := uint32(2)                                                                                                // u32
	minRequiredVersion_701_minor := uint32(0)                                                                                                // u32
	supportsBits_702_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_702_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_701_major, minRequiredVersion_701_minor, supportsBits_702_seenBits, supportsBits_702_validBits
	return nil
}
func (ϟa *GlClearBufferfi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_703_major := uint32(3) // u32
	minRequiredVersion_703_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_DEPTH_STENCIL:
		glErrorInvalidValueIf_704_condition := (ϟa.Drawbuffer) != (GLint(int32(0))) // bool
		if glErrorInvalidValueIf_704_condition {
			return
		}
		_ = glErrorInvalidValueIf_704_condition
	default:
		glErrorInvalidEnum_705_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_705_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_703_major, minRequiredVersion_703_minor
	return nil
}
func (ϟa *GlClearBufferfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_706_major := uint32(3) // u32
	minRequiredVersion_706_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
		glErrorInvalidValueIf_707_condition := (ϟa.Drawbuffer) >= (GLint(FramebufferConstants_MAX_DRAW_BUFFERS)) // bool
		if glErrorInvalidValueIf_707_condition {
			return
		}
		ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = glErrorInvalidValueIf_707_condition
	case GLenum_GL_DEPTH:
		glErrorInvalidValueIf_708_condition := (ϟa.Drawbuffer) != (GLint(int32(0))) // bool
		if glErrorInvalidValueIf_708_condition {
			return
		}
		ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = glErrorInvalidValueIf_708_condition
	default:
		glErrorInvalidEnum_709_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_709_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_706_major, minRequiredVersion_706_minor
	return nil
}
func (ϟa *GlClearBufferiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_710_major := uint32(3) // u32
	minRequiredVersion_710_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
		glErrorInvalidValueIf_711_condition := (ϟa.Drawbuffer) >= (GLint(FramebufferConstants_MAX_DRAW_BUFFERS)) // bool
		if glErrorInvalidValueIf_711_condition {
			return
		}
		ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = glErrorInvalidValueIf_711_condition
	case GLenum_GL_STENCIL:
		glErrorInvalidValueIf_712_condition := (ϟa.Drawbuffer) != (GLint(int32(0))) // bool
		if glErrorInvalidValueIf_712_condition {
			return
		}
		ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = glErrorInvalidValueIf_712_condition
	default:
		glErrorInvalidEnum_713_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_713_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_710_major, minRequiredVersion_710_minor
	return nil
}
func (ϟa *GlClearBufferuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_714_major := uint32(3) // u32
	minRequiredVersion_714_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
		glErrorInvalidValueIf_715_condition := (ϟa.Drawbuffer) >= (GLint(FramebufferConstants_MAX_DRAW_BUFFERS)) // bool
		if glErrorInvalidValueIf_715_condition {
			return
		}
		ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = glErrorInvalidValueIf_715_condition
	default:
		glErrorInvalidEnum_716_param := ϟa.Buffer // GLenum
		return
		_ = glErrorInvalidEnum_716_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_714_major, minRequiredVersion_714_minor
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_717_major := uint32(2)    // u32
	minRequiredVersion_717_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_719_msg := "No context bound" // string
		return
		_ = error_719_msg
	}
	GetContext_718_result := context // Contextʳ
	ctx := GetContext_718_result     // Contextʳ
	ctx.Clearing.ClearColor = Color{Red: ϟa.R, Green: ϟa.G, Blue: ϟa.B, Alpha: ϟa.A}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_717_major, minRequiredVersion_717_minor, context, GetContext_718_result, ctx
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_720_major := uint32(2)    // u32
	minRequiredVersion_720_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_722_msg := "No context bound" // string
		return
		_ = error_722_msg
	}
	GetContext_721_result := context // Contextʳ
	ctx := GetContext_721_result     // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_720_major, minRequiredVersion_720_minor, context, GetContext_721_result, ctx
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_723_major := uint32(2)    // u32
	minRequiredVersion_723_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_725_msg := "No context bound" // string
		return
		_ = error_725_msg
	}
	GetContext_724_result := context // Contextʳ
	ctx := GetContext_724_result     // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_723_major, minRequiredVersion_723_minor, context, GetContext_724_result, ctx
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_726_major := uint32(2)    // u32
	minRequiredVersion_726_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_728_msg := "No context bound" // string
		return
		_ = error_728_msg
	}
	GetContext_727_result := context // Contextʳ
	ctx := GetContext_727_result     // Contextʳ
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_726_major, minRequiredVersion_726_minor, context, GetContext_727_result, ctx
	return nil
}
func (ϟa *GlColorMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_729_major := uint32(3) // u32
	minRequiredVersion_729_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_729_major, minRequiredVersion_729_minor
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_730_major := uint32(2)                                   // u32
	minRequiredVersion_730_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_732_msg := "No context bound" // string
		return
		_ = error_732_msg
	}
	GetContext_731_result := context // Contextʳ
	ctx := GetContext_731_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Framebuffers, f.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_730_major, minRequiredVersion_730_minor, f, context, GetContext_731_result, ctx
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_733_major := uint32(2)                                    // u32
	minRequiredVersion_733_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_735_msg := "No context bound" // string
		return
		_ = error_735_msg
	}
	GetContext_734_result := context // Contextʳ
	ctx := GetContext_734_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Renderbuffers, r.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_733_major, minRequiredVersion_733_minor, r, context, GetContext_734_result, ctx
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_736_major := uint32(2)    // u32
	minRequiredVersion_736_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_738_msg := "No context bound" // string
		return
		_ = error_738_msg
	}
	GetContext_737_result := context // Contextʳ
	ctx := GetContext_737_result     // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_736_major, minRequiredVersion_736_minor, context, GetContext_737_result, ctx
	return nil
}
func (ϟa *GlFramebufferParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_739_major := uint32(3) // u32
	minRequiredVersion_739_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_740_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_740_param
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	case GLenum_GL_FRAMEBUFFER_DEFAULT_LAYERS:
		minRequiredVersion_741_major := uint32(3) // u32
		minRequiredVersion_741_minor := uint32(2) // u32
		_, _ = minRequiredVersion_741_major, minRequiredVersion_741_minor
	default:
		glErrorInvalidEnum_742_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_742_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_739_major, minRequiredVersion_739_minor
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_743_major := uint32(2) // u32
	minRequiredVersion_743_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_744_major := uint32(3) // u32
		minRequiredVersion_744_minor := uint32(0) // u32
		_, _ = minRequiredVersion_744_major, minRequiredVersion_744_minor
	default:
		glErrorInvalidEnum_745_param := ϟa.FramebufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_745_param
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_746_major := uint32(3) // u32
		minRequiredVersion_746_minor := uint32(0) // u32
		_, _ = minRequiredVersion_746_major, minRequiredVersion_746_minor
	default:
		glErrorInvalidEnum_747_param := ϟa.FramebufferAttachment // GLenum
		return
		_ = glErrorInvalidEnum_747_param
	}
	switch ϟa.RenderbufferTarget {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_748_param := ϟa.RenderbufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_748_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_750_msg := "No context bound" // string
		return
		_ = error_750_msg
	}
	GetContext_749_result := context // Contextʳ
	ctx := GetContext_749_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_743_major, minRequiredVersion_743_minor, context, GetContext_749_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_751_major := uint32(3) // u32
	minRequiredVersion_751_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_752_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_752_param
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	default:
		glErrorInvalidEnum_753_param := ϟa.Attachment // GLenum
		return
		_ = glErrorInvalidEnum_753_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_751_major, minRequiredVersion_751_minor
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_754_major := uint32(2) // u32
	minRequiredVersion_754_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_755_major := uint32(3) // u32
		minRequiredVersion_755_minor := uint32(0) // u32
		_, _ = minRequiredVersion_755_major, minRequiredVersion_755_minor
	default:
		glErrorInvalidEnum_756_param := ϟa.FramebufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_756_param
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_757_major := uint32(3) // u32
		minRequiredVersion_757_minor := uint32(0) // u32
		_, _ = minRequiredVersion_757_major, minRequiredVersion_757_minor
	default:
		glErrorInvalidEnum_758_param := ϟa.FramebufferAttachment // GLenum
		return
		_ = glErrorInvalidEnum_758_param
	}
	switch ϟa.TextureTarget {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_759_major := uint32(3) // u32
		minRequiredVersion_759_minor := uint32(1) // u32
		_, _ = minRequiredVersion_759_major, minRequiredVersion_759_minor
	default:
		glErrorInvalidEnum_760_param := ϟa.TextureTarget // GLenum
		return
		_ = glErrorInvalidEnum_760_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_762_msg := "No context bound" // string
		return
		_ = error_762_msg
	}
	GetContext_761_result := context // Contextʳ
	ctx := GetContext_761_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_754_major, minRequiredVersion_754_minor, context, GetContext_761_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTextureLayer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_763_major := uint32(3) // u32
	minRequiredVersion_763_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_764_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_764_param
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	default:
		glErrorInvalidEnum_765_param := ϟa.Attachment // GLenum
		return
		_ = glErrorInvalidEnum_765_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_763_major, minRequiredVersion_763_minor
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_766_major := uint32(2)                                   // u32
	minRequiredVersion_766_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_768_msg := "No context bound" // string
		return
		_ = error_768_msg
	}
	GetContext_767_result := context // Contextʳ
	ctx := GetContext_767_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // FramebufferId
		ctx.Instances.Framebuffers[id] = &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}}
		f.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_766_major, minRequiredVersion_766_minor, f, context, GetContext_767_result, ctx
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_769_major := uint32(2)                                    // u32
	minRequiredVersion_769_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_771_msg := "No context bound" // string
		return
		_ = error_771_msg
	}
	GetContext_770_result := context // Contextʳ
	ctx := GetContext_770_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = &Renderbuffer{}
		r.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_769_major, minRequiredVersion_769_minor, r, context, GetContext_770_result, ctx
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_772_major := uint32(2) // u32
	minRequiredVersion_772_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_773_major := uint32(3) // u32
		minRequiredVersion_773_minor := uint32(0) // u32
		_, _ = minRequiredVersion_773_major, minRequiredVersion_773_minor
	default:
		glErrorInvalidEnum_774_param := ϟa.FramebufferTarget // GLenum
		return
		_ = glErrorInvalidEnum_774_param
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL:
		minRequiredVersion_775_major := uint32(3) // u32
		minRequiredVersion_775_minor := uint32(0) // u32
		_, _ = minRequiredVersion_775_major, minRequiredVersion_775_minor
	default:
		glErrorInvalidEnum_776_param := ϟa.Attachment // GLenum
		return
		_ = glErrorInvalidEnum_776_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME, GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		minRequiredVersion_777_major := uint32(3) // u32
		minRequiredVersion_777_minor := uint32(0) // u32
		_, _ = minRequiredVersion_777_major, minRequiredVersion_777_minor
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_LAYERED:
		minRequiredVersion_778_major := uint32(3) // u32
		minRequiredVersion_778_minor := uint32(2) // u32
		_, _ = minRequiredVersion_778_major, minRequiredVersion_778_minor
	default:
		glErrorInvalidEnum_779_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_779_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_781_msg := "No context bound" // string
		return
		_ = error_781_msg
	}
	GetContext_780_result := context // Contextʳ
	ctx := GetContext_780_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_772_major, minRequiredVersion_772_minor, context, GetContext_780_result, ctx, target, framebufferId, framebuffer, a
	return nil
}
func (ϟa *GlGetFramebufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_782_major := uint32(3) // u32
	minRequiredVersion_782_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_783_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_783_param
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	case GLenum_GL_FRAMEBUFFER_DEFAULT_LAYERS:
		minRequiredVersion_784_major := uint32(3) // u32
		minRequiredVersion_784_minor := uint32(2) // u32
		_, _ = minRequiredVersion_784_major, minRequiredVersion_784_minor
	default:
		glErrorInvalidEnum_785_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_785_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_782_major, minRequiredVersion_782_minor
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_786_major := uint32(2) // u32
	minRequiredVersion_786_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_787_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_787_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_RENDERBUFFER_ALPHA_SIZE, GLenum_GL_RENDERBUFFER_BLUE_SIZE, GLenum_GL_RENDERBUFFER_DEPTH_SIZE, GLenum_GL_RENDERBUFFER_GREEN_SIZE, GLenum_GL_RENDERBUFFER_HEIGHT, GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT, GLenum_GL_RENDERBUFFER_RED_SIZE, GLenum_GL_RENDERBUFFER_STENCIL_SIZE, GLenum_GL_RENDERBUFFER_WIDTH:
	case GLenum_GL_RENDERBUFFER_SAMPLES:
		minRequiredVersion_788_major := uint32(3) // u32
		minRequiredVersion_788_minor := uint32(0) // u32
		_, _ = minRequiredVersion_788_major, minRequiredVersion_788_minor
	default:
		glErrorInvalidEnum_789_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_789_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_791_msg := "No context bound" // string
		return
		_ = error_791_msg
	}
	GetContext_790_result := context            // Contextʳ
	ctx := GetContext_790_result                // Contextʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_786_major, minRequiredVersion_786_minor, context, GetContext_790_result, ctx, id, rb
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_792_major := uint32(3) // u32
	minRequiredVersion_792_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_793_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_793_param
	}
	a := ϟa.Attachments.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLenumˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = minRequiredVersion_792_major, minRequiredVersion_792_minor, a
	return nil
}
func (ϟa *GlInvalidateSubFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_794_major := uint32(3) // u32
	minRequiredVersion_794_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		glErrorInvalidEnum_795_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_795_param
	}
	a := ϟa.Attachments.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.NumAttachments), ϟs) // GLenumˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = minRequiredVersion_794_major, minRequiredVersion_794_minor, a
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_796_major := uint32(2)    // u32
	minRequiredVersion_796_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_798_msg := "No context bound" // string
		return
		_ = error_798_msg
	}
	GetContext_797_result := context // Contextʳ
	ctx := GetContext_797_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_796_major, minRequiredVersion_796_minor, context, GetContext_797_result, ctx
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_799_major := uint32(2)    // u32
	minRequiredVersion_799_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_801_msg := "No context bound" // string
		return
		_ = error_801_msg
	}
	GetContext_800_result := context // Contextʳ
	ctx := GetContext_800_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_799_major, minRequiredVersion_799_minor, context, GetContext_800_result, ctx
	return nil
}
func (ϟa *GlReadBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_802_major := uint32(3) // u32
	minRequiredVersion_802_minor := uint32(0) // u32
	switch ϟa.Src {
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_NONE:
	default:
		glErrorInvalidEnum_803_param := ϟa.Src // GLenum
		return
		_ = glErrorInvalidEnum_803_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_802_major, minRequiredVersion_802_minor
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_804_major := uint32(2) // u32
	minRequiredVersion_804_minor := uint32(0) // u32
	switch ϟa.Format {
	case GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER:
	default:
		glErrorInvalidEnum_805_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_805_param
	}
	switch ϟa.Type {
	case GLenum_GL_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT:
	case GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_806_major := uint32(3) // u32
		minRequiredVersion_806_minor := uint32(1) // u32
		_, _ = minRequiredVersion_806_major, minRequiredVersion_806_minor
	default:
		glErrorInvalidEnum_807_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_807_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_809_msg := "No context bound" // string
		return
		_ = error_809_msg
	}
	GetContext_808_result := context // Contextʳ
	ctx := GetContext_808_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_PACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (Voidᵖ{})) {
		requiredSize := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type) // u32
		ϟa.Data.Slice(uint64(uint32(0)), uint64(requiredSize), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = requiredSize
	}
	_, _, _, _, _ = minRequiredVersion_804_major, minRequiredVersion_804_minor, context, GetContext_808_result, ctx
	return nil
}
func (ϟa *GlReadnPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_810_major := uint32(3) // u32
	minRequiredVersion_810_minor := uint32(2) // u32
	switch ϟa.Format {
	case GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER:
	default:
		glErrorInvalidEnum_811_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_811_param
	}
	switch ϟa.Type {
	case GLenum_GL_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
	default:
		glErrorInvalidEnum_812_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_812_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_814_msg := "No context bound" // string
		return
		_ = error_814_msg
	}
	GetContext_813_result := context // Contextʳ
	ctx := GetContext_813_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_PACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (Voidᵖ{})) {
		requiredSize := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type) // u32
		glErrorInvalidOperationIf_815_condition := (GLsizei(requiredSize)) > (ϟa.BufSize)                              // bool
		if glErrorInvalidOperationIf_815_condition {
			return
		}
		ϟa.Data.Slice(uint64(uint32(0)), uint64(requiredSize), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = requiredSize, glErrorInvalidOperationIf_815_condition
	}
	_, _, _, _, _ = minRequiredVersion_810_major, minRequiredVersion_810_minor, context, GetContext_813_result, ctx
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_816_major := uint32(2) // u32
	minRequiredVersion_816_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_817_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_817_param
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGBA4, GLenum_GL_STENCIL_INDEX8:
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_818_major := uint32(3) // u32
		minRequiredVersion_818_minor := uint32(0) // u32
		_, _ = minRequiredVersion_818_major, minRequiredVersion_818_minor
	default:
		glErrorInvalidEnum_819_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_819_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_821_msg := "No context bound" // string
		return
		_ = error_821_msg
	}
	GetContext_820_result := context            // Contextʳ
	ctx := GetContext_820_result                // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target) // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)   // Renderbufferʳ
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_816_major, minRequiredVersion_816_minor, context, GetContext_820_result, ctx, id, rb
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_822_major := uint32(3) // u32
	minRequiredVersion_822_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		glErrorInvalidEnum_823_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_823_param
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		glErrorInvalidEnum_824_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_824_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_822_major, minRequiredVersion_822_minor
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_825_major := uint32(2)    // u32
	minRequiredVersion_825_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_827_msg := "No context bound" // string
		return
		_ = error_827_msg
	}
	GetContext_826_result := context // Contextʳ
	ctx := GetContext_826_result     // Contextʳ
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_825_major, minRequiredVersion_825_minor, context, GetContext_826_result, ctx
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_828_major := uint32(2) // u32
	minRequiredVersion_828_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_829_param := ϟa.Face // GLenum
		return
		_ = glErrorInvalidEnum_829_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_831_msg := "No context bound" // string
		return
		_ = error_831_msg
	}
	GetContext_830_result := context // Contextʳ
	ctx := GetContext_830_result     // Contextʳ
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
	_, _, _, _, _ = minRequiredVersion_828_major, minRequiredVersion_828_minor, context, GetContext_830_result, ctx
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_832_major := uint32(2) // u32
	minRequiredVersion_832_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_833_major := uint32(3) // u32
		minRequiredVersion_833_minor := uint32(0) // u32
		_, _ = minRequiredVersion_833_major, minRequiredVersion_833_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_834_major := uint32(3) // u32
		minRequiredVersion_834_minor := uint32(1) // u32
		_, _ = minRequiredVersion_834_major, minRequiredVersion_834_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS:
		minRequiredVersion_835_major := uint32(3) // u32
		minRequiredVersion_835_minor := uint32(2) // u32
		_, _ = minRequiredVersion_835_major, minRequiredVersion_835_minor
	default:
		glErrorInvalidEnum_836_param := ϟa.Capability // GLenum
		return
		_ = glErrorInvalidEnum_836_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_838_msg := "No context bound" // string
		return
		_ = error_838_msg
	}
	GetContext_837_result := context // Contextʳ
	ctx := GetContext_837_result     // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_832_major, minRequiredVersion_832_minor, context, GetContext_837_result, ctx
	return nil
}
func (ϟa *GlDisablei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_839_major := uint32(3) // u32
	minRequiredVersion_839_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		glErrorInvalidEnum_840_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_840_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_839_major, minRequiredVersion_839_minor
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_841_major := uint32(2) // u32
	minRequiredVersion_841_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_842_major := uint32(3) // u32
		minRequiredVersion_842_minor := uint32(0) // u32
		_, _ = minRequiredVersion_842_major, minRequiredVersion_842_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_843_major := uint32(3) // u32
		minRequiredVersion_843_minor := uint32(1) // u32
		_, _ = minRequiredVersion_843_major, minRequiredVersion_843_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS:
		minRequiredVersion_844_major := uint32(3) // u32
		minRequiredVersion_844_minor := uint32(2) // u32
		_, _ = minRequiredVersion_844_major, minRequiredVersion_844_minor
	default:
		glErrorInvalidEnum_845_param := ϟa.Capability // GLenum
		return
		_ = glErrorInvalidEnum_845_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_847_msg := "No context bound" // string
		return
		_ = error_847_msg
	}
	GetContext_846_result := context // Contextʳ
	ctx := GetContext_846_result     // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_841_major, minRequiredVersion_841_minor, context, GetContext_846_result, ctx
	return nil
}
func (ϟa *GlEnablei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_848_major := uint32(3) // u32
	minRequiredVersion_848_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		glErrorInvalidEnum_849_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_849_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_848_major, minRequiredVersion_848_minor
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_850_major := uint32(2) // u32
	minRequiredVersion_850_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_850_major, minRequiredVersion_850_minor
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_851_major := uint32(2) // u32
	minRequiredVersion_851_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_851_major, minRequiredVersion_851_minor
	return nil
}
func (ϟa *GlFlushMappedBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_852_major := uint32(3) // u32
	minRequiredVersion_852_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_853_major := uint32(3) // u32
		minRequiredVersion_853_minor := uint32(2) // u32
		_, _ = minRequiredVersion_853_major, minRequiredVersion_853_minor
	default:
		glErrorInvalidEnum_854_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_854_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_852_major, minRequiredVersion_852_minor
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_855_major := uint32(2) // u32
	minRequiredVersion_855_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_855_major, minRequiredVersion_855_minor
	return nil
}
func (ϟa *GlGetGraphicsResetStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_856_major := uint32(3) // u32
	minRequiredVersion_856_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_856_major, minRequiredVersion_856_minor
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_857_major := uint32(2) // u32
	minRequiredVersion_857_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_GENERATE_MIPMAP_HINT:
	case GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT:
		minRequiredVersion_858_major := uint32(3) // u32
		minRequiredVersion_858_minor := uint32(0) // u32
		_, _ = minRequiredVersion_858_major, minRequiredVersion_858_minor
	default:
		glErrorInvalidEnum_859_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_859_param
	}
	switch ϟa.Mode {
	case GLenum_GL_DONT_CARE, GLenum_GL_FASTEST, GLenum_GL_NICEST:
	default:
		glErrorInvalidEnum_860_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_860_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_862_msg := "No context bound" // string
		return
		_ = error_862_msg
	}
	GetContext_861_result := context // Contextʳ
	ctx := GetContext_861_result     // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_857_major, minRequiredVersion_857_minor, context, GetContext_861_result, ctx
	return nil
}
func (ϟa *GlActiveShaderProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_863_major := uint32(3) // u32
	minRequiredVersion_863_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_863_major, minRequiredVersion_863_minor
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_864_major := uint32(2)    // u32
	minRequiredVersion_864_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_866_msg := "No context bound" // string
		return
		_ = error_866_msg
	}
	GetContext_865_result := context            // Contextʳ
	ctx := GetContext_865_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)   // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_864_major, minRequiredVersion_864_minor, context, GetContext_865_result, ctx, p, s
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_867_major := uint32(2)    // u32
	minRequiredVersion_867_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_869_msg := "No context bound" // string
		return
		_ = error_869_msg
	}
	GetContext_868_result := context            // Contextʳ
	ctx := GetContext_868_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_867_major, minRequiredVersion_867_minor, context, GetContext_868_result, ctx, p
	return nil
}
func (ϟa *GlBindProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_870_major := uint32(3) // u32
	minRequiredVersion_870_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_870_major, minRequiredVersion_870_minor
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_871_major := uint32(2) // u32
	minRequiredVersion_871_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_871_major, minRequiredVersion_871_minor
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_872_major := uint32(2)    // u32
	minRequiredVersion_872_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_874_msg := "No context bound" // string
		return
		_ = error_874_msg
	}
	GetContext_873_result := context // Contextʳ
	ctx := GetContext_873_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = &Program{Shaders: GLenumːShaderIdᵐ{}, AttributeBindings: StringːAttributeLocationᵐ{}, Attributes: S32ːVertexAttributeᵐ{}, Uniforms: UniformLocationːUniformᵐ{}}
	ϟa.Result = id
	_, _, _, _, _, _ = minRequiredVersion_872_major, minRequiredVersion_872_minor, context, GetContext_873_result, ctx, id
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_875_major := uint32(2) // u32
	minRequiredVersion_875_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_COMPUTE_SHADER:
		minRequiredVersion_876_major := uint32(3) // u32
		minRequiredVersion_876_minor := uint32(1) // u32
		_, _ = minRequiredVersion_876_major, minRequiredVersion_876_minor
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_877_major := uint32(3) // u32
		minRequiredVersion_877_minor := uint32(2) // u32
		_, _ = minRequiredVersion_877_major, minRequiredVersion_877_minor
	default:
		glErrorInvalidEnum_878_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_878_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_880_msg := "No context bound" // string
		return
		_ = error_880_msg
	}
	GetContext_879_result := context // Contextʳ
	ctx := GetContext_879_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = &Shader{Compiled: false, Deletable: false}
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _, _, _, _, _, _ = minRequiredVersion_875_major, minRequiredVersion_875_minor, context, GetContext_879_result, ctx, id, s
	return nil
}
func (ϟa *GlCreateShaderProgramv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_881_major := uint32(3) // u32
	minRequiredVersion_881_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_882_major := uint32(3) // u32
		minRequiredVersion_882_minor := uint32(2) // u32
		_, _ = minRequiredVersion_882_major, minRequiredVersion_882_minor
	default:
		glErrorInvalidEnum_883_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_883_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_885_msg := "No context bound" // string
		return
		_ = error_885_msg
	}
	GetContext_884_result := context                                             // Contextʳ
	ctx := GetContext_884_result                                                 // Contextʳ
	sources := ϟa.Strings.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		source := strings.TrimRight(string(Charᵖ(sources.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
		_ = source
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _ = minRequiredVersion_881_major, minRequiredVersion_881_minor, context, GetContext_884_result, ctx, sources
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_886_major := uint32(2)    // u32
	minRequiredVersion_886_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_888_msg := "No context bound" // string
		return
		_ = error_888_msg
	}
	GetContext_887_result := context // Contextʳ
	ctx := GetContext_887_result     // Contextʳ
	delete(ctx.Instances.Programs, ϟa.Program)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_886_major, minRequiredVersion_886_minor, context, GetContext_887_result, ctx
	return nil
}
func (ϟa *GlDeleteProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_889_major := uint32(3)                            // u32
	minRequiredVersion_889_minor := uint32(1)                            // u32
	p := ϟa.Pipelines.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.N), ϟs) // PipelineIdˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = minRequiredVersion_889_major, minRequiredVersion_889_minor, p
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_890_major := uint32(2)    // u32
	minRequiredVersion_890_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_892_msg := "No context bound" // string
		return
		_ = error_892_msg
	}
	GetContext_891_result := context          // Contextʳ
	ctx := GetContext_891_result              // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
	s.Deletable = true
	delete(ctx.Instances.Shaders, ϟa.Shader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_890_major, minRequiredVersion_890_minor, context, GetContext_891_result, ctx, s
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_893_major := uint32(2)    // u32
	minRequiredVersion_893_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_895_msg := "No context bound" // string
		return
		_ = error_895_msg
	}
	GetContext_894_result := context            // Contextʳ
	ctx := GetContext_894_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)   // Shaderʳ
	delete(p.Shaders, s.Type)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_893_major, minRequiredVersion_893_minor, context, GetContext_894_result, ctx, p, s
	return nil
}
func (ϟa *GlDispatchCompute) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_896_major := uint32(3) // u32
	minRequiredVersion_896_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_896_major, minRequiredVersion_896_minor
	return nil
}
func (ϟa *GlDispatchComputeIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_897_major := uint32(3) // u32
	minRequiredVersion_897_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_897_major, minRequiredVersion_897_minor
	return nil
}
func (ϟa *GlGenProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_898_major := uint32(3) // u32
	minRequiredVersion_898_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Pipelines.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.N), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_898_major, minRequiredVersion_898_minor
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_899_major := uint32(2)                     // u32
	minRequiredVersion_899_minor := uint32(0)                     // u32
	writeString_900_buffer_size := ϟa.BufferSize                  // GLsizei
	writeString_900_buffer_bytes_written := ϟa.BufferBytesWritten // GLsizeiᵖ
	writeString_900_buffer := ϟa.Name                             // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_900_buffer) != (GLcharᵖ{})) && ((writeString_900_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_900_buffer_size // GLsizei
		if (writeString_900_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_900_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_900_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_900_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = minRequiredVersion_899_major, minRequiredVersion_899_minor, writeString_900_buffer_size, writeString_900_buffer_bytes_written, writeString_900_buffer
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_901_major := uint32(2)                     // u32
	minRequiredVersion_901_minor := uint32(0)                     // u32
	writeString_902_buffer_size := ϟa.BufferSize                  // GLsizei
	writeString_902_buffer_bytes_written := ϟa.BufferBytesWritten // GLsizeiᵖ
	writeString_902_buffer := ϟa.Name                             // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_902_buffer) != (GLcharᵖ{})) && ((writeString_902_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_902_buffer_size // GLsizei
		if (writeString_902_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_902_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_902_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_902_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = minRequiredVersion_901_major, minRequiredVersion_901_minor, writeString_902_buffer_size, writeString_902_buffer_bytes_written, writeString_902_buffer
	return nil
}
func (ϟa *GlGetActiveUniformBlockName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_903_major := uint32(3)                     // u32
	minRequiredVersion_903_minor := uint32(0)                     // u32
	writeString_904_buffer_size := ϟa.BufferSize                  // GLsizei
	writeString_904_buffer_bytes_written := ϟa.BufferBytesWritten // GLsizeiᵖ
	writeString_904_buffer := ϟa.Name                             // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_904_buffer) != (GLcharᵖ{})) && ((writeString_904_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_904_buffer_size // GLsizei
		if (writeString_904_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_904_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_904_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_904_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _ = minRequiredVersion_903_major, minRequiredVersion_903_minor, writeString_904_buffer_size, writeString_904_buffer_bytes_written, writeString_904_buffer
	return nil
}
func (ϟa *GlGetActiveUniformBlockiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_905_major := uint32(3) // u32
	minRequiredVersion_905_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS, GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES, GLenum_GL_UNIFORM_BLOCK_BINDING, GLenum_GL_UNIFORM_BLOCK_DATA_SIZE, GLenum_GL_UNIFORM_BLOCK_NAME_LENGTH, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
	default:
		glErrorInvalidEnum_906_param := ϟa.ParameterName // GLenum
		return
		_ = glErrorInvalidEnum_906_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_905_major, minRequiredVersion_905_minor
	return nil
}
func (ϟa *GlGetActiveUniformsiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_907_major := uint32(3) // u32
	minRequiredVersion_907_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_ARRAY_STRIDE, GLenum_GL_UNIFORM_BLOCK_INDEX, GLenum_GL_UNIFORM_IS_ROW_MAJOR, GLenum_GL_UNIFORM_MATRIX_STRIDE, GLenum_GL_UNIFORM_NAME_LENGTH, GLenum_GL_UNIFORM_OFFSET, GLenum_GL_UNIFORM_SIZE, GLenum_GL_UNIFORM_TYPE:
	default:
		glErrorInvalidEnum_908_param := ϟa.ParameterName // GLenum
		return
		_ = glErrorInvalidEnum_908_param
	}
	ϟa.UniformIndices.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_907_major, minRequiredVersion_907_minor
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_909_major := uint32(2)    // u32
	minRequiredVersion_909_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_911_msg := "No context bound" // string
		return
		_ = error_911_msg
	}
	GetContext_910_result := context            // Contextʳ
	ctx := GetContext_910_result                // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program) // Programʳ
	min_912_a := int32(ϟa.BufferLength)         // s32
	min_912_b := int32(len(p.Shaders))          // s32
	min_912_result := func() (result int32) {
		switch (min_912_a) < (min_912_b) {
		case true:
			return min_912_a
		case false:
			return min_912_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_912_a) < (min_912_b), ϟa))
			return result
		}
	}() // s32
	l := min_912_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Shaders.Slice(uint64(int32(0)), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	if (ϟa.ShadersLengthWritten) != (GLsizeiᵖ{}) {
		ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_909_major, minRequiredVersion_909_minor, context, GetContext_910_result, ctx, p, min_912_a, min_912_b, min_912_result, l
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_913_major := uint32(2) // u32
	minRequiredVersion_913_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_913_major, minRequiredVersion_913_minor
	return nil
}
func (ϟa *GlGetFragDataLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_914_major := uint32(3) // u32
	minRequiredVersion_914_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_914_major, minRequiredVersion_914_minor
	return nil
}
func (ϟa *GlGetProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_915_major := uint32(3) // u32
	minRequiredVersion_915_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Length) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
		ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.BufSize), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_915_major, minRequiredVersion_915_minor
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_916_major := uint32(2)    // u32
	minRequiredVersion_916_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_918_msg := "No context bound" // string
		return
		_ = error_918_msg
	}
	GetContext_917_result := context                               // Contextʳ
	ctx := GetContext_917_result                                   // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)                    // Programʳ
	writeString_919_buffer_size := ϟa.BufferLength                 // GLsizei
	writeString_919_buffer_bytes_written := ϟa.StringLengthWritten // GLsizeiᵖ
	writeString_919_buffer := ϟa.Info                              // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_919_buffer) != (GLcharᵖ{})) && ((writeString_919_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_919_buffer_size // GLsizei
		if (writeString_919_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_919_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_919_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_919_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_916_major, minRequiredVersion_916_minor, context, GetContext_917_result, ctx, p, writeString_919_buffer_size, writeString_919_buffer_bytes_written, writeString_919_buffer
	return nil
}
func (ϟa *GlGetProgramInterfaceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_920_major := uint32(3) // u32
	minRequiredVersion_920_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_921_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_921_param
	}
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_RESOURCES, GLenum_GL_MAX_NAME_LENGTH, GLenum_GL_MAX_NUM_ACTIVE_VARIABLES:
	default:
		glErrorInvalidEnum_922_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_922_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_920_major, minRequiredVersion_920_minor
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_923_major := uint32(3)         // u32
	minRequiredVersion_923_minor := uint32(1)         // u32
	writeString_924_buffer_size := ϟa.BufSize         // GLsizei
	writeString_924_buffer_bytes_written := ϟa.Length // GLsizeiᵖ
	writeString_924_buffer := ϟa.InfoLog              // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_924_buffer) != (GLcharᵖ{})) && ((writeString_924_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_924_buffer_size // GLsizei
		if (writeString_924_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_924_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_924_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_924_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _ = minRequiredVersion_923_major, minRequiredVersion_923_minor, writeString_924_buffer_size, writeString_924_buffer_bytes_written, writeString_924_buffer
	return nil
}
func (ϟa *GlGetProgramPipelineiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_925_major := uint32(3) // u32
	minRequiredVersion_925_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_PROGRAM, GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_VALIDATE_STATUS, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_926_major := uint32(3) // u32
		minRequiredVersion_926_minor := uint32(2) // u32
		_, _ = minRequiredVersion_926_major, minRequiredVersion_926_minor
	default:
		glErrorInvalidEnum_927_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_927_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_925_major, minRequiredVersion_925_minor
	return nil
}
func (ϟa *GlGetProgramResourceIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_928_major := uint32(3) // u32
	minRequiredVersion_928_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_929_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_929_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_928_major, minRequiredVersion_928_minor
	return nil
}
func (ϟa *GlGetProgramResourceLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_930_major := uint32(3) // u32
	minRequiredVersion_930_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM:
	default:
		glErrorInvalidEnum_931_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_931_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_930_major, minRequiredVersion_930_minor
	return nil
}
func (ϟa *GlGetProgramResourceName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_932_major := uint32(3) // u32
	minRequiredVersion_932_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_933_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_933_param
	}
	writeString_934_buffer_size := ϟa.BufSize         // GLsizei
	writeString_934_buffer_bytes_written := ϟa.Length // GLsizeiᵖ
	writeString_934_buffer := ϟa.Name                 // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_934_buffer) != (GLcharᵖ{})) && ((writeString_934_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_934_buffer_size // GLsizei
		if (writeString_934_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_934_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_934_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_934_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _ = minRequiredVersion_932_major, minRequiredVersion_932_minor, writeString_934_buffer_size, writeString_934_buffer_bytes_written, writeString_934_buffer
	return nil
}
func (ϟa *GlGetProgramResourceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_935_major := uint32(3) // u32
	minRequiredVersion_935_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		glErrorInvalidEnum_936_param := ϟa.ProgramInterface // GLenum
		return
		_ = glErrorInvalidEnum_936_param
	}
	p := ϟa.Props.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.PropCount), ϟs) // GLenumˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Length) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
		ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Params.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Params.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.BufSize), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	_, _, _ = minRequiredVersion_935_major, minRequiredVersion_935_minor, p
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_937_major := uint32(2) // u32
	minRequiredVersion_937_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_ACTIVE_ATTRIBUTES, GLenum_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH, GLenum_GL_ACTIVE_UNIFORMS, GLenum_GL_ACTIVE_UNIFORM_MAX_LENGTH, GLenum_GL_ATTACHED_SHADERS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_LINK_STATUS, GLenum_GL_VALIDATE_STATUS:
	case GLenum_GL_ACTIVE_UNIFORM_BLOCKS, GLenum_GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH, GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_MODE, GLenum_GL_TRANSFORM_FEEDBACK_VARYINGS, GLenum_GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		minRequiredVersion_938_major := uint32(3) // u32
		minRequiredVersion_938_minor := uint32(0) // u32
		_, _ = minRequiredVersion_938_major, minRequiredVersion_938_minor
	case GLenum_GL_ACTIVE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_939_major := uint32(3) // u32
		minRequiredVersion_939_minor := uint32(1) // u32
		_, _ = minRequiredVersion_939_major, minRequiredVersion_939_minor
	case GLenum_GL_GEOMETRY_INPUT_TYPE, GLenum_GL_GEOMETRY_OUTPUT_TYPE, GLenum_GL_GEOMETRY_VERTICES_OUT, GLenum_GL_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_TESS_CONTROL_OUTPUT_VERTICES, GLenum_GL_TESS_GEN_MODE, GLenum_GL_TESS_GEN_POINT_MODE, GLenum_GL_TESS_GEN_SPACING, GLenum_GL_TESS_GEN_VERTEX_ORDER:
		minRequiredVersion_940_major := uint32(3) // u32
		minRequiredVersion_940_minor := uint32(2) // u32
		_, _ = minRequiredVersion_940_major, minRequiredVersion_940_minor
	default:
		glErrorInvalidEnum_941_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_941_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_937_major, minRequiredVersion_937_minor
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
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
	GetContext_943_result := context                               // Contextʳ
	ctx := GetContext_943_result                                   // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)                      // Shaderʳ
	writeString_945_buffer_size := ϟa.BufferLength                 // GLsizei
	writeString_945_buffer_bytes_written := ϟa.StringLengthWritten // GLsizeiᵖ
	writeString_945_buffer := ϟa.Info                              // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_945_buffer) != (GLcharᵖ{})) && ((writeString_945_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_945_buffer_size // GLsizei
		if (writeString_945_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_945_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_945_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_945_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_942_major, minRequiredVersion_942_minor, context, GetContext_943_result, ctx, s, writeString_945_buffer_size, writeString_945_buffer_bytes_written, writeString_945_buffer
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_946_major := uint32(2) // u32
	minRequiredVersion_946_minor := uint32(0) // u32
	switch ϟa.ShaderType {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	default:
		glErrorInvalidEnum_947_param := ϟa.ShaderType // GLenum
		return
		_ = glErrorInvalidEnum_947_param
	}
	switch ϟa.PrecisionType {
	case GLenum_GL_HIGH_FLOAT, GLenum_GL_HIGH_INT, GLenum_GL_LOW_FLOAT, GLenum_GL_LOW_INT, GLenum_GL_MEDIUM_FLOAT, GLenum_GL_MEDIUM_INT:
	default:
		glErrorInvalidEnum_948_param := ϟa.PrecisionType // GLenum
		return
		_ = glErrorInvalidEnum_948_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Range.Slice(uint64(0), uint64(2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_946_major, minRequiredVersion_946_minor
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_949_major := uint32(2)    // u32
	minRequiredVersion_949_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_951_msg := "No context bound" // string
		return
		_ = error_951_msg
	}
	GetContext_950_result := context                               // Contextʳ
	ctx := GetContext_950_result                                   // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)                      // Shaderʳ
	writeString_952_buffer_size := ϟa.BufferLength                 // GLsizei
	writeString_952_buffer_bytes_written := ϟa.StringLengthWritten // GLsizeiᵖ
	writeString_952_buffer := ϟa.Source                            // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_952_buffer) != (GLcharᵖ{})) && ((writeString_952_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_952_buffer_size // GLsizei
		if (writeString_952_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_952_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_952_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_952_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_949_major, minRequiredVersion_949_minor, context, GetContext_950_result, ctx, s, writeString_952_buffer_size, writeString_952_buffer_bytes_written, writeString_952_buffer
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_953_major := uint32(2) // u32
	minRequiredVersion_953_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_COMPILE_STATUS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_SHADER_SOURCE_LENGTH, GLenum_GL_SHADER_TYPE:
	default:
		glErrorInvalidEnum_954_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_954_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_956_msg := "No context bound" // string
		return
		_ = error_956_msg
	}
	GetContext_955_result := context          // Contextʳ
	ctx := GetContext_955_result              // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_953_major, minRequiredVersion_953_minor, context, GetContext_955_result, ctx, s
	return nil
}
func (ϟa *GlGetUniformBlockIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_957_major := uint32(3) // u32
	minRequiredVersion_957_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_957_major, minRequiredVersion_957_minor
	return nil
}
func (ϟa *GlGetUniformIndices) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_958_major := uint32(3)                                              // u32
	minRequiredVersion_958_minor := uint32(0)                                              // u32
	names := ϟa.UniformNames.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs) // GLcharᶜᵖˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.UniformCount; i++ {
		name := strings.TrimRight(string(Charᵖ(names.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
		_ = name
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.UniformIndices.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _ = minRequiredVersion_958_major, minRequiredVersion_958_minor, names
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_959_major := uint32(2) // u32
	minRequiredVersion_959_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_959_major, minRequiredVersion_959_minor
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_960_major := uint32(2) // u32
	minRequiredVersion_960_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_960_major, minRequiredVersion_960_minor
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_961_major := uint32(2) // u32
	minRequiredVersion_961_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_961_major, minRequiredVersion_961_minor
	return nil
}
func (ϟa *GlGetUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_962_major := uint32(3) // u32
	minRequiredVersion_962_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_962_major, minRequiredVersion_962_minor
	return nil
}
func (ϟa *GlGetnUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_963_major := uint32(3) // u32
	minRequiredVersion_963_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_963_major, minRequiredVersion_963_minor
	return nil
}
func (ϟa *GlGetnUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_964_major := uint32(3) // u32
	minRequiredVersion_964_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_964_major, minRequiredVersion_964_minor
	return nil
}
func (ϟa *GlGetnUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_965_major := uint32(3) // u32
	minRequiredVersion_965_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_965_major, minRequiredVersion_965_minor
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_966_major := uint32(2)    // u32
	minRequiredVersion_966_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_968_msg := "No context bound" // string
		return
		_ = error_968_msg
	}
	GetContext_967_result := context // Contextʳ
	ctx := GetContext_967_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Programs.Contains(ϟa.Program) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_966_major, minRequiredVersion_966_minor, context, GetContext_967_result, ctx
	return nil
}
func (ϟa *GlIsProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_969_major := uint32(3) // u32
	minRequiredVersion_969_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_969_major, minRequiredVersion_969_minor
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_970_major := uint32(2)    // u32
	minRequiredVersion_970_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_972_msg := "No context bound" // string
		return
		_ = error_972_msg
	}
	GetContext_971_result := context // Contextʳ
	ctx := GetContext_971_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Shaders.Contains(ϟa.Shader) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_970_major, minRequiredVersion_970_minor, context, GetContext_971_result, ctx
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_973_major := uint32(2) // u32
	minRequiredVersion_973_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_973_major, minRequiredVersion_973_minor
	return nil
}
func (ϟa *GlMemoryBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_974_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_974_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_975_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_975_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_974_major, minRequiredVersion_974_minor, supportsBits_975_seenBits, supportsBits_975_validBits
	return nil
}
func (ϟa *GlMemoryBarrierByRegion) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_976_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_976_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_977_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_977_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_976_major, minRequiredVersion_976_minor, supportsBits_977_seenBits, supportsBits_977_validBits
	return nil
}
func (ϟa *GlProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_978_major := uint32(3) // u32
	minRequiredVersion_978_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		glErrorInvalidEnum_979_param := ϟa.BinaryFormat // GLenum
		return
		_ = glErrorInvalidEnum_979_param
	}
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_978_major, minRequiredVersion_978_minor
	return nil
}
func (ϟa *GlProgramParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_980_major := uint32(3) // u32
	minRequiredVersion_980_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT:
	case GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_981_major := uint32(3) // u32
		minRequiredVersion_981_minor := uint32(1) // u32
		_, _ = minRequiredVersion_981_major, minRequiredVersion_981_minor
	default:
		glErrorInvalidEnum_982_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_982_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_980_major, minRequiredVersion_980_minor
	return nil
}
func (ϟa *GlProgramUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_983_major := uint32(3) // u32
	minRequiredVersion_983_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_983_major, minRequiredVersion_983_minor
	return nil
}
func (ϟa *GlProgramUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_984_major := uint32(3) // u32
	minRequiredVersion_984_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_984_major, minRequiredVersion_984_minor
	return nil
}
func (ϟa *GlProgramUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_985_major := uint32(3) // u32
	minRequiredVersion_985_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_985_major, minRequiredVersion_985_minor
	return nil
}
func (ϟa *GlProgramUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_986_major := uint32(3) // u32
	minRequiredVersion_986_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_986_major, minRequiredVersion_986_minor
	return nil
}
func (ϟa *GlProgramUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_987_major := uint32(3) // u32
	minRequiredVersion_987_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_987_major, minRequiredVersion_987_minor
	return nil
}
func (ϟa *GlProgramUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_988_major := uint32(3) // u32
	minRequiredVersion_988_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_988_major, minRequiredVersion_988_minor
	return nil
}
func (ϟa *GlProgramUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_989_major := uint32(3) // u32
	minRequiredVersion_989_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_989_major, minRequiredVersion_989_minor
	return nil
}
func (ϟa *GlProgramUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_990_major := uint32(3) // u32
	minRequiredVersion_990_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_990_major, minRequiredVersion_990_minor
	return nil
}
func (ϟa *GlProgramUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_991_major := uint32(3) // u32
	minRequiredVersion_991_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_991_major, minRequiredVersion_991_minor
	return nil
}
func (ϟa *GlProgramUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_992_major := uint32(3) // u32
	minRequiredVersion_992_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_992_major, minRequiredVersion_992_minor
	return nil
}
func (ϟa *GlProgramUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_993_major := uint32(3) // u32
	minRequiredVersion_993_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_993_major, minRequiredVersion_993_minor
	return nil
}
func (ϟa *GlProgramUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_994_major := uint32(3) // u32
	minRequiredVersion_994_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_994_major, minRequiredVersion_994_minor
	return nil
}
func (ϟa *GlProgramUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_995_major := uint32(3) // u32
	minRequiredVersion_995_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_995_major, minRequiredVersion_995_minor
	return nil
}
func (ϟa *GlProgramUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_996_major := uint32(3) // u32
	minRequiredVersion_996_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_996_major, minRequiredVersion_996_minor
	return nil
}
func (ϟa *GlProgramUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_997_major := uint32(3) // u32
	minRequiredVersion_997_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_997_major, minRequiredVersion_997_minor
	return nil
}
func (ϟa *GlProgramUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_998_major := uint32(3) // u32
	minRequiredVersion_998_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_998_major, minRequiredVersion_998_minor
	return nil
}
func (ϟa *GlProgramUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_999_major := uint32(3) // u32
	minRequiredVersion_999_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_999_major, minRequiredVersion_999_minor
	return nil
}
func (ϟa *GlProgramUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1000_major := uint32(3) // u32
	minRequiredVersion_1000_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1000_major, minRequiredVersion_1000_minor
	return nil
}
func (ϟa *GlProgramUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1001_major := uint32(3) // u32
	minRequiredVersion_1001_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1001_major, minRequiredVersion_1001_minor
	return nil
}
func (ϟa *GlProgramUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1002_major := uint32(3) // u32
	minRequiredVersion_1002_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1002_major, minRequiredVersion_1002_minor
	return nil
}
func (ϟa *GlProgramUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1003_major := uint32(3) // u32
	minRequiredVersion_1003_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1003_major, minRequiredVersion_1003_minor
	return nil
}
func (ϟa *GlProgramUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1004_major := uint32(3) // u32
	minRequiredVersion_1004_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1004_major, minRequiredVersion_1004_minor
	return nil
}
func (ϟa *GlProgramUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1005_major := uint32(3) // u32
	minRequiredVersion_1005_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1005_major, minRequiredVersion_1005_minor
	return nil
}
func (ϟa *GlProgramUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1006_major := uint32(3) // u32
	minRequiredVersion_1006_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1006_major, minRequiredVersion_1006_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1007_major := uint32(3) // u32
	minRequiredVersion_1007_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1007_major, minRequiredVersion_1007_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1008_major := uint32(3) // u32
	minRequiredVersion_1008_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1008_major, minRequiredVersion_1008_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1009_major := uint32(3) // u32
	minRequiredVersion_1009_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1009_major, minRequiredVersion_1009_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1010_major := uint32(3) // u32
	minRequiredVersion_1010_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1010_major, minRequiredVersion_1010_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1011_major := uint32(3) // u32
	minRequiredVersion_1011_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1011_major, minRequiredVersion_1011_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1012_major := uint32(3) // u32
	minRequiredVersion_1012_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1012_major, minRequiredVersion_1012_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1013_major := uint32(3) // u32
	minRequiredVersion_1013_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1013_major, minRequiredVersion_1013_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1014_major := uint32(3) // u32
	minRequiredVersion_1014_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1014_major, minRequiredVersion_1014_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1015_major := uint32(3) // u32
	minRequiredVersion_1015_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1015_major, minRequiredVersion_1015_minor
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1016_major := uint32(2) // u32
	minRequiredVersion_1016_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1016_major, minRequiredVersion_1016_minor
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1017_major := uint32(2) // u32
	minRequiredVersion_1017_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		glErrorInvalidEnum_1018_param := ϟa.BinaryFormat // GLenum
		return
		_ = glErrorInvalidEnum_1018_param
	}
	ϟa.Shaders.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.BinarySize), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1017_major, minRequiredVersion_1017_minor
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1019_major := uint32(2)                                  // u32
	minRequiredVersion_1019_minor := uint32(0)                                  // u32
	sources := ϟa.Source.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	lengths := ϟa.Length.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1021_msg := "No context bound" // string
		return
		_ = error_1021_msg
	}
	GetContext_1020_result := context         // Contextʳ
	ctx := GetContext_1020_result             // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader) // Shaderʳ
	s.Source = ""
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1019_major, minRequiredVersion_1019_minor, sources, lengths, context, GetContext_1020_result, ctx, s
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1022_major := uint32(2)   // u32
	minRequiredVersion_1022_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1024_msg := "No context bound" // string
		return
		_ = error_1024_msg
	}
	GetContext_1023_result := context // Contextʳ
	ctx := GetContext_1023_result     // Contextʳ
	v := MakeGLfloatˢ(uint64(1), ϟs)  // GLfloatˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1022_major, minRequiredVersion_1022_minor, context, GetContext_1023_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1025_major := uint32(2)   // u32
	minRequiredVersion_1025_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1027_msg := "No context bound" // string
		return
		_ = error_1027_msg
	}
	GetContext_1026_result := context                                     // Contextʳ
	ctx := GetContext_1026_result                                         // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLfloatˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1025_major, minRequiredVersion_1025_minor, context, GetContext_1026_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1028_major := uint32(2)   // u32
	minRequiredVersion_1028_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1030_msg := "No context bound" // string
		return
		_ = error_1030_msg
	}
	GetContext_1029_result := context // Contextʳ
	ctx := GetContext_1029_result     // Contextʳ
	v := MakeGLintˢ(uint64(1), ϟs)    // GLintˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1028_major, minRequiredVersion_1028_minor, context, GetContext_1029_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1031_major := uint32(2)   // u32
	minRequiredVersion_1031_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1033_msg := "No context bound" // string
		return
		_ = error_1033_msg
	}
	GetContext_1032_result := context                                     // Contextʳ
	ctx := GetContext_1032_result                                         // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1031_major, minRequiredVersion_1031_minor, context, GetContext_1032_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1034_major := uint32(3) // u32
	minRequiredVersion_1034_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1034_major, minRequiredVersion_1034_minor
	return nil
}
func (ϟa *GlUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1035_major := uint32(3) // u32
	minRequiredVersion_1035_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1035_major, minRequiredVersion_1035_minor
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1036_major := uint32(2)   // u32
	minRequiredVersion_1036_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1038_msg := "No context bound" // string
		return
		_ = error_1038_msg
	}
	GetContext_1037_result := context // Contextʳ
	ctx := GetContext_1037_result     // Contextʳ
	v := MakeVec2fˢ(uint64(1), ϟs)    // Vec2fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2f{Elements: [2]GLfloat{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1036_major, minRequiredVersion_1036_minor, context, GetContext_1037_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1039_major := uint32(2)   // u32
	minRequiredVersion_1039_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1041_msg := "No context bound" // string
		return
		_ = error_1041_msg
	}
	GetContext_1040_result := context                                             // Contextʳ
	ctx := GetContext_1040_result                                                 // Contextʳ
	v := Vec2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1039_major, minRequiredVersion_1039_minor, context, GetContext_1040_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1042_major := uint32(2)   // u32
	minRequiredVersion_1042_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1044_msg := "No context bound" // string
		return
		_ = error_1044_msg
	}
	GetContext_1043_result := context // Contextʳ
	ctx := GetContext_1043_result     // Contextʳ
	v := MakeVec2iˢ(uint64(1), ϟs)    // Vec2iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2i{Elements: [2]GLint{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1042_major, minRequiredVersion_1042_minor, context, GetContext_1043_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1045_major := uint32(2)   // u32
	minRequiredVersion_1045_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1047_msg := "No context bound" // string
		return
		_ = error_1047_msg
	}
	GetContext_1046_result := context                                             // Contextʳ
	ctx := GetContext_1046_result                                                 // Contextʳ
	v := Vec2iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1045_major, minRequiredVersion_1045_minor, context, GetContext_1046_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1048_major := uint32(3) // u32
	minRequiredVersion_1048_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1048_major, minRequiredVersion_1048_minor
	return nil
}
func (ϟa *GlUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1049_major := uint32(3) // u32
	minRequiredVersion_1049_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1049_major, minRequiredVersion_1049_minor
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1050_major := uint32(2)   // u32
	minRequiredVersion_1050_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1052_msg := "No context bound" // string
		return
		_ = error_1052_msg
	}
	GetContext_1051_result := context // Contextʳ
	ctx := GetContext_1051_result     // Contextʳ
	v := MakeVec3fˢ(uint64(1), ϟs)    // Vec3fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3f{Elements: [3]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1050_major, minRequiredVersion_1050_minor, context, GetContext_1051_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1053_major := uint32(2)   // u32
	minRequiredVersion_1053_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1055_msg := "No context bound" // string
		return
		_ = error_1055_msg
	}
	GetContext_1054_result := context                                             // Contextʳ
	ctx := GetContext_1054_result                                                 // Contextʳ
	v := Vec3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1053_major, minRequiredVersion_1053_minor, context, GetContext_1054_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1056_major := uint32(2)   // u32
	minRequiredVersion_1056_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1058_msg := "No context bound" // string
		return
		_ = error_1058_msg
	}
	GetContext_1057_result := context // Contextʳ
	ctx := GetContext_1057_result     // Contextʳ
	v := MakeVec3iˢ(uint64(1), ϟs)    // Vec3iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3i{Elements: [3]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1056_major, minRequiredVersion_1056_minor, context, GetContext_1057_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1059_major := uint32(2)   // u32
	minRequiredVersion_1059_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1061_msg := "No context bound" // string
		return
		_ = error_1061_msg
	}
	GetContext_1060_result := context                                             // Contextʳ
	ctx := GetContext_1060_result                                                 // Contextʳ
	v := Vec3iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1059_major, minRequiredVersion_1059_minor, context, GetContext_1060_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1062_major := uint32(3) // u32
	minRequiredVersion_1062_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1062_major, minRequiredVersion_1062_minor
	return nil
}
func (ϟa *GlUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1063_major := uint32(3) // u32
	minRequiredVersion_1063_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1063_major, minRequiredVersion_1063_minor
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1064_major := uint32(2)   // u32
	minRequiredVersion_1064_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1066_msg := "No context bound" // string
		return
		_ = error_1066_msg
	}
	GetContext_1065_result := context // Contextʳ
	ctx := GetContext_1065_result     // Contextʳ
	v := MakeVec4fˢ(uint64(1), ϟs)    // Vec4fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4f{Elements: [4]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1064_major, minRequiredVersion_1064_minor, context, GetContext_1065_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1067_major := uint32(2)   // u32
	minRequiredVersion_1067_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1069_msg := "No context bound" // string
		return
		_ = error_1069_msg
	}
	GetContext_1068_result := context                                             // Contextʳ
	ctx := GetContext_1068_result                                                 // Contextʳ
	v := Vec4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1067_major, minRequiredVersion_1067_minor, context, GetContext_1068_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1070_major := uint32(2)   // u32
	minRequiredVersion_1070_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1072_msg := "No context bound" // string
		return
		_ = error_1072_msg
	}
	GetContext_1071_result := context // Contextʳ
	ctx := GetContext_1071_result     // Contextʳ
	v := MakeVec4iˢ(uint64(1), ϟs)    // Vec4iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4i{Elements: [4]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_1070_major, minRequiredVersion_1070_minor, context, GetContext_1071_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1073_major := uint32(2)   // u32
	minRequiredVersion_1073_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1075_msg := "No context bound" // string
		return
		_ = error_1075_msg
	}
	GetContext_1074_result := context                                             // Contextʳ
	ctx := GetContext_1074_result                                                 // Contextʳ
	v := Vec4iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1073_major, minRequiredVersion_1073_minor, context, GetContext_1074_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1076_major := uint32(3) // u32
	minRequiredVersion_1076_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1076_major, minRequiredVersion_1076_minor
	return nil
}
func (ϟa *GlUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1077_major := uint32(3) // u32
	minRequiredVersion_1077_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1077_major, minRequiredVersion_1077_minor
	return nil
}
func (ϟa *GlUniformBlockBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1078_major := uint32(3) // u32
	minRequiredVersion_1078_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1078_major, minRequiredVersion_1078_minor
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1079_major := uint32(2)   // u32
	minRequiredVersion_1079_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1081_msg := "No context bound" // string
		return
		_ = error_1081_msg
	}
	GetContext_1080_result := context                                             // Contextʳ
	ctx := GetContext_1080_result                                                 // Contextʳ
	v := Mat2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1079_major, minRequiredVersion_1079_minor, context, GetContext_1080_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1082_major := uint32(3) // u32
	minRequiredVersion_1082_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1082_major, minRequiredVersion_1082_minor
	return nil
}
func (ϟa *GlUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1083_major := uint32(3) // u32
	minRequiredVersion_1083_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1083_major, minRequiredVersion_1083_minor
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1084_major := uint32(2)   // u32
	minRequiredVersion_1084_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1086_msg := "No context bound" // string
		return
		_ = error_1086_msg
	}
	GetContext_1085_result := context                                             // Contextʳ
	ctx := GetContext_1085_result                                                 // Contextʳ
	v := Mat3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1084_major, minRequiredVersion_1084_minor, context, GetContext_1085_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1087_major := uint32(3) // u32
	minRequiredVersion_1087_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1087_major, minRequiredVersion_1087_minor
	return nil
}
func (ϟa *GlUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1088_major := uint32(3) // u32
	minRequiredVersion_1088_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1088_major, minRequiredVersion_1088_minor
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1089_major := uint32(2)   // u32
	minRequiredVersion_1089_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1091_msg := "No context bound" // string
		return
		_ = error_1091_msg
	}
	GetContext_1090_result := context                                             // Contextʳ
	ctx := GetContext_1090_result                                                 // Contextʳ
	v := Mat4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1089_major, minRequiredVersion_1089_minor, context, GetContext_1090_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1092_major := uint32(3) // u32
	minRequiredVersion_1092_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1092_major, minRequiredVersion_1092_minor
	return nil
}
func (ϟa *GlUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1093_major := uint32(3) // u32
	minRequiredVersion_1093_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1093_major, minRequiredVersion_1093_minor
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1094_major := uint32(2)   // u32
	minRequiredVersion_1094_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1096_msg := "No context bound" // string
		return
		_ = error_1096_msg
	}
	GetContext_1095_result := context // Contextʳ
	ctx := GetContext_1095_result     // Contextʳ
	ctx.BoundProgram = ϟa.Program
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1094_major, minRequiredVersion_1094_minor, context, GetContext_1095_result, ctx
	return nil
}
func (ϟa *GlUseProgramStages) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1097_major := uint32(3)                                                                                                                                        // u32
	minRequiredVersion_1097_minor := uint32(1)                                                                                                                                        // u32
	supportsBits_1098_seenBits := ϟa.Stages                                                                                                                                           // GLbitfield
	supportsBits_1098_validBits := (GLbitfield_GL_ALL_SHADER_BITS) | ((GLbitfield_GL_COMPUTE_SHADER_BIT) | ((GLbitfield_GL_FRAGMENT_SHADER_BIT) | (GLbitfield_GL_VERTEX_SHADER_BIT))) // GLbitfield
	if (GLbitfield_GL_ALL_SHADER_BITS)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_COMPUTE_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_FRAGMENT_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_VERTEX_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_1097_major, minRequiredVersion_1097_minor, supportsBits_1098_seenBits, supportsBits_1098_validBits
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1099_major := uint32(2) // u32
	minRequiredVersion_1099_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1099_major, minRequiredVersion_1099_minor
	return nil
}
func (ϟa *GlValidateProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1100_major := uint32(3) // u32
	minRequiredVersion_1100_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1100_major, minRequiredVersion_1100_minor
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1101_major := uint32(2) // u32
	minRequiredVersion_1101_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		glErrorInvalidEnum_1102_param := ϟa.Mode // GLenum
		return
		_ = glErrorInvalidEnum_1102_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1104_msg := "No context bound" // string
		return
		_ = error_1104_msg
	}
	GetContext_1103_result := context // Contextʳ
	ctx := GetContext_1103_result     // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1101_major, minRequiredVersion_1101_minor, context, GetContext_1103_result, ctx
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1105_major := uint32(2)   // u32
	minRequiredVersion_1105_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1107_msg := "No context bound" // string
		return
		_ = error_1107_msg
	}
	GetContext_1106_result := context // Contextʳ
	ctx := GetContext_1106_result     // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1105_major, minRequiredVersion_1105_minor, context, GetContext_1106_result, ctx
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1108_major := uint32(2) // u32
	minRequiredVersion_1108_minor := uint32(0) // u32
	switch ϟa.Orientation {
	case GLenum_GL_CCW, GLenum_GL_CW:
	default:
		glErrorInvalidEnum_1109_param := ϟa.Orientation // GLenum
		return
		_ = glErrorInvalidEnum_1109_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1111_msg := "No context bound" // string
		return
		_ = error_1111_msg
	}
	GetContext_1110_result := context // Contextʳ
	ctx := GetContext_1110_result     // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1108_major, minRequiredVersion_1108_minor, context, GetContext_1110_result, ctx
	return nil
}
func (ϟa *GlGetMultisamplefv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1112_major := uint32(3) // u32
	minRequiredVersion_1112_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_SAMPLE_POSITION:
	default:
		glErrorInvalidEnum_1113_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1113_param
	}
	ϟa.Val.Slice(uint64(0), uint64(2), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1112_major, minRequiredVersion_1112_minor
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1114_major := uint32(2)   // u32
	minRequiredVersion_1114_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1116_msg := "No context bound" // string
		return
		_ = error_1116_msg
	}
	GetContext_1115_result := context // Contextʳ
	ctx := GetContext_1115_result     // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1114_major, minRequiredVersion_1114_minor, context, GetContext_1115_result, ctx
	return nil
}
func (ϟa *GlMinSampleShading) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1117_major := uint32(3) // u32
	minRequiredVersion_1117_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1117_major, minRequiredVersion_1117_minor
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1118_major := uint32(2)   // u32
	minRequiredVersion_1118_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1120_msg := "No context bound" // string
		return
		_ = error_1120_msg
	}
	GetContext_1119_result := context // Contextʳ
	ctx := GetContext_1119_result     // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1118_major, minRequiredVersion_1118_minor, context, GetContext_1119_result, ctx
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1121_major := uint32(2)   // u32
	minRequiredVersion_1121_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1123_msg := "No context bound" // string
		return
		_ = error_1123_msg
	}
	GetContext_1122_result := context // Contextʳ
	ctx := GetContext_1122_result     // Contextʳ
	ctx.Rasterizing.Viewport = Rect{X: ϟa.X, Y: ϟa.Y, Width: ϟa.Width, Height: ϟa.Height}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1121_major, minRequiredVersion_1121_minor, context, GetContext_1122_result, ctx
	return nil
}
func (ϟa *GlGetBooleani_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1124_major := uint32(3) // u32
	minRequiredVersion_1124_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE, GLenum_GL_VIEWPORT:
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
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
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1127_major := uint32(2) // u32
	minRequiredVersion_1127_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_1128_major := uint32(3) // u32
		minRequiredVersion_1128_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1128_major, minRequiredVersion_1128_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1129_major := uint32(3) // u32
		minRequiredVersion_1129_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1129_major, minRequiredVersion_1129_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1130_major := uint32(3) // u32
		minRequiredVersion_1130_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1130_major, minRequiredVersion_1130_minor
	default:
		glErrorInvalidEnum_1131_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1131_param
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLbooleanˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1133_msg := "No context bound" // string
		return
		_ = error_1133_msg
	}
	GetContext_1132_result := context // Contextʳ
	ctx := GetContext_1132_result     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1127_major, minRequiredVersion_1127_minor, v, context, GetContext_1132_result, ctx
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1134_major := uint32(2) // u32
	minRequiredVersion_1134_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_1135_major := uint32(3) // u32
		minRequiredVersion_1135_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1135_major, minRequiredVersion_1135_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1136_major := uint32(3) // u32
		minRequiredVersion_1136_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1136_major, minRequiredVersion_1136_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1137_major := uint32(3) // u32
		minRequiredVersion_1137_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1137_major, minRequiredVersion_1137_minor
	default:
		glErrorInvalidEnum_1138_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1138_param
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLfloatˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1140_msg := "No context bound" // string
		return
		_ = error_1140_msg
	}
	GetContext_1139_result := context // Contextʳ
	ctx := GetContext_1139_result     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1134_major, minRequiredVersion_1134_minor, v, context, GetContext_1139_result, ctx
	return nil
}
func (ϟa *GlGetInteger64i_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1141_major := uint32(3) // u32
	minRequiredVersion_1141_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1142_major := uint32(3) // u32
		minRequiredVersion_1142_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1142_major, minRequiredVersion_1142_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1143_major := uint32(3) // u32
		minRequiredVersion_1143_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1143_major, minRequiredVersion_1143_minor
	default:
		glErrorInvalidEnum_1144_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1144_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1141_major, minRequiredVersion_1141_minor
	return nil
}
func (ϟa *GlGetInteger64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1145_major := uint32(3) // u32
	minRequiredVersion_1145_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1146_major := uint32(3) // u32
		minRequiredVersion_1146_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1146_major, minRequiredVersion_1146_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1147_major := uint32(3) // u32
		minRequiredVersion_1147_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1147_major, minRequiredVersion_1147_minor
	default:
		glErrorInvalidEnum_1148_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1148_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1145_major, minRequiredVersion_1145_minor
	return nil
}
func (ϟa *GlGetIntegeri_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1149_major := uint32(3) // u32
	minRequiredVersion_1149_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1150_major := uint32(3) // u32
		minRequiredVersion_1150_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1150_major, minRequiredVersion_1150_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1151_major := uint32(3) // u32
		minRequiredVersion_1151_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1151_major, minRequiredVersion_1151_minor
	default:
		glErrorInvalidEnum_1152_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1152_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1149_major, minRequiredVersion_1149_minor
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1153_major := uint32(2) // u32
	minRequiredVersion_1153_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_1154_major := uint32(3) // u32
		minRequiredVersion_1154_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1154_major, minRequiredVersion_1154_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_1155_major := uint32(3) // u32
		minRequiredVersion_1155_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1155_major, minRequiredVersion_1155_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_1156_major := uint32(3) // u32
		minRequiredVersion_1156_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1156_major, minRequiredVersion_1156_minor
	default:
		glErrorInvalidEnum_1157_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1157_param
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1159_msg := "No context bound" // string
		return
		_ = error_1159_msg
	}
	GetContext_1158_result := context // Contextʳ
	ctx := GetContext_1158_result     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1153_major, minRequiredVersion_1153_minor, v, context, GetContext_1158_result, ctx
	return nil
}
func (ϟa *GlGetInternalformativ) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1160_major := uint32(3) // u32
	minRequiredVersion_1160_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1161_major := uint32(3) // u32
		minRequiredVersion_1161_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1161_major, minRequiredVersion_1161_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY:
		minRequiredVersion_1162_major := uint32(3) // u32
		minRequiredVersion_1162_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1162_major, minRequiredVersion_1162_minor
	default:
		glErrorInvalidEnum_1163_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1163_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1164_major := uint32(3) // u32
		minRequiredVersion_1164_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1164_major, minRequiredVersion_1164_minor
	default:
		glErrorInvalidEnum_1165_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1165_param
	}
	switch ϟa.Pname {
	case GLenum_GL_NUM_SAMPLE_COUNTS, GLenum_GL_SAMPLES:
	default:
		glErrorInvalidEnum_1166_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1166_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1160_major, minRequiredVersion_1160_minor
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1167_major := uint32(2) // u32
	minRequiredVersion_1167_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_EXTENSIONS, GLenum_GL_RENDERER, GLenum_GL_SHADING_LANGUAGE_VERSION, GLenum_GL_VENDOR, GLenum_GL_VERSION:
	default:
		glErrorInvalidEnum_1168_param := ϟa.Param // GLenum
		return
		_ = glErrorInvalidEnum_1168_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = GLubyteᶜᵖ{}
	_, _ = minRequiredVersion_1167_major, minRequiredVersion_1167_minor
	return nil
}
func (ϟa *GlGetStringi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1169_major := uint32(3) // u32
	minRequiredVersion_1169_minor := uint32(0) // u32
	switch ϟa.Name {
	case GLenum_GL_EXTENSIONS:
	default:
		glErrorInvalidEnum_1170_param := ϟa.Name // GLenum
		return
		_ = glErrorInvalidEnum_1170_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1169_major, minRequiredVersion_1169_minor
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1171_major := uint32(2) // u32
	minRequiredVersion_1171_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_1172_major := uint32(3) // u32
		minRequiredVersion_1172_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1172_major, minRequiredVersion_1172_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_1173_major := uint32(3) // u32
		minRequiredVersion_1173_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1173_major, minRequiredVersion_1173_minor
	default:
		glErrorInvalidEnum_1174_param := ϟa.Capability // GLenum
		return
		_ = glErrorInvalidEnum_1174_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1176_msg := "No context bound" // string
		return
		_ = error_1176_msg
	}
	GetContext_1175_result := context // Contextʳ
	ctx := GetContext_1175_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Capabilities.Get(ϟa.Capability) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_1171_major, minRequiredVersion_1171_minor, context, GetContext_1175_result, ctx
	return nil
}
func (ϟa *GlIsEnabledi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1177_major := uint32(3) // u32
	minRequiredVersion_1177_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		glErrorInvalidEnum_1178_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1178_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1177_major, minRequiredVersion_1177_minor
	return nil
}
func (ϟa *GlClientWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1179_major := uint32(3)                           // u32
	minRequiredVersion_1179_minor := uint32(0)                           // u32
	supportsBits_1180_seenBits := ϟa.SyncFlags                           // GLbitfield
	supportsBits_1180_validBits := GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT // GLbitfield
	if (GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT)&(ϟa.SyncFlags) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _, _, _ = minRequiredVersion_1179_major, minRequiredVersion_1179_minor, supportsBits_1180_seenBits, supportsBits_1180_validBits
	return nil
}
func (ϟa *GlDeleteSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1181_major := uint32(3) // u32
	minRequiredVersion_1181_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1181_major, minRequiredVersion_1181_minor
	return nil
}
func (ϟa *GlFenceSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1182_major := uint32(3) // u32
	minRequiredVersion_1182_minor := uint32(0) // u32
	switch ϟa.Condition {
	case GLenum_GL_SYNC_GPU_COMMANDS_COMPLETE:
	default:
		glErrorInvalidEnum_1183_param := ϟa.Condition // GLenum
		return
		_ = glErrorInvalidEnum_1183_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1182_major, minRequiredVersion_1182_minor
	return nil
}
func (ϟa *GlGetSynciv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1184_major := uint32(3) // u32
	minRequiredVersion_1184_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Pname {
	case GLenum_GL_OBJECT_TYPE, GLenum_GL_SYNC_CONDITION, GLenum_GL_SYNC_FLAGS, GLenum_GL_SYNC_STATUS:
		if ((ϟa.Values) != (GLintᵖ{})) && ((ϟa.BufSize) > (GLsizei(int32(0)))) {
			ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
			if (ϟa.Length) != (GLsizeiᵖ{}) {
				ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(int32(1)), ϟa, ϟs, ϟd, ϟl, ϟb)
			}
		}
	default:
		glErrorInvalidEnum_1185_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1185_param
	}
	_, _ = minRequiredVersion_1184_major, minRequiredVersion_1184_minor
	return nil
}
func (ϟa *GlIsSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1186_major := uint32(3) // u32
	minRequiredVersion_1186_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1186_major, minRequiredVersion_1186_minor
	return nil
}
func (ϟa *GlWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1187_major := uint32(3) // u32
	minRequiredVersion_1187_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1187_major, minRequiredVersion_1187_minor
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1188_major := uint32(2) // u32
	minRequiredVersion_1188_minor := uint32(0) // u32
	switch ϟa.Unit {
	case GLenum_GL_TEXTURE0, GLenum_GL_TEXTURE1, GLenum_GL_TEXTURE10, GLenum_GL_TEXTURE11, GLenum_GL_TEXTURE12, GLenum_GL_TEXTURE13, GLenum_GL_TEXTURE14, GLenum_GL_TEXTURE15, GLenum_GL_TEXTURE16, GLenum_GL_TEXTURE17, GLenum_GL_TEXTURE18, GLenum_GL_TEXTURE19, GLenum_GL_TEXTURE2, GLenum_GL_TEXTURE20, GLenum_GL_TEXTURE21, GLenum_GL_TEXTURE22, GLenum_GL_TEXTURE23, GLenum_GL_TEXTURE24, GLenum_GL_TEXTURE25, GLenum_GL_TEXTURE26, GLenum_GL_TEXTURE27, GLenum_GL_TEXTURE28, GLenum_GL_TEXTURE29, GLenum_GL_TEXTURE3, GLenum_GL_TEXTURE30, GLenum_GL_TEXTURE31, GLenum_GL_TEXTURE4, GLenum_GL_TEXTURE5, GLenum_GL_TEXTURE6, GLenum_GL_TEXTURE7, GLenum_GL_TEXTURE8, GLenum_GL_TEXTURE9:
	default:
		glErrorInvalidEnum_1189_param := ϟa.Unit // GLenum
		return
		_ = glErrorInvalidEnum_1189_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1191_msg := "No context bound" // string
		return
		_ = error_1191_msg
	}
	GetContext_1190_result := context // Contextʳ
	ctx := GetContext_1190_result     // Contextʳ
	ctx.ActiveTextureUnit = ϟa.Unit
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1188_major, minRequiredVersion_1188_minor, context, GetContext_1190_result, ctx
	return nil
}
func (ϟa *GlBindImageTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1192_major := uint32(3) // u32
	minRequiredVersion_1192_minor := uint32(1) // u32
	switch ϟa.Access {
	case GLenum_GL_READ_ONLY, GLenum_GL_READ_WRITE, GLenum_GL_WRITE_ONLY:
	default:
		glErrorInvalidEnum_1193_param := ϟa.Access // GLenum
		return
		_ = glErrorInvalidEnum_1193_param
	}
	switch ϟa.Format {
	case GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM:
	default:
		glErrorInvalidEnum_1194_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1194_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1192_major, minRequiredVersion_1192_minor
	return nil
}
func (ϟa *GlBindSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1195_major := uint32(3) // u32
	minRequiredVersion_1195_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1195_major, minRequiredVersion_1195_minor
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1196_major := uint32(2) // u32
	minRequiredVersion_1196_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1197_major := uint32(3) // u32
		minRequiredVersion_1197_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1197_major, minRequiredVersion_1197_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1198_major := uint32(3) // u32
		minRequiredVersion_1198_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1198_major, minRequiredVersion_1198_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1199_major := uint32(3) // u32
		minRequiredVersion_1199_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1199_major, minRequiredVersion_1199_minor
	default:
		glErrorInvalidEnum_1200_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1200_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1202_msg := "No context bound" // string
		return
		_ = error_1202_msg
	}
	GetContext_1201_result := context // Contextʳ
	ctx := GetContext_1201_result     // Contextʳ
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = (&Texture{ID: ϟa.Texture, Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	}
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	tu.Bindings[ϟa.Target] = ϟa.Texture
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1196_major, minRequiredVersion_1196_minor, context, GetContext_1201_result, ctx, tu
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1203_major := uint32(2) // u32
	minRequiredVersion_1203_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1204_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1204_param
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_1205_major := uint32(3) // u32
		minRequiredVersion_1205_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1205_major, minRequiredVersion_1205_minor
	case GLenum_GL_ATC_RGB_AMD, GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD, GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD:
		requiresExtension_1206_ext := ExtensionId_GL_AMD_compressed_ATC_texture // ExtensionId
		_ = requiresExtension_1206_ext
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1207_major := uint32(3) // u32
		minRequiredVersion_1207_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1207_major, minRequiredVersion_1207_minor
	default:
		glErrorInvalidEnum_1208_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1208_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1210_msg := "No context bound" // string
		return
		_ = error_1210_msg
	}
	GetContext_1209_result := context                 // Contextʳ
	ctx := GetContext_1209_result                     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1203_major, minRequiredVersion_1203_minor, context, GetContext_1209_result, ctx, tu
	return nil
}
func (ϟa *GlCompressedTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1211_major := uint32(3) // u32
	minRequiredVersion_1211_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1212_major := uint32(3) // u32
		minRequiredVersion_1212_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1212_major, minRequiredVersion_1212_minor
	default:
		glErrorInvalidEnum_1213_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1213_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1214_major := uint32(3) // u32
		minRequiredVersion_1214_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1214_major, minRequiredVersion_1214_minor
	default:
		glErrorInvalidEnum_1215_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1215_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1217_msg := "No context bound" // string
		return
		_ = error_1217_msg
	}
	GetContext_1216_result := context // Contextʳ
	ctx := GetContext_1216_result     // Contextʳ
	if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
		U8ᵖ(ϟa.Data).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.ImageSize), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1211_major, minRequiredVersion_1211_minor, context, GetContext_1216_result, ctx
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1218_major := uint32(2) // u32
	minRequiredVersion_1218_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1219_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1219_param
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_1220_major := uint32(3) // u32
		minRequiredVersion_1220_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1220_major, minRequiredVersion_1220_minor
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1221_major := uint32(3) // u32
		minRequiredVersion_1221_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1221_major, minRequiredVersion_1221_minor
	default:
		glErrorInvalidEnum_1222_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1222_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1224_msg := "No context bound" // string
		return
		_ = error_1224_msg
	}
	GetContext_1223_result := context // Contextʳ
	ctx := GetContext_1223_result     // Contextʳ
	if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
		U8ᵖ(ϟa.Data).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.ImageSize), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1218_major, minRequiredVersion_1218_minor, context, GetContext_1223_result, ctx
	return nil
}
func (ϟa *GlCompressedTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1225_major := uint32(3) // u32
	minRequiredVersion_1225_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1226_major := uint32(3) // u32
		minRequiredVersion_1226_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1226_major, minRequiredVersion_1226_minor
	default:
		glErrorInvalidEnum_1227_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1227_param
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1228_major := uint32(3) // u32
		minRequiredVersion_1228_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1228_major, minRequiredVersion_1228_minor
	default:
		glErrorInvalidEnum_1229_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1229_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1231_msg := "No context bound" // string
		return
		_ = error_1231_msg
	}
	GetContext_1230_result := context // Contextʳ
	ctx := GetContext_1230_result     // Contextʳ
	if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
		U8ᵖ(ϟa.Data).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.ImageSize), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1225_major, minRequiredVersion_1225_minor, context, GetContext_1230_result, ctx
	return nil
}
func (ϟa *GlCopyImageSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1232_major := uint32(3) // u32
	minRequiredVersion_1232_minor := uint32(2) // u32
	switch ϟa.SrcTarget {
	case GLenum_GL_RENDERBUFFER, GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1233_param := ϟa.SrcTarget // GLenum
		return
		_ = glErrorInvalidEnum_1233_param
	}
	switch ϟa.DstTarget {
	case GLenum_GL_RENDERBUFFER, GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		glErrorInvalidEnum_1234_param := ϟa.DstTarget // GLenum
		return
		_ = glErrorInvalidEnum_1234_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1232_major, minRequiredVersion_1232_minor
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1235_major := uint32(2) // u32
	minRequiredVersion_1235_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1236_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1236_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_1237_major := uint32(3) // u32
		minRequiredVersion_1237_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1237_major, minRequiredVersion_1237_minor
	default:
		glErrorInvalidEnum_1238_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1238_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1235_major, minRequiredVersion_1235_minor
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1239_major := uint32(2) // u32
	minRequiredVersion_1239_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1240_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1240_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1239_major, minRequiredVersion_1239_minor
	return nil
}
func (ϟa *GlCopyTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1241_major := uint32(3) // u32
	minRequiredVersion_1241_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1242_major := uint32(3) // u32
		minRequiredVersion_1242_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1242_major, minRequiredVersion_1242_minor
	default:
		glErrorInvalidEnum_1243_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1243_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1241_major, minRequiredVersion_1241_minor
	return nil
}
func (ϟa *GlDeleteSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1244_major := uint32(3) // u32
	minRequiredVersion_1244_minor := uint32(0) // u32
	ϟa.Samplers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1244_major, minRequiredVersion_1244_minor
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1245_major := uint32(2)                              // u32
	minRequiredVersion_1245_minor := uint32(0)                              // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1247_msg := "No context bound" // string
		return
		_ = error_1247_msg
	}
	GetContext_1246_result := context // Contextʳ
	ctx := GetContext_1246_result     // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Textures, t.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1245_major, minRequiredVersion_1245_minor, t, context, GetContext_1246_result, ctx
	return nil
}
func (ϟa *GlGenSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1248_major := uint32(3) // u32
	minRequiredVersion_1248_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Samplers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_1248_major, minRequiredVersion_1248_minor
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1249_major := uint32(2)                              // u32
	minRequiredVersion_1249_minor := uint32(0)                              // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1251_msg := "No context bound" // string
		return
		_ = error_1251_msg
	}
	GetContext_1250_result := context // Contextʳ
	ctx := GetContext_1250_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // TextureId
		ctx.Instances.Textures[id] = (&Texture{ID: id, Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
		t.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1249_major, minRequiredVersion_1249_minor, t, context, GetContext_1250_result, ctx
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1252_major := uint32(2) // u32
	minRequiredVersion_1252_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1253_major := uint32(3) // u32
		minRequiredVersion_1253_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1253_major, minRequiredVersion_1253_minor
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1254_major := uint32(3) // u32
		minRequiredVersion_1254_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1254_major, minRequiredVersion_1254_minor
	default:
		glErrorInvalidEnum_1255_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1255_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1252_major, minRequiredVersion_1252_minor
	return nil
}
func (ϟa *GlGetSamplerParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1256_major := uint32(3)     // u32
	minRequiredVersion_1256_minor := uint32(2)     // u32
	GetSamplerParameter_1257_sampler := ϟa.Sampler // SamplerId
	GetSamplerParameter_1257_pname := ϟa.Pname     // GLenum
	GetSamplerParameter_1257_params := ϟa.Params   // GLintᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetSamplerParameter_1257_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetSamplerParameter_1257_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		GetSamplerParameter_1257_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		glErrorInvalidEnum_1258_param := GetSamplerParameter_1257_pname // GLenum
		return
		_ = glErrorInvalidEnum_1258_param
	}
	_, _, _, _, _ = minRequiredVersion_1256_major, minRequiredVersion_1256_minor, GetSamplerParameter_1257_sampler, GetSamplerParameter_1257_pname, GetSamplerParameter_1257_params
	return nil
}
func (ϟa *GlGetSamplerParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1259_major := uint32(3)     // u32
	minRequiredVersion_1259_minor := uint32(2)     // u32
	GetSamplerParameter_1260_sampler := ϟa.Sampler // SamplerId
	GetSamplerParameter_1260_pname := ϟa.Pname     // GLenum
	GetSamplerParameter_1260_params := ϟa.Params   // GLuintᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetSamplerParameter_1260_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetSamplerParameter_1260_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		GetSamplerParameter_1260_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		glErrorInvalidEnum_1261_param := GetSamplerParameter_1260_pname // GLenum
		return
		_ = glErrorInvalidEnum_1261_param
	}
	_, _, _, _, _ = minRequiredVersion_1259_major, minRequiredVersion_1259_minor, GetSamplerParameter_1260_sampler, GetSamplerParameter_1260_pname, GetSamplerParameter_1260_params
	return nil
}
func (ϟa *GlGetSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1262_major := uint32(3)     // u32
	minRequiredVersion_1262_minor := uint32(0)     // u32
	GetSamplerParameter_1263_sampler := ϟa.Sampler // SamplerId
	GetSamplerParameter_1263_pname := ϟa.Pname     // GLenum
	GetSamplerParameter_1263_params := ϟa.Params   // GLfloatᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetSamplerParameter_1263_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetSamplerParameter_1263_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		GetSamplerParameter_1263_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		glErrorInvalidEnum_1264_param := GetSamplerParameter_1263_pname // GLenum
		return
		_ = glErrorInvalidEnum_1264_param
	}
	_, _, _, _, _ = minRequiredVersion_1262_major, minRequiredVersion_1262_minor, GetSamplerParameter_1263_sampler, GetSamplerParameter_1263_pname, GetSamplerParameter_1263_params
	return nil
}
func (ϟa *GlGetSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1265_major := uint32(3)     // u32
	minRequiredVersion_1265_minor := uint32(0)     // u32
	GetSamplerParameter_1266_sampler := ϟa.Sampler // SamplerId
	GetSamplerParameter_1266_pname := ϟa.Pname     // GLenum
	GetSamplerParameter_1266_params := ϟa.Params   // GLintᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetSamplerParameter_1266_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetSamplerParameter_1266_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		GetSamplerParameter_1266_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		glErrorInvalidEnum_1267_param := GetSamplerParameter_1266_pname // GLenum
		return
		_ = glErrorInvalidEnum_1267_param
	}
	_, _, _, _, _ = minRequiredVersion_1265_major, minRequiredVersion_1265_minor, GetSamplerParameter_1266_sampler, GetSamplerParameter_1266_pname, GetSamplerParameter_1266_params
	return nil
}
func (ϟa *GlGetTexLevelParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1268_major := uint32(3) // u32
	minRequiredVersion_1268_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1269_major := uint32(3) // u32
		minRequiredVersion_1269_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1269_major, minRequiredVersion_1269_minor
	default:
		glErrorInvalidEnum_1270_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1270_param
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	case GLenum_GL_TEXTURE_BUFFER_DATA_STORE_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET, GLenum_GL_TEXTURE_BUFFER_SIZE:
		minRequiredVersion_1271_major := uint32(3) // u32
		minRequiredVersion_1271_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1271_major, minRequiredVersion_1271_minor
	default:
		glErrorInvalidEnum_1272_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1272_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_1268_major, minRequiredVersion_1268_minor
	return nil
}
func (ϟa *GlGetTexLevelParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1273_major := uint32(3) // u32
	minRequiredVersion_1273_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1274_major := uint32(3) // u32
		minRequiredVersion_1274_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1274_major, minRequiredVersion_1274_minor
	default:
		glErrorInvalidEnum_1275_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1275_param
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	case GLenum_GL_TEXTURE_BUFFER_DATA_STORE_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET, GLenum_GL_TEXTURE_BUFFER_SIZE:
		minRequiredVersion_1276_major := uint32(3) // u32
		minRequiredVersion_1276_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1276_major, minRequiredVersion_1276_minor
	default:
		glErrorInvalidEnum_1277_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1277_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_1273_major, minRequiredVersion_1273_minor
	return nil
}
func (ϟa *GlGetTexParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1278_major := uint32(3) // u32
	minRequiredVersion_1278_minor := uint32(2) // u32
	GetTexParameter_1279_target := ϟa.Target   // GLenum
	GetTexParameter_1279_parameter := ϟa.Pname // GLenum
	GetTexParameter_1279_params := ϟa.Params   // GLintᵖ
	minRequiredVersion_1280_major := uint32(2) // u32
	minRequiredVersion_1280_minor := uint32(0) // u32
	switch GetTexParameter_1279_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1281_major := uint32(3) // u32
		minRequiredVersion_1281_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1281_major, minRequiredVersion_1281_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1282_major := uint32(3) // u32
		minRequiredVersion_1282_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1282_major, minRequiredVersion_1282_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1283_major := uint32(3) // u32
		minRequiredVersion_1283_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1283_major, minRequiredVersion_1283_minor
	default:
		glErrorInvalidEnum_1284_param := GetTexParameter_1279_target // GLenum
		return
		_ = glErrorInvalidEnum_1284_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetTexParameter_1279_parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetTexParameter_1279_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1285_major := uint32(3) // u32
		minRequiredVersion_1285_minor := uint32(0) // u32
		GetTexParameter_1279_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1285_major, minRequiredVersion_1285_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1286_major := uint32(3) // u32
		minRequiredVersion_1286_minor := uint32(1) // u32
		GetTexParameter_1279_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1286_major, minRequiredVersion_1286_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1287_major := uint32(3) // u32
		minRequiredVersion_1287_minor := uint32(2) // u32
		GetTexParameter_1279_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1287_major, minRequiredVersion_1287_minor
	default:
		glErrorInvalidEnum_1288_param := GetTexParameter_1279_parameter // GLenum
		return
		_ = glErrorInvalidEnum_1288_param
	}
	_, _, _, _, _, _, _ = minRequiredVersion_1278_major, minRequiredVersion_1278_minor, GetTexParameter_1279_target, GetTexParameter_1279_parameter, GetTexParameter_1279_params, minRequiredVersion_1280_major, minRequiredVersion_1280_minor
	return nil
}
func (ϟa *GlGetTexParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1289_major := uint32(3) // u32
	minRequiredVersion_1289_minor := uint32(2) // u32
	GetTexParameter_1290_target := ϟa.Target   // GLenum
	GetTexParameter_1290_parameter := ϟa.Pname // GLenum
	GetTexParameter_1290_params := ϟa.Params   // GLuintᵖ
	minRequiredVersion_1291_major := uint32(2) // u32
	minRequiredVersion_1291_minor := uint32(0) // u32
	switch GetTexParameter_1290_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1292_major := uint32(3) // u32
		minRequiredVersion_1292_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1292_major, minRequiredVersion_1292_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1293_major := uint32(3) // u32
		minRequiredVersion_1293_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1293_major, minRequiredVersion_1293_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1294_major := uint32(3) // u32
		minRequiredVersion_1294_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1294_major, minRequiredVersion_1294_minor
	default:
		glErrorInvalidEnum_1295_param := GetTexParameter_1290_target // GLenum
		return
		_ = glErrorInvalidEnum_1295_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetTexParameter_1290_parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetTexParameter_1290_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1296_major := uint32(3) // u32
		minRequiredVersion_1296_minor := uint32(0) // u32
		GetTexParameter_1290_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1296_major, minRequiredVersion_1296_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1297_major := uint32(3) // u32
		minRequiredVersion_1297_minor := uint32(1) // u32
		GetTexParameter_1290_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1297_major, minRequiredVersion_1297_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1298_major := uint32(3) // u32
		minRequiredVersion_1298_minor := uint32(2) // u32
		GetTexParameter_1290_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1298_major, minRequiredVersion_1298_minor
	default:
		glErrorInvalidEnum_1299_param := GetTexParameter_1290_parameter // GLenum
		return
		_ = glErrorInvalidEnum_1299_param
	}
	_, _, _, _, _, _, _ = minRequiredVersion_1289_major, minRequiredVersion_1289_minor, GetTexParameter_1290_target, GetTexParameter_1290_parameter, GetTexParameter_1290_params, minRequiredVersion_1291_major, minRequiredVersion_1291_minor
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1300_major := uint32(2)     // u32
	minRequiredVersion_1300_minor := uint32(0)     // u32
	GetTexParameter_1301_target := ϟa.Target       // GLenum
	GetTexParameter_1301_parameter := ϟa.Parameter // GLenum
	GetTexParameter_1301_params := ϟa.Values       // GLfloatᵖ
	minRequiredVersion_1302_major := uint32(2)     // u32
	minRequiredVersion_1302_minor := uint32(0)     // u32
	switch GetTexParameter_1301_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1303_major := uint32(3) // u32
		minRequiredVersion_1303_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1303_major, minRequiredVersion_1303_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1304_major := uint32(3) // u32
		minRequiredVersion_1304_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1304_major, minRequiredVersion_1304_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1305_major := uint32(3) // u32
		minRequiredVersion_1305_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1305_major, minRequiredVersion_1305_minor
	default:
		glErrorInvalidEnum_1306_param := GetTexParameter_1301_target // GLenum
		return
		_ = glErrorInvalidEnum_1306_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetTexParameter_1301_parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetTexParameter_1301_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1307_major := uint32(3) // u32
		minRequiredVersion_1307_minor := uint32(0) // u32
		GetTexParameter_1301_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1307_major, minRequiredVersion_1307_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1308_major := uint32(3) // u32
		minRequiredVersion_1308_minor := uint32(1) // u32
		GetTexParameter_1301_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1308_major, minRequiredVersion_1308_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1309_major := uint32(3) // u32
		minRequiredVersion_1309_minor := uint32(2) // u32
		GetTexParameter_1301_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1309_major, minRequiredVersion_1309_minor
	default:
		glErrorInvalidEnum_1310_param := GetTexParameter_1301_parameter // GLenum
		return
		_ = glErrorInvalidEnum_1310_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1312_msg := "No context bound" // string
		return
		_ = error_1312_msg
	}
	GetContext_1311_result := context                 // Contextʳ
	ctx := GetContext_1311_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	id := tu.Bindings.Get(ϟa.Target)                  // TextureId
	t := ctx.Instances.Textures.Get(id)               // Textureʳ
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
	_, _, _, _, _, _, _, _, _, _, _, _, _ = minRequiredVersion_1300_major, minRequiredVersion_1300_minor, GetTexParameter_1301_target, GetTexParameter_1301_parameter, GetTexParameter_1301_params, minRequiredVersion_1302_major, minRequiredVersion_1302_minor, context, GetContext_1311_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1313_major := uint32(2)     // u32
	minRequiredVersion_1313_minor := uint32(0)     // u32
	GetTexParameter_1314_target := ϟa.Target       // GLenum
	GetTexParameter_1314_parameter := ϟa.Parameter // GLenum
	GetTexParameter_1314_params := ϟa.Values       // GLintᵖ
	minRequiredVersion_1315_major := uint32(2)     // u32
	minRequiredVersion_1315_minor := uint32(0)     // u32
	switch GetTexParameter_1314_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1316_major := uint32(3) // u32
		minRequiredVersion_1316_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1316_major, minRequiredVersion_1316_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1317_major := uint32(3) // u32
		minRequiredVersion_1317_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1317_major, minRequiredVersion_1317_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1318_major := uint32(3) // u32
		minRequiredVersion_1318_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1318_major, minRequiredVersion_1318_minor
	default:
		glErrorInvalidEnum_1319_param := GetTexParameter_1314_target // GLenum
		return
		_ = glErrorInvalidEnum_1319_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch GetTexParameter_1314_parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		GetTexParameter_1314_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1320_major := uint32(3) // u32
		minRequiredVersion_1320_minor := uint32(0) // u32
		GetTexParameter_1314_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1320_major, minRequiredVersion_1320_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1321_major := uint32(3) // u32
		minRequiredVersion_1321_minor := uint32(1) // u32
		GetTexParameter_1314_params.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1321_major, minRequiredVersion_1321_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1322_major := uint32(3) // u32
		minRequiredVersion_1322_minor := uint32(2) // u32
		GetTexParameter_1314_params.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1322_major, minRequiredVersion_1322_minor
	default:
		glErrorInvalidEnum_1323_param := GetTexParameter_1314_parameter // GLenum
		return
		_ = glErrorInvalidEnum_1323_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1325_msg := "No context bound" // string
		return
		_ = error_1325_msg
	}
	GetContext_1324_result := context                 // Contextʳ
	ctx := GetContext_1324_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	id := tu.Bindings.Get(ϟa.Target)                  // TextureId
	t := ctx.Instances.Textures.Get(id)               // Textureʳ
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
	_, _, _, _, _, _, _, _, _, _, _, _, _ = minRequiredVersion_1313_major, minRequiredVersion_1313_minor, GetTexParameter_1314_target, GetTexParameter_1314_parameter, GetTexParameter_1314_params, minRequiredVersion_1315_major, minRequiredVersion_1315_minor, context, GetContext_1324_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlIsSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1326_major := uint32(3) // u32
	minRequiredVersion_1326_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1326_major, minRequiredVersion_1326_minor
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1327_major := uint32(2)   // u32
	minRequiredVersion_1327_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1329_msg := "No context bound" // string
		return
		_ = error_1329_msg
	}
	GetContext_1328_result := context // Contextʳ
	ctx := GetContext_1328_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Textures.Contains(ϟa.Texture) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_1327_major, minRequiredVersion_1327_minor, context, GetContext_1328_result, ctx
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1330_major := uint32(2) // u32
	minRequiredVersion_1330_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_PACK_ALIGNMENT, GLenum_GL_UNPACK_ALIGNMENT:
	case GLenum_GL_PACK_IMAGE_HEIGHT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_IMAGES, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS:
		minRequiredVersion_1331_major := uint32(3) // u32
		minRequiredVersion_1331_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1331_major, minRequiredVersion_1331_minor
	default:
		glErrorInvalidEnum_1332_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1332_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1334_msg := "No context bound" // string
		return
		_ = error_1334_msg
	}
	GetContext_1333_result := context // Contextʳ
	ctx := GetContext_1333_result     // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1330_major, minRequiredVersion_1330_minor, context, GetContext_1333_result, ctx
	return nil
}
func (ϟa *GlSamplerParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1335_major := uint32(3)   // u32
	minRequiredVersion_1335_minor := uint32(2)   // u32
	SamplerParameterv_1336_sampler := ϟa.Sampler // SamplerId
	SamplerParameterv_1336_pname := ϟa.Pname     // GLenum
	SamplerParameterv_1336_params := ϟa.Param    // GLintᶜᵖ
	minRequiredVersion_1337_major := uint32(3)   // u32
	minRequiredVersion_1337_minor := uint32(0)   // u32
	switch SamplerParameterv_1336_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		SamplerParameterv_1336_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1338_major := uint32(3) // u32
		minRequiredVersion_1338_minor := uint32(2) // u32
		SamplerParameterv_1336_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1338_major, minRequiredVersion_1338_minor
	default:
		glErrorInvalidEnum_1339_param := SamplerParameterv_1336_pname // GLenum
		return
		_ = glErrorInvalidEnum_1339_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_1335_major, minRequiredVersion_1335_minor, SamplerParameterv_1336_sampler, SamplerParameterv_1336_pname, SamplerParameterv_1336_params, minRequiredVersion_1337_major, minRequiredVersion_1337_minor
	return nil
}
func (ϟa *GlSamplerParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1340_major := uint32(3)   // u32
	minRequiredVersion_1340_minor := uint32(2)   // u32
	SamplerParameterv_1341_sampler := ϟa.Sampler // SamplerId
	SamplerParameterv_1341_pname := ϟa.Pname     // GLenum
	SamplerParameterv_1341_params := ϟa.Param    // GLuintᶜᵖ
	minRequiredVersion_1342_major := uint32(3)   // u32
	minRequiredVersion_1342_minor := uint32(0)   // u32
	switch SamplerParameterv_1341_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		SamplerParameterv_1341_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1343_major := uint32(3) // u32
		minRequiredVersion_1343_minor := uint32(2) // u32
		SamplerParameterv_1341_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1343_major, minRequiredVersion_1343_minor
	default:
		glErrorInvalidEnum_1344_param := SamplerParameterv_1341_pname // GLenum
		return
		_ = glErrorInvalidEnum_1344_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_1340_major, minRequiredVersion_1340_minor, SamplerParameterv_1341_sampler, SamplerParameterv_1341_pname, SamplerParameterv_1341_params, minRequiredVersion_1342_major, minRequiredVersion_1342_minor
	return nil
}
func (ϟa *GlSamplerParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1345_major := uint32(3) // u32
	minRequiredVersion_1345_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1346_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1346_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1345_major, minRequiredVersion_1345_minor
	return nil
}
func (ϟa *GlSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1347_major := uint32(3)   // u32
	minRequiredVersion_1347_minor := uint32(0)   // u32
	SamplerParameterv_1348_sampler := ϟa.Sampler // SamplerId
	SamplerParameterv_1348_pname := ϟa.Pname     // GLenum
	SamplerParameterv_1348_params := ϟa.Param    // GLfloatᶜᵖ
	minRequiredVersion_1349_major := uint32(3)   // u32
	minRequiredVersion_1349_minor := uint32(0)   // u32
	switch SamplerParameterv_1348_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		SamplerParameterv_1348_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1350_major := uint32(3) // u32
		minRequiredVersion_1350_minor := uint32(2) // u32
		SamplerParameterv_1348_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1350_major, minRequiredVersion_1350_minor
	default:
		glErrorInvalidEnum_1351_param := SamplerParameterv_1348_pname // GLenum
		return
		_ = glErrorInvalidEnum_1351_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_1347_major, minRequiredVersion_1347_minor, SamplerParameterv_1348_sampler, SamplerParameterv_1348_pname, SamplerParameterv_1348_params, minRequiredVersion_1349_major, minRequiredVersion_1349_minor
	return nil
}
func (ϟa *GlSamplerParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1352_major := uint32(3) // u32
	minRequiredVersion_1352_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		glErrorInvalidEnum_1353_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1353_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1352_major, minRequiredVersion_1352_minor
	return nil
}
func (ϟa *GlSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1354_major := uint32(3)   // u32
	minRequiredVersion_1354_minor := uint32(0)   // u32
	SamplerParameterv_1355_sampler := ϟa.Sampler // SamplerId
	SamplerParameterv_1355_pname := ϟa.Pname     // GLenum
	SamplerParameterv_1355_params := ϟa.Param    // GLintᶜᵖ
	minRequiredVersion_1356_major := uint32(3)   // u32
	minRequiredVersion_1356_minor := uint32(0)   // u32
	switch SamplerParameterv_1355_pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		SamplerParameterv_1355_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1357_major := uint32(3) // u32
		minRequiredVersion_1357_minor := uint32(2) // u32
		SamplerParameterv_1355_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1357_major, minRequiredVersion_1357_minor
	default:
		glErrorInvalidEnum_1358_param := SamplerParameterv_1355_pname // GLenum
		return
		_ = glErrorInvalidEnum_1358_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_1354_major, minRequiredVersion_1354_minor, SamplerParameterv_1355_sampler, SamplerParameterv_1355_pname, SamplerParameterv_1355_params, minRequiredVersion_1356_major, minRequiredVersion_1356_minor
	return nil
}
func (ϟa *GlTexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1359_major := uint32(3) // u32
	minRequiredVersion_1359_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_BUFFER:
	default:
		glErrorInvalidEnum_1360_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1360_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_R16, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGBA16, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI:
	default:
		glErrorInvalidEnum_1361_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1361_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1359_major, minRequiredVersion_1359_minor
	return nil
}
func (ϟa *GlTexBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1362_major := uint32(3) // u32
	minRequiredVersion_1362_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_BUFFER:
	default:
		glErrorInvalidEnum_1363_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1363_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_R16, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGBA16, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI:
	default:
		glErrorInvalidEnum_1364_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1364_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1362_major, minRequiredVersion_1362_minor
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1365_major := uint32(2) // u32
	minRequiredVersion_1365_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1366_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1366_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1367_major := uint32(3) // u32
		minRequiredVersion_1367_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1367_major, minRequiredVersion_1367_minor
	case GLenum_GL_STENCIL_INDEX:
		minRequiredVersion_1368_major := uint32(3) // u32
		minRequiredVersion_1368_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1368_major, minRequiredVersion_1368_minor
	default:
		glErrorInvalidEnum_1369_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1369_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1370_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1370_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1371_major := uint32(3) // u32
		minRequiredVersion_1371_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1371_major, minRequiredVersion_1371_minor
	default:
		glErrorInvalidEnum_1372_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1372_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1374_msg := "No context bound" // string
		return
		_ = error_1374_msg
	}
	GetContext_1373_result := context                 // Contextʳ
	ctx := GetContext_1373_result                     // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_1365_major, minRequiredVersion_1365_minor, context, GetContext_1373_result, ctx, tu
	return nil
}
func (ϟa *GlTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1375_major := uint32(3) // u32
	minRequiredVersion_1375_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1376_major := uint32(3) // u32
		minRequiredVersion_1376_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1376_major, minRequiredVersion_1376_minor
	default:
		glErrorInvalidEnum_1377_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1377_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	case GLenum_GL_STENCIL_INDEX:
		minRequiredVersion_1378_major := uint32(3) // u32
		minRequiredVersion_1378_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1378_major, minRequiredVersion_1378_minor
	default:
		glErrorInvalidEnum_1379_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1379_param
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1380_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1380_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		glErrorInvalidEnum_1381_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1381_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1383_msg := "No context bound" // string
		return
		_ = error_1383_msg
	}
	GetContext_1382_result := context // Contextʳ
	ctx := GetContext_1382_result     // Contextʳ
	if ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) && ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) {
		size := (externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)) * (uint32(ϟa.Depth)) // u32
		U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(size), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = size
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1375_major, minRequiredVersion_1375_minor, context, GetContext_1382_result, ctx
	return nil
}
func (ϟa *GlTexParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1384_major := uint32(3) // u32
	minRequiredVersion_1384_minor := uint32(2) // u32
	TexParameterv_1385_target := ϟa.Target     // GLenum
	TexParameterv_1385_pname := ϟa.Pname       // GLenum
	TexParameterv_1385_params := ϟa.Params     // GLintᶜᵖ
	switch TexParameterv_1385_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1386_major := uint32(3) // u32
		minRequiredVersion_1386_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1386_major, minRequiredVersion_1386_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1387_major := uint32(3) // u32
		minRequiredVersion_1387_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1387_major, minRequiredVersion_1387_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1388_major := uint32(3) // u32
		minRequiredVersion_1388_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1388_major, minRequiredVersion_1388_minor
	default:
		glErrorInvalidEnum_1389_param := TexParameterv_1385_target // GLenum
		return
		_ = glErrorInvalidEnum_1389_param
	}
	switch TexParameterv_1385_pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		TexParameterv_1385_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1390_major := uint32(3) // u32
		minRequiredVersion_1390_minor := uint32(0) // u32
		TexParameterv_1385_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1390_major, minRequiredVersion_1390_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1391_major := uint32(3) // u32
		minRequiredVersion_1391_minor := uint32(1) // u32
		TexParameterv_1385_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1391_major, minRequiredVersion_1391_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1392_major := uint32(3) // u32
		minRequiredVersion_1392_minor := uint32(2) // u32
		TexParameterv_1385_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1392_major, minRequiredVersion_1392_minor
	default:
		glErrorInvalidEnum_1393_param := TexParameterv_1385_pname // GLenum
		return
		_ = glErrorInvalidEnum_1393_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1384_major, minRequiredVersion_1384_minor, TexParameterv_1385_target, TexParameterv_1385_pname, TexParameterv_1385_params
	return nil
}
func (ϟa *GlTexParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1394_major := uint32(3) // u32
	minRequiredVersion_1394_minor := uint32(2) // u32
	TexParameterv_1395_target := ϟa.Target     // GLenum
	TexParameterv_1395_pname := ϟa.Pname       // GLenum
	TexParameterv_1395_params := ϟa.Params     // GLuintᶜᵖ
	switch TexParameterv_1395_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1396_major := uint32(3) // u32
		minRequiredVersion_1396_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1396_major, minRequiredVersion_1396_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1397_major := uint32(3) // u32
		minRequiredVersion_1397_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1397_major, minRequiredVersion_1397_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1398_major := uint32(3) // u32
		minRequiredVersion_1398_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1398_major, minRequiredVersion_1398_minor
	default:
		glErrorInvalidEnum_1399_param := TexParameterv_1395_target // GLenum
		return
		_ = glErrorInvalidEnum_1399_param
	}
	switch TexParameterv_1395_pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		TexParameterv_1395_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1400_major := uint32(3) // u32
		minRequiredVersion_1400_minor := uint32(0) // u32
		TexParameterv_1395_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1400_major, minRequiredVersion_1400_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1401_major := uint32(3) // u32
		minRequiredVersion_1401_minor := uint32(1) // u32
		TexParameterv_1395_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1401_major, minRequiredVersion_1401_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1402_major := uint32(3) // u32
		minRequiredVersion_1402_minor := uint32(2) // u32
		TexParameterv_1395_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1402_major, minRequiredVersion_1402_minor
	default:
		glErrorInvalidEnum_1403_param := TexParameterv_1395_pname // GLenum
		return
		_ = glErrorInvalidEnum_1403_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1394_major, minRequiredVersion_1394_minor, TexParameterv_1395_target, TexParameterv_1395_pname, TexParameterv_1395_params
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1404_major := uint32(2) // u32
	minRequiredVersion_1404_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1405_major := uint32(3) // u32
		minRequiredVersion_1405_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1405_major, minRequiredVersion_1405_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1406_major := uint32(3) // u32
		minRequiredVersion_1406_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1406_major, minRequiredVersion_1406_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1407_major := uint32(3) // u32
		minRequiredVersion_1407_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1407_major, minRequiredVersion_1407_minor
	default:
		glErrorInvalidEnum_1408_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1408_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1409_major := uint32(3) // u32
		minRequiredVersion_1409_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1409_major, minRequiredVersion_1409_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1410_major := uint32(3) // u32
		minRequiredVersion_1410_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1410_major, minRequiredVersion_1410_minor
	default:
		glErrorInvalidEnum_1411_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1411_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1413_msg := "No context bound" // string
		return
		_ = error_1413_msg
	}
	GetContext_1412_result := context                 // Contextʳ
	ctx := GetContext_1412_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1404_major, minRequiredVersion_1404_minor, context, GetContext_1412_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1414_major := uint32(2) // u32
	minRequiredVersion_1414_minor := uint32(0) // u32
	TexParameterv_1415_target := ϟa.Target     // GLenum
	TexParameterv_1415_pname := ϟa.Pname       // GLenum
	TexParameterv_1415_params := ϟa.Params     // GLfloatᶜᵖ
	switch TexParameterv_1415_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1416_major := uint32(3) // u32
		minRequiredVersion_1416_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1416_major, minRequiredVersion_1416_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1417_major := uint32(3) // u32
		minRequiredVersion_1417_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1417_major, minRequiredVersion_1417_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1418_major := uint32(3) // u32
		minRequiredVersion_1418_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1418_major, minRequiredVersion_1418_minor
	default:
		glErrorInvalidEnum_1419_param := TexParameterv_1415_target // GLenum
		return
		_ = glErrorInvalidEnum_1419_param
	}
	switch TexParameterv_1415_pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		TexParameterv_1415_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1420_major := uint32(3) // u32
		minRequiredVersion_1420_minor := uint32(0) // u32
		TexParameterv_1415_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1420_major, minRequiredVersion_1420_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1421_major := uint32(3) // u32
		minRequiredVersion_1421_minor := uint32(1) // u32
		TexParameterv_1415_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1421_major, minRequiredVersion_1421_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1422_major := uint32(3) // u32
		minRequiredVersion_1422_minor := uint32(2) // u32
		TexParameterv_1415_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1422_major, minRequiredVersion_1422_minor
	default:
		glErrorInvalidEnum_1423_param := TexParameterv_1415_pname // GLenum
		return
		_ = glErrorInvalidEnum_1423_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1414_major, minRequiredVersion_1414_minor, TexParameterv_1415_target, TexParameterv_1415_pname, TexParameterv_1415_params
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1424_major := uint32(2) // u32
	minRequiredVersion_1424_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1425_major := uint32(3) // u32
		minRequiredVersion_1425_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1425_major, minRequiredVersion_1425_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1426_major := uint32(3) // u32
		minRequiredVersion_1426_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1426_major, minRequiredVersion_1426_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1427_major := uint32(3) // u32
		minRequiredVersion_1427_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1427_major, minRequiredVersion_1427_minor
	default:
		glErrorInvalidEnum_1428_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1428_param
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1429_major := uint32(3) // u32
		minRequiredVersion_1429_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1429_major, minRequiredVersion_1429_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1430_major := uint32(3) // u32
		minRequiredVersion_1430_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1430_major, minRequiredVersion_1430_minor
	default:
		glErrorInvalidEnum_1431_param := ϟa.Parameter // GLenum
		return
		_ = glErrorInvalidEnum_1431_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1433_msg := "No context bound" // string
		return
		_ = error_1433_msg
	}
	GetContext_1432_result := context                 // Contextʳ
	ctx := GetContext_1432_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_1424_major, minRequiredVersion_1424_minor, context, GetContext_1432_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1434_major := uint32(2) // u32
	minRequiredVersion_1434_minor := uint32(0) // u32
	TexParameterv_1435_target := ϟa.Target     // GLenum
	TexParameterv_1435_pname := ϟa.Pname       // GLenum
	TexParameterv_1435_params := ϟa.Params     // GLintᶜᵖ
	switch TexParameterv_1435_target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1436_major := uint32(3) // u32
		minRequiredVersion_1436_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1436_major, minRequiredVersion_1436_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1437_major := uint32(3) // u32
		minRequiredVersion_1437_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1437_major, minRequiredVersion_1437_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1438_major := uint32(3) // u32
		minRequiredVersion_1438_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1438_major, minRequiredVersion_1438_minor
	default:
		glErrorInvalidEnum_1439_param := TexParameterv_1435_target // GLenum
		return
		_ = glErrorInvalidEnum_1439_param
	}
	switch TexParameterv_1435_pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
		TexParameterv_1435_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1440_major := uint32(3) // u32
		minRequiredVersion_1440_minor := uint32(0) // u32
		TexParameterv_1435_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1440_major, minRequiredVersion_1440_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1441_major := uint32(3) // u32
		minRequiredVersion_1441_minor := uint32(1) // u32
		TexParameterv_1435_params.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1441_major, minRequiredVersion_1441_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1442_major := uint32(3) // u32
		minRequiredVersion_1442_minor := uint32(2) // u32
		TexParameterv_1435_params.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = minRequiredVersion_1442_major, minRequiredVersion_1442_minor
	default:
		glErrorInvalidEnum_1443_param := TexParameterv_1435_pname // GLenum
		return
		_ = glErrorInvalidEnum_1443_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1434_major, minRequiredVersion_1434_minor, TexParameterv_1435_target, TexParameterv_1435_pname, TexParameterv_1435_params
	return nil
}
func (ϟa *GlTexStorage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1444_major := uint32(3) // u32
	minRequiredVersion_1444_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	default:
		glErrorInvalidEnum_1445_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1445_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8, GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1446_major := uint32(3) // u32
		minRequiredVersion_1446_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1446_major, minRequiredVersion_1446_minor
	default:
		glErrorInvalidEnum_1447_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1447_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1444_major, minRequiredVersion_1444_minor
	return nil
}
func (ϟa *GlTexStorage2DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1448_major := uint32(3) // u32
	minRequiredVersion_1448_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
	default:
		glErrorInvalidEnum_1449_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1449_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1450_major := uint32(3) // u32
		minRequiredVersion_1450_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1450_major, minRequiredVersion_1450_minor
	default:
		glErrorInvalidEnum_1451_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1451_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1448_major, minRequiredVersion_1448_minor
	return nil
}
func (ϟa *GlTexStorage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1452_major := uint32(3) // u32
	minRequiredVersion_1452_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1453_major := uint32(3) // u32
		minRequiredVersion_1453_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1453_major, minRequiredVersion_1453_minor
	default:
		glErrorInvalidEnum_1454_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1454_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8, GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1455_major := uint32(3) // u32
		minRequiredVersion_1455_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1455_major, minRequiredVersion_1455_minor
	default:
		glErrorInvalidEnum_1456_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1456_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1452_major, minRequiredVersion_1452_minor
	return nil
}
func (ϟa *GlTexStorage3DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1457_major := uint32(3) // u32
	minRequiredVersion_1457_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY:
	default:
		glErrorInvalidEnum_1458_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1458_param
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		glErrorInvalidEnum_1459_param := ϟa.Internalformat // GLenum
		return
		_ = glErrorInvalidEnum_1459_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1457_major, minRequiredVersion_1457_minor
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1460_major := uint32(2) // u32
	minRequiredVersion_1460_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		glErrorInvalidEnum_1461_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1461_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1462_major := uint32(3) // u32
		minRequiredVersion_1462_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1462_major, minRequiredVersion_1462_minor
	default:
		glErrorInvalidEnum_1463_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1463_param
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1464_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1464_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1465_major := uint32(3) // u32
		minRequiredVersion_1465_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1465_major, minRequiredVersion_1465_minor
	default:
		glErrorInvalidEnum_1466_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1466_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1468_msg := "No context bound" // string
		return
		_ = error_1468_msg
	}
	GetContext_1467_result := context                 // Contextʳ
	ctx := GetContext_1467_result                     // Contextʳ
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
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = minRequiredVersion_1460_major, minRequiredVersion_1460_minor, context, GetContext_1467_result, ctx, tu, image, pbo, url, src_width, src_stride, src_size, dst_stride, dst_offset, src_data, line_bytes
	return nil
}
func (ϟa *GlTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1469_major := uint32(3) // u32
	minRequiredVersion_1469_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1470_major := uint32(3) // u32
		minRequiredVersion_1470_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1470_major, minRequiredVersion_1470_minor
	default:
		glErrorInvalidEnum_1471_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1471_param
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	default:
		glErrorInvalidEnum_1472_param := ϟa.Format // GLenum
		return
		_ = glErrorInvalidEnum_1472_param
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1473_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1473_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		glErrorInvalidEnum_1474_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1474_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1476_msg := "No context bound" // string
		return
		_ = error_1476_msg
	}
	GetContext_1475_result := context // Contextʳ
	ctx := GetContext_1475_result     // Contextʳ
	if ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) && ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) {
		size := (externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)) * (uint32(ϟa.Depth)) // u32
		U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(size), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = size
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1469_major, minRequiredVersion_1469_minor, context, GetContext_1475_result, ctx
	return nil
}
func (ϟa *GlBeginTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1477_major := uint32(3) // u32
	minRequiredVersion_1477_minor := uint32(0) // u32
	switch ϟa.PrimitiveMode {
	case GLenum_GL_LINES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES:
	default:
		glErrorInvalidEnum_1478_param := ϟa.PrimitiveMode // GLenum
		return
		_ = glErrorInvalidEnum_1478_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1477_major, minRequiredVersion_1477_minor
	return nil
}
func (ϟa *GlBindTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1479_major := uint32(3) // u32
	minRequiredVersion_1479_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK:
	default:
		glErrorInvalidEnum_1480_param := ϟa.Target // GLenum
		return
		_ = glErrorInvalidEnum_1480_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1479_major, minRequiredVersion_1479_minor
	return nil
}
func (ϟa *GlDeleteTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1481_major := uint32(3) // u32
	minRequiredVersion_1481_minor := uint32(0) // u32
	ϟa.Ids.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.N), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1481_major, minRequiredVersion_1481_minor
	return nil
}
func (ϟa *GlEndTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1482_major := uint32(3) // u32
	minRequiredVersion_1482_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1482_major, minRequiredVersion_1482_minor
	return nil
}
func (ϟa *GlGenTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1483_major := uint32(3) // u32
	minRequiredVersion_1483_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Ids.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.N), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_1483_major, minRequiredVersion_1483_minor
	return nil
}
func (ϟa *GlGetTransformFeedbackVarying) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1484_major := uint32(3)         // u32
	minRequiredVersion_1484_minor := uint32(0)         // u32
	writeString_1485_buffer_size := ϟa.BufSize         // GLsizei
	writeString_1485_buffer_bytes_written := ϟa.Length // GLsizeiᵖ
	writeString_1485_buffer := ϟa.Name                 // GLcharᵖ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if ((writeString_1485_buffer) != (GLcharᵖ{})) && ((writeString_1485_buffer_size) > (GLsizei(int32(0)))) {
		buffer_size2 := writeString_1485_buffer_size // GLsizei
		if (writeString_1485_buffer_bytes_written) != (GLsizeiᵖ{}) {
			length := GLsizei(ϟa.Length.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
			writeString_1485_buffer_bytes_written.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(length, ϟa, ϟs, ϟd, ϟl, ϟb)
			writeString_1485_buffer.Slice(uint64(GLsizei(int32(0))), uint64((length)+(GLsizei(int32(1)))), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
			_ = length
		} else {
			writeString_1485_buffer.Slice(uint64(GLsizei(int32(0))), uint64(buffer_size2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		_ = buffer_size2
	}
	ϟa.Size.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Size.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = minRequiredVersion_1484_major, minRequiredVersion_1484_minor, writeString_1485_buffer_size, writeString_1485_buffer_bytes_written, writeString_1485_buffer
	return nil
}
func (ϟa *GlIsTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1486_major := uint32(3) // u32
	minRequiredVersion_1486_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1486_major, minRequiredVersion_1486_minor
	return nil
}
func (ϟa *GlPauseTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1487_major := uint32(3) // u32
	minRequiredVersion_1487_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1487_major, minRequiredVersion_1487_minor
	return nil
}
func (ϟa *GlResumeTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1488_major := uint32(3) // u32
	minRequiredVersion_1488_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1488_major, minRequiredVersion_1488_minor
	return nil
}
func (ϟa *GlTransformFeedbackVaryings) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1489_major := uint32(3) // u32
	minRequiredVersion_1489_minor := uint32(0) // u32
	switch ϟa.BufferMode {
	case GLenum_GL_INTERLEAVED_ATTRIBS, GLenum_GL_SEPARATE_ATTRIBS:
	default:
		glErrorInvalidEnum_1490_param := ϟa.BufferMode // GLenum
		return
		_ = glErrorInvalidEnum_1490_param
	}
	names := ϟa.Varyings.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		name := strings.TrimRight(string(Charᵖ(names.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
		_ = name
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = minRequiredVersion_1489_major, minRequiredVersion_1489_minor, names
	return nil
}
func (ϟa *GlBindVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1491_major := uint32(3)   // u32
	minRequiredVersion_1491_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1493_msg := "No context bound" // string
		return
		_ = error_1493_msg
	}
	GetContext_1492_result := context // Contextʳ
	ctx := GetContext_1492_result     // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = &VertexArray{}
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1491_major, minRequiredVersion_1491_minor, context, GetContext_1492_result, ctx
	return nil
}
func (ϟa *GlBindVertexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1494_major := uint32(3) // u32
	minRequiredVersion_1494_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1494_major, minRequiredVersion_1494_minor
	return nil
}
func (ϟa *GlDeleteVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1495_major := uint32(3)   // u32
	minRequiredVersion_1495_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1497_msg := "No context bound" // string
		return
		_ = error_1497_msg
	}
	GetContext_1496_result := context                                     // Contextʳ
	ctx := GetContext_1496_result                                         // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1495_major, minRequiredVersion_1495_minor, context, GetContext_1496_result, ctx, a
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1498_major := uint32(2)   // u32
	minRequiredVersion_1498_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1500_msg := "No context bound" // string
		return
		_ = error_1500_msg
	}
	GetContext_1499_result := context // Contextʳ
	ctx := GetContext_1499_result     // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1498_major, minRequiredVersion_1498_minor, context, GetContext_1499_result, ctx
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1501_major := uint32(2)   // u32
	minRequiredVersion_1501_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1503_msg := "No context bound" // string
		return
		_ = error_1503_msg
	}
	GetContext_1502_result := context // Contextʳ
	ctx := GetContext_1502_result     // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1501_major, minRequiredVersion_1501_minor, context, GetContext_1502_result, ctx
	return nil
}
func (ϟa *GlGenVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1504_major := uint32(3)                            // u32
	minRequiredVersion_1504_minor := uint32(0)                            // u32
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1506_msg := "No context bound" // string
		return
		_ = error_1506_msg
	}
	GetContext_1505_result := context // Contextʳ
	ctx := GetContext_1505_result     // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = &VertexArray{}
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1504_major, minRequiredVersion_1504_minor, a, context, GetContext_1505_result, ctx
	return nil
}
func (ϟa *GlGetVertexAttribIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1507_major := uint32(3) // u32
	minRequiredVersion_1507_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1508_major := uint32(3) // u32
		minRequiredVersion_1508_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1508_major, minRequiredVersion_1508_minor
	default:
		glErrorInvalidEnum_1509_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1509_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1507_major, minRequiredVersion_1507_minor
	return nil
}
func (ϟa *GlGetVertexAttribIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1510_major := uint32(3) // u32
	minRequiredVersion_1510_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1511_major := uint32(3) // u32
		minRequiredVersion_1511_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1511_major, minRequiredVersion_1511_minor
	default:
		glErrorInvalidEnum_1512_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1512_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1510_major, minRequiredVersion_1510_minor
	return nil
}
func (ϟa *GlGetVertexAttribPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1513_major := uint32(2) // u32
	minRequiredVersion_1513_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_POINTER:
	default:
		glErrorInvalidEnum_1514_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1514_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1513_major, minRequiredVersion_1513_minor
	return nil
}
func (ϟa *GlGetVertexAttribfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1515_major := uint32(2) // u32
	minRequiredVersion_1515_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1516_major := uint32(3) // u32
		minRequiredVersion_1516_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1516_major, minRequiredVersion_1516_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1517_major := uint32(3) // u32
		minRequiredVersion_1517_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1517_major, minRequiredVersion_1517_minor
	default:
		glErrorInvalidEnum_1518_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1518_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1515_major, minRequiredVersion_1515_minor
	return nil
}
func (ϟa *GlGetVertexAttribiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1519_major := uint32(2) // u32
	minRequiredVersion_1519_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1520_major := uint32(3) // u32
		minRequiredVersion_1520_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1520_major, minRequiredVersion_1520_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1521_major := uint32(3) // u32
		minRequiredVersion_1521_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1521_major, minRequiredVersion_1521_minor
	default:
		glErrorInvalidEnum_1522_param := ϟa.Pname // GLenum
		return
		_ = glErrorInvalidEnum_1522_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1519_major, minRequiredVersion_1519_minor
	return nil
}
func (ϟa *GlIsVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1523_major := uint32(3) // u32
	minRequiredVersion_1523_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1523_major, minRequiredVersion_1523_minor
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1524_major := uint32(2) // u32
	minRequiredVersion_1524_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1524_major, minRequiredVersion_1524_minor
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1525_major := uint32(2) // u32
	minRequiredVersion_1525_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1525_major, minRequiredVersion_1525_minor
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1526_major := uint32(2) // u32
	minRequiredVersion_1526_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1526_major, minRequiredVersion_1526_minor
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1527_major := uint32(2) // u32
	minRequiredVersion_1527_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(2), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1527_major, minRequiredVersion_1527_minor
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1528_major := uint32(2) // u32
	minRequiredVersion_1528_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1528_major, minRequiredVersion_1528_minor
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1529_major := uint32(2) // u32
	minRequiredVersion_1529_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(3), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1529_major, minRequiredVersion_1529_minor
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1530_major := uint32(2) // u32
	minRequiredVersion_1530_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1530_major, minRequiredVersion_1530_minor
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1531_major := uint32(2) // u32
	minRequiredVersion_1531_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1531_major, minRequiredVersion_1531_minor
	return nil
}
func (ϟa *GlVertexAttribBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1532_major := uint32(3) // u32
	minRequiredVersion_1532_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1532_major, minRequiredVersion_1532_minor
	return nil
}
func (ϟa *GlVertexAttribDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1533_major := uint32(3) // u32
	minRequiredVersion_1533_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1533_major, minRequiredVersion_1533_minor
	return nil
}
func (ϟa *GlVertexAttribFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1534_major := uint32(3) // u32
	minRequiredVersion_1534_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_1535_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1535_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1534_major, minRequiredVersion_1534_minor
	return nil
}
func (ϟa *GlVertexAttribI4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1536_major := uint32(3) // u32
	minRequiredVersion_1536_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1536_major, minRequiredVersion_1536_minor
	return nil
}
func (ϟa *GlVertexAttribI4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1537_major := uint32(3) // u32
	minRequiredVersion_1537_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1537_major, minRequiredVersion_1537_minor
	return nil
}
func (ϟa *GlVertexAttribI4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1538_major := uint32(3) // u32
	minRequiredVersion_1538_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1538_major, minRequiredVersion_1538_minor
	return nil
}
func (ϟa *GlVertexAttribI4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1539_major := uint32(3) // u32
	minRequiredVersion_1539_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1539_major, minRequiredVersion_1539_minor
	return nil
}
func (ϟa *GlVertexAttribIFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1540_major := uint32(3) // u32
	minRequiredVersion_1540_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_1541_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1541_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1540_major, minRequiredVersion_1540_minor
	return nil
}
func (ϟa *GlVertexAttribIPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1542_major := uint32(3) // u32
	minRequiredVersion_1542_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1543_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1543_ext
	case GLenum_GL_BYTE, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		glErrorInvalidEnum_1544_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1544_param
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1542_major, minRequiredVersion_1542_minor
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1545_major := uint32(2) // u32
	minRequiredVersion_1545_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1546_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1546_ext
	case GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_1547_major := uint32(3) // u32
		minRequiredVersion_1547_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1547_major, minRequiredVersion_1547_minor
	default:
		glErrorInvalidEnum_1548_param := ϟa.Type // GLenum
		return
		_ = glErrorInvalidEnum_1548_param
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	if (context) == ((*Context)(nil)) {
		error_1550_msg := "No context bound" // string
		return
		_ = error_1550_msg
	}
	GetContext_1549_result := context               // Contextʳ
	ctx := GetContext_1549_result                   // Contextʳ
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayʳ
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = ϟa.Data
	a.Buffer = ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1545_major, minRequiredVersion_1545_minor, context, GetContext_1549_result, ctx, a
	return nil
}
func (ϟa *GlVertexBindingDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1551_major := uint32(3) // u32
	minRequiredVersion_1551_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1551_major, minRequiredVersion_1551_minor
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
	CreateContext_1552_result := ctx // Contextʳ
	ϟc.EGLContexts[context] = CreateContext_1552_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1552_result
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1553_context := ϟc.EGLContexts.Get(ϟa.Context) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1553_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1553_context
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
	CreateContext_1554_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1554_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1554_result
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
	CreateContext_1555_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1555_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1555_result
	return nil
}
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1556_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1556_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1556_context
	return nil
}
func (ϟa *GlXMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1557_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1557_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1557_context
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
	CreateContext_1558_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1558_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1558_result
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
	CreateContext_1559_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1559_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1559_result
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1560_context := ϟc.WGLContexts.Get(ϟa.Hglrc) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1560_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1560_context
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
	CreateContext_1561_result := ctx // Contextʳ
	ϟc.CGLContexts[context] = CreateContext_1561_result
	ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(context, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1561_result
	return nil
}
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (ϟe error) {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1562_context := ϟc.CGLContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1562_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1562_context
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
		error_1564_msg := "No context bound" // string
		return
		_ = error_1564_msg
	}
	GetContext_1563_result := context // Contextʳ
	ctx := GetContext_1563_result     // Contextʳ
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
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_1563_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
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
