////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
	"android.googlesource.com/platform/tools/gpu/service"
)

func init() {
	s := schemaBuilder{
		arrays:       make(map[int]*service.ArrayInfo),
		staticArrays: make(map[int]*service.StaticArrayInfo),
		maps:         make(map[int]*service.MapInfo),
		enums:        make(map[int]*service.EnumInfo),
		structs:      make(map[int]*service.StructInfo),
		classes:      make(map[int]*service.ClassInfo),
	}
	schema.RegisterAtom(service.AtomInfo{
		Type: 0,
		Name: "init",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "color_fmt",
				Type: s.getEnumInfo(13),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "depth_fmt",
				Type: s.getEnumInfo(13),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_fmt",
				Type: s.getEnumInfo(13),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 1,
		Name: "startTimer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "index",
				Type: schema.U8,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 2,
		Name: "stopTimer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "index",
				Type: schema.U8,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.U64,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type:             3,
		Name:             "flushPostBuffer",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 4,
		Name: "eglCreateContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "version",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "context",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 5,
		Name: "eglMakeCurrent",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "context",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type:             6,
		Name:             "eglSwapBuffers",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     true,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 7,
		Name: "wglSwapBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "hdc",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     true,
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/dd369060%28v=vs.85%29]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 8,
		Name: "glEnableClientState",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(42),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 9,
		Name: "glDisableClientState",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(42),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 10,
		Name: "glGetProgramBinaryOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "bytes_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "binary_format",
				Type: schema.U32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "binary",
				Type: schema.Pointer,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 11,
		Name: "glProgramBinaryOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary_format",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary_size",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 12,
		Name: "glStartTilingQCOM",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "x",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "y",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "preserveMask",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 13,
		Name: "glEndTilingQCOM",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "preserve_mask",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 14,
		Name: "glDiscardFramebufferEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "numAttachments",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attachments",
				Type: s.getArrayInfo(2),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_discard_framebuffer.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 15,
		Name: "glInsertEventMarkerEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "marker",
				Type: schema.String,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 16,
		Name: "glPushGroupMarkerEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "marker",
				Type: schema.String,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type:             17,
		Name:             "glPopGroupMarkerEXT",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 18,
		Name: "glTexStorage1DEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "levels",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 19,
		Name: "glTexStorage2DEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "levels",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 20,
		Name: "glTexStorage3DEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "levels",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "depth",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 21,
		Name: "glTextureStorage1DEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "texture",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "levels",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 22,
		Name: "glTextureStorage2DEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "texture",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "levels",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 23,
		Name: "glTextureStorage3DEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "texture",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "levels",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "depth",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 24,
		Name: "glGenVertexArraysOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "arrays",
				Type: s.getArrayInfo(12),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 25,
		Name: "glBindVertexArrayOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "array",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 26,
		Name: "glDeleteVertexArraysOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "arrays",
				Type: s.getArrayInfo(12),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 27,
		Name: "glIsVertexArrayOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "array",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 28,
		Name: "glEGLImageTargetTexture2DOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(72),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "image",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 29,
		Name: "glEGLImageTargetRenderbufferStorageOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "image",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 30,
		Name: "glGetGraphicsResetStatusEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(74),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_robustness.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 31,
		Name: "glBindAttribLocation",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.String,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindAttribLocation.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 32,
		Name: "glBlendFunc",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "src_factor",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dst_factor",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFunc.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 33,
		Name: "glBlendFuncSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "src_factor_rgb",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dst_factor_rgb",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "src_factor_alpha",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dst_factor_alpha",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFuncSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 34,
		Name: "glBlendEquation",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "equation",
				Type: s.getEnumInfo(68),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquation.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 35,
		Name: "glBlendEquationSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "rgb",
				Type: s.getEnumInfo(68),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "alpha",
				Type: s.getEnumInfo(68),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquationSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 36,
		Name: "glBlendColor",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "red",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "green",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "blue",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "alpha",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendColor.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 37,
		Name: "glEnableVertexAttribArray",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 38,
		Name: "glDisableVertexAttribArray",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 39,
		Name: "glVertexAttribPointer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(45),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "normalized",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stride",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttribPointer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 40,
		Name: "glGetActiveAttrib",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_bytes_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "vector_count",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(46),
				Out:  true,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 41,
		Name: "glGetActiveUniform",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_bytes_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "size",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(47),
				Out:  true,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 42,
		Name: "glGetError",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(48),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetError.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 43,
		Name: "glGetProgramiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(52),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 44,
		Name: "glGetShaderiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(53),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderiv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 45,
		Name: "glGetUniformLocation",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.String,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniformLocation.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 46,
		Name: "glGetAttribLocation",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.String,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttribLocation.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 47,
		Name: "glPixelStorei",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(54),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPixelStorei.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 48,
		Name: "glTexParameteri",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(59),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 49,
		Name: "glTexParameterf",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(59),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 50,
		Name: "glGetTexParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(59),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 51,
		Name: "glGetTexParameterfv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(59),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(3),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 52,
		Name: "glUniform1i",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 53,
		Name: "glUniform2i",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 54,
		Name: "glUniform3i",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value2",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 55,
		Name: "glUniform4i",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value2",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value3",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 56,
		Name: "glUniform1iv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 57,
		Name: "glUniform2iv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 58,
		Name: "glUniform3iv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 59,
		Name: "glUniform4iv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 60,
		Name: "glUniform1f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 61,
		Name: "glUniform2f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 62,
		Name: "glUniform3f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value2",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 63,
		Name: "glUniform4f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value2",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value3",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 64,
		Name: "glUniform1fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 65,
		Name: "glUniform2fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 66,
		Name: "glUniform3fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 67,
		Name: "glUniform4fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 68,
		Name: "glUniformMatrix2fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "transpose",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 69,
		Name: "glUniformMatrix3fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "transpose",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 70,
		Name: "glUniformMatrix4fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "transpose",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 71,
		Name: "glGetUniformfv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 72,
		Name: "glGetUniformiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 73,
		Name: "glVertexAttrib1f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 74,
		Name: "glVertexAttrib2f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 75,
		Name: "glVertexAttrib3f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value2",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 76,
		Name: "glVertexAttrib4f",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value0",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value1",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value2",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value3",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 77,
		Name: "glVertexAttrib1fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 78,
		Name: "glVertexAttrib2fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 79,
		Name: "glVertexAttrib3fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 80,
		Name: "glVertexAttrib4fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 81,
		Name: "glGetShaderPrecisionFormat",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader_type",
				Type: s.getEnumInfo(33),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "precision_type",
				Type: s.getEnumInfo(64),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "range",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
			service.ParameterInfo{
				Name: "precision",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 82,
		Name: "glDepthMask",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "enabled",
				Type: schema.Bool,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthMask.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 83,
		Name: "glDepthFunc",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "function",
				Type: s.getEnumInfo(65),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthFunc.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 84,
		Name: "glDepthRangef",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "near",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "far",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthRangef.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 85,
		Name: "glColorMask",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "red",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "green",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "blue",
				Type: schema.Bool,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "alpha",
				Type: schema.Bool,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glColorMask.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 86,
		Name: "glStencilMask",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "mask",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMask.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 87,
		Name: "glStencilMaskSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "face",
				Type: s.getEnumInfo(39),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "mask",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMaskSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 88,
		Name: "glStencilFuncSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "face",
				Type: s.getEnumInfo(39),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "function",
				Type: s.getEnumInfo(65),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "reference_value",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "mask",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilFuncSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 89,
		Name: "glStencilOpSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "face",
				Type: s.getEnumInfo(39),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_fail",
				Type: s.getEnumInfo(66),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_pass_depth_fail",
				Type: s.getEnumInfo(66),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_pass_depth_pass",
				Type: s.getEnumInfo(66),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilOpSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 90,
		Name: "glFrontFace",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "orientation",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFrontFace.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 91,
		Name: "glViewport",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "x",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "y",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glViewport.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 92,
		Name: "glScissor",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "x",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "y",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glScissor.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 93,
		Name: "glActiveTexture",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "unit",
				Type: s.getEnumInfo(31),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glActiveTexture.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 94,
		Name: "glGenTextures",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "textures",
				Type: s.getArrayInfo(11),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenTextures.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 95,
		Name: "glDeleteTextures",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "textures",
				Type: s.getArrayInfo(11),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteTextures.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 96,
		Name: "glIsTexture",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "texture",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsTexture.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 97,
		Name: "glBindTexture",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "texture",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindTexture.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 98,
		Name: "glTexImage2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "internal_format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "border",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(19),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexImage2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 99,
		Name: "glTexSubImage2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "xoffset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "yoffset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(19),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexSubImage2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 100,
		Name: "glCopyTexImage2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(12),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "x",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "y",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "border",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexImage2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 101,
		Name: "glCopyTexSubImage2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "xoffset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "yoffset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "x",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "y",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 102,
		Name: "glCompressedTexImage2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(17),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "border",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "image_size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexImage2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 103,
		Name: "glCompressedTexSubImage2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "xoffset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "yoffset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(17),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "image_size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 104,
		Name: "glGenerateMipmap",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenerateMipmap.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 105,
		Name: "glReadPixels",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "x",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "y",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(9),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(19),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReadPixels.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 106,
		Name: "glGenFramebuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffers",
				Type: s.getArrayInfo(5),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenFramebuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 107,
		Name: "glBindFramebuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffer",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindFramebuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 108,
		Name: "glCheckFramebufferStatus",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(26),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 109,
		Name: "glDeleteFramebuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffers",
				Type: s.getArrayInfo(5),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteFramebuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 110,
		Name: "glIsFramebuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsFramebuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 111,
		Name: "glGenRenderbuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffers",
				Type: s.getArrayInfo(7),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenRenderbuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 112,
		Name: "glBindRenderbuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(27),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffer",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindRenderbuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 113,
		Name: "glRenderbufferStorage",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(27),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(13),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glRenderbufferStorage.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 114,
		Name: "glDeleteRenderbuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffers",
				Type: s.getArrayInfo(7),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 115,
		Name: "glIsRenderbuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "renderbuffer",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsRenderbuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 116,
		Name: "glGetRenderbufferParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(27),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(28),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 117,
		Name: "glGenBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffers",
				Type: s.getArrayInfo(1),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 118,
		Name: "glBindBuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindBuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 119,
		Name: "glBufferData",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "usage",
				Type: s.getEnumInfo(32),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferData.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 120,
		Name: "glBufferSubData",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "offset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "data",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferSubData.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 121,
		Name: "glDeleteBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffers",
				Type: s.getArrayInfo(1),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteBuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 122,
		Name: "glIsBuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "buffer",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsBuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 123,
		Name: "glGetBufferParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(30),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetBufferParameteriv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 124,
		Name: "glCreateShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(33),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateShader.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 125,
		Name: "glDeleteShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteShader.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 126,
		Name: "glShaderSource",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "source",
				Type: s.getArrayInfo(10),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "length",
				Type: s.getArrayInfo(8),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderSource.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 127,
		Name: "glShaderBinary",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shaders",
				Type: s.getArrayInfo(9),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary_format",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary_size",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderBinary.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 128,
		Name: "glGetShaderInfoLog",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "string_length_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "info",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderInfoLog.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 129,
		Name: "glGetShaderSource",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "string_length_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "source",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderSource.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type:             130,
		Name:             "glReleaseShaderCompiler",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 131,
		Name: "glCompileShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompileShader.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 132,
		Name: "glIsShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsShader.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 133,
		Name: "glCreateProgram",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 134,
		Name: "glDeleteProgram",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 135,
		Name: "glAttachShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glAttachShader.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 136,
		Name: "glDetachShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDetachShader.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 137,
		Name: "glGetAttachedShaders",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shaders_length_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "shaders",
				Type: s.getArrayInfo(9),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttachedShaders.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 138,
		Name: "glLinkProgram",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLinkProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 139,
		Name: "glGetProgramInfoLog",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffer_length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "string_length_written",
				Type: schema.S32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "info",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgramInfoLog.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 140,
		Name: "glUseProgram",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUseProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 141,
		Name: "glIsProgram",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 142,
		Name: "glValidateProgram",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glValidateProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 143,
		Name: "glClearColor",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "r",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "g",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "a",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearColor.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 144,
		Name: "glClearDepthf",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "depth",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearDepthf.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 145,
		Name: "glClearStencil",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "stencil",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearStencil.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 146,
		Name: "glClear",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "mask",
				Type: s.getEnumInfo(86),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClear.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 147,
		Name: "glCullFace",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "mode",
				Type: s.getEnumInfo(39),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCullFace.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 148,
		Name: "glPolygonOffset",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "scale_factor",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "units",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPolygonOffset.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 149,
		Name: "glLineWidth",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "width",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLineWidth.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 150,
		Name: "glSampleCoverage",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "value",
				Type: schema.Float,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "invert",
				Type: schema.Bool,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glSampleCoverage.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 151,
		Name: "glHint",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(49),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "mode",
				Type: s.getEnumInfo(50),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glHint.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 152,
		Name: "glFramebufferRenderbuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer_target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffer_attachment",
				Type: s.getEnumInfo(20),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffer_target",
				Type: s.getEnumInfo(27),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffer",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 153,
		Name: "glFramebufferTexture2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer_target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffer_attachment",
				Type: s.getEnumInfo(20),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "texture_target",
				Type: s.getEnumInfo(8),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "texture",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "level",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferTexture2D.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 154,
		Name: "glGetFramebufferAttachmentParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer_target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attachment",
				Type: s.getEnumInfo(20),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(25),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 155,
		Name: "glDrawElements",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "draw_mode",
				Type: s.getEnumInfo(0),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "element_count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "indices_type",
				Type: s.getEnumInfo(1),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "indices",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       true,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawElements.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 156,
		Name: "glDrawArrays",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "draw_mode",
				Type: s.getEnumInfo(0),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "first_index",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "index_count",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       true,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawArrays.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type:             157,
		Name:             "glFlush",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFlush.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type:             158,
		Name:             "glFinish",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFinish.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 159,
		Name: "glGetBooleanv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(38),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(0),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 160,
		Name: "glGetFloatv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(38),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(3),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 161,
		Name: "glGetIntegerv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(38),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: s.getArrayInfo(8),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 162,
		Name: "glGetString",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(44),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetString.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 163,
		Name: "glEnable",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "capability",
				Type: s.getEnumInfo(43),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnable.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 164,
		Name: "glDisable",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "capability",
				Type: s.getEnumInfo(43),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisable.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 165,
		Name: "glIsEnabled",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "capability",
				Type: s.getEnumInfo(43),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsEnabled.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 166,
		Name: "glMapBufferRange",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(69),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "offset",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "access",
				Type: s.getEnumInfo(87),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Pointer,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 167,
		Name: "glUnmapBuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(69),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 168,
		Name: "glInvalidateFramebuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(24),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attachments",
				Type: s.getArrayInfo(4),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glInvalidateFramebuffer.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 169,
		Name: "glRenderbufferStorageMultisample",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(27),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "samples",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "format",
				Type: s.getEnumInfo(13),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "width",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "height",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.opengl.org/registry/specs/EXT/framebuffer_multisample.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 170,
		Name: "glBlitFramebuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "srcX0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "srcY0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "srcX1",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "srcY1",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dstX0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dstY0",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dstX1",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dstY1",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "mask",
				Type: s.getEnumInfo(86),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "filter",
				Type: s.getEnumInfo(60),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBlitFramebuffer.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 171,
		Name: "glGenQueries",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: s.getArrayInfo(6),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenQueries.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 172,
		Name: "glBeginQuery",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBeginQuery.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 173,
		Name: "glEndQuery",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glEndQuery.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 174,
		Name: "glDeleteQueries",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: s.getArrayInfo(6),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteQueries.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 175,
		Name: "glIsQuery",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glIsQuery.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 176,
		Name: "glGetQueryiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(78),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryiv.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 177,
		Name: "glGetQueryObjectuiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(81),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryObjectuiv.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 178,
		Name: "glGenQueriesEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: s.getArrayInfo(6),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 179,
		Name: "glBeginQueryEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 180,
		Name: "glEndQueryEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 181,
		Name: "glDeleteQueriesEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: s.getArrayInfo(6),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 182,
		Name: "glIsQueryEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 183,
		Name: "glQueryCounterEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 184,
		Name: "glGetQueryivEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(84),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(78),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 185,
		Name: "glGetQueryObjectivEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(81),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 186,
		Name: "glGetQueryObjectuivEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(81),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 187,
		Name: "glGetQueryObjecti64vEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(81),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.S64,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 188,
		Name: "glGetQueryObjectui64vEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(81),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.U64,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAPI(api{}, service.StructInfo{
		Name: "state",
		Kind: service.TypeKindStruct,
		Fields: service.FieldInfoArray{
			&service.FieldInfo{
				Name: "Blending",
				Type: s.getClassInfo(27),
			},
			&service.FieldInfo{
				Name: "Rasterizing",
				Type: s.getClassInfo(28),
			},
			&service.FieldInfo{
				Name: "Clearing",
				Type: s.getClassInfo(29),
			},
			&service.FieldInfo{
				Name: "BoundFramebuffers",
				Type: s.getMapInfo(6),
			},
			&service.FieldInfo{
				Name: "BoundRenderbuffers",
				Type: s.getMapInfo(12),
			},
			&service.FieldInfo{
				Name: "BoundBuffers",
				Type: s.getMapInfo(2),
			},
			&service.FieldInfo{
				Name: "BoundProgram",
				Type: schema.U32,
			},
			&service.FieldInfo{
				Name: "BoundVertexArray",
				Type: schema.U32,
			},
			&service.FieldInfo{
				Name: "VertexAttributeArrays",
				Type: s.getMapInfo(23),
			},
			&service.FieldInfo{
				Name: "TextureUnits",
				Type: s.getMapInfo(18),
			},
			&service.FieldInfo{
				Name: "ActiveTextureUnit",
				Type: s.getEnumInfo(31),
			},
			&service.FieldInfo{
				Name: "Capabilities",
				Type: s.getMapInfo(1),
			},
			&service.FieldInfo{
				Name: "Internals",
				Type: s.getClassInfo(30),
			},
			&service.FieldInfo{
				Name: "GenerateMipmapHint",
				Type: s.getEnumInfo(50),
			},
			&service.FieldInfo{
				Name: "PixelStorage",
				Type: s.getMapInfo(14),
			},
			&service.FieldInfo{
				Name: "Instances",
				Type: s.getClassInfo(31),
			},
		},
	})
}

type schemaBuilder struct {
	arrays       map[int]*service.ArrayInfo
	staticArrays map[int]*service.StaticArrayInfo
	maps         map[int]*service.MapInfo
	enums        map[int]*service.EnumInfo
	structs      map[int]*service.StructInfo
	classes      map[int]*service.ClassInfo
}

func (s schemaBuilder) getArrayInfo(id int) *service.ArrayInfo {
	e, f := s.arrays[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateArrayInfo("BoolArray", service.TypeKindArray, schema.Bool)
	case 1:
		e = service.CreateArrayInfo("BufferIdArray", service.TypeKindArray, schema.U32)
	case 2:
		e = service.CreateArrayInfo("DiscardFramebufferAttachmentArray", service.TypeKindArray, s.getEnumInfo(51))
	case 3:
		e = service.CreateArrayInfo("F32Array", service.TypeKindArray, schema.Float)
	case 4:
		e = service.CreateArrayInfo("FramebufferAttachmentArray", service.TypeKindArray, s.getEnumInfo(20))
	case 5:
		e = service.CreateArrayInfo("FramebufferIdArray", service.TypeKindArray, schema.U32)
	case 6:
		e = service.CreateArrayInfo("QueryIdArray", service.TypeKindArray, schema.U32)
	case 7:
		e = service.CreateArrayInfo("RenderbufferIdArray", service.TypeKindArray, schema.U32)
	case 8:
		e = service.CreateArrayInfo("S32Array", service.TypeKindArray, schema.S32)
	case 9:
		e = service.CreateArrayInfo("ShaderIdArray", service.TypeKindArray, schema.U32)
	case 10:
		e = service.CreateArrayInfo("StringArray", service.TypeKindArray, schema.String)
	case 11:
		e = service.CreateArrayInfo("TextureIdArray", service.TypeKindArray, schema.U32)
	case 12:
		e = service.CreateArrayInfo("VertexArrayIdArray", service.TypeKindArray, schema.U32)
	}
	s.arrays[id] = e
	return e
}
func (s schemaBuilder) getStaticArrayInfo(id int) *service.StaticArrayInfo {
	e, f := s.staticArrays[id]
	if f {
		return e
	}
	switch id {
	}
	s.staticArrays[id] = e
	return e
}
func (s schemaBuilder) getMapInfo(id int) *service.MapInfo {
	e, f := s.maps[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateMapInfo("AttributeLocation_stringMap", service.TypeKindMap, schema.String, schema.U32)
	case 1:
		e = service.CreateMapInfo("Bool_CapabilityMap", service.TypeKindMap, s.getEnumInfo(43), schema.Bool)
	case 2:
		e = service.CreateMapInfo("BufferId_BufferTargetMap", service.TypeKindMap, s.getEnumInfo(29), schema.U32)
	case 3:
		e = service.CreateMapInfo("BufferRef_BufferIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(9))
	case 4:
		e = service.CreateMapInfo("CubemapLevel_s32Map", service.TypeKindMap, schema.S32, s.getClassInfo(6))
	case 5:
		e = service.CreateMapInfo("FramebufferAttachmentInfo_FramebufferAttachmentMap", service.TypeKindMap, s.getEnumInfo(20), s.getClassInfo(7))
	case 6:
		e = service.CreateMapInfo("FramebufferId_FramebufferTargetMap", service.TypeKindMap, s.getEnumInfo(24), schema.U32)
	case 7:
		e = service.CreateMapInfo("FramebufferRef_FramebufferIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(8))
	case 8:
		e = service.CreateMapInfo("Image_CubeMapImageTargetMap", service.TypeKindMap, s.getEnumInfo(6), s.getClassInfo(2))
	case 9:
		e = service.CreateMapInfo("Image_s32Map", service.TypeKindMap, schema.S32, s.getClassInfo(2))
	case 10:
		e = service.CreateMapInfo("ProgramRef_ProgramIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(23))
	case 11:
		e = service.CreateMapInfo("QueryRef_QueryIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(26))
	case 12:
		e = service.CreateMapInfo("RenderbufferId_RenderbufferTargetMap", service.TypeKindMap, s.getEnumInfo(27), schema.U32)
	case 13:
		e = service.CreateMapInfo("RenderbufferRef_RenderbufferIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(4))
	case 14:
		e = service.CreateMapInfo("S32_PixelStoreParameterMap", service.TypeKindMap, s.getEnumInfo(54), schema.S32)
	case 15:
		e = service.CreateMapInfo("ShaderId_ShaderTypeMap", service.TypeKindMap, s.getEnumInfo(33), schema.U32)
	case 16:
		e = service.CreateMapInfo("ShaderRef_ShaderIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(10))
	case 17:
		e = service.CreateMapInfo("TextureId_TextureTargetMap", service.TypeKindMap, s.getEnumInfo(5), schema.U32)
	case 18:
		e = service.CreateMapInfo("TextureId_TextureTargetMap_TextureUnitMap", service.TypeKindMap, s.getEnumInfo(31), s.getMapInfo(17))
	case 19:
		e = service.CreateMapInfo("TextureRef_TextureIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(5))
	case 20:
		e = service.CreateMapInfo("U32_FaceModeMap", service.TypeKindMap, s.getEnumInfo(39), schema.U32)
	case 21:
		e = service.CreateMapInfo("Uniform_UniformLocationMap", service.TypeKindMap, schema.S32, s.getClassInfo(22))
	case 22:
		e = service.CreateMapInfo("VertexArrayRef_VertexArrayIdMap", service.TypeKindMap, schema.U32, s.getClassInfo(24))
	case 23:
		e = service.CreateMapInfo("VertexAttributeArrayRef_AttributeLocationMap", service.TypeKindMap, schema.U32, s.getClassInfo(25))
	case 24:
		e = service.CreateMapInfo("VertexAttribute_s32Map", service.TypeKindMap, schema.S32, s.getClassInfo(11))
	}
	s.maps[id] = e
	return e
}
func (s schemaBuilder) getEnumInfo(id int) *service.EnumInfo {
	e, f := s.enums[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateEnumInfo(
			"DrawMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_LINE_LOOP",
					Value: 2,
				},
				service.EnumEntry{
					Name:  "GL_LINE_STRIP",
					Value: 3,
				},
				service.EnumEntry{
					Name:  "GL_LINES",
					Value: 1,
				},
				service.EnumEntry{
					Name:  "GL_POINTS",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "GL_TRIANGLE_FAN",
					Value: 6,
				},
				service.EnumEntry{
					Name:  "GL_TRIANGLE_STRIP",
					Value: 5,
				},
				service.EnumEntry{
					Name:  "GL_TRIANGLES",
					Value: 4,
				},
			},
			service.EnumInfoArray{},
		)
	case 1:
		e = service.CreateEnumInfo(
			"IndicesType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_UNSIGNED_BYTE",
					Value: 5121,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_SHORT",
					Value: 5123,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_INT",
					Value: 5125,
				},
			},
			service.EnumInfoArray{},
		)
	case 2:
		e = service.CreateEnumInfo(
			"TextureTarget_GLES_1_1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_2D",
					Value: 3553,
				},
			},
			service.EnumInfoArray{},
		)
	case 3:
		e = service.CreateEnumInfo(
			"TextureTarget_GLES_2_0",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP",
					Value: 34067,
				},
			},
			service.EnumInfoArray{},
		)
	case 4:
		e = service.CreateEnumInfo(
			"TextureTarget_OES_EGL_image_external",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_EXTERNAL_OES",
					Value: 36197,
				},
			},
			service.EnumInfoArray{},
		)
	case 5:
		e = service.CreateEnumInfo(
			"TextureTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(2),
				s.getEnumInfo(3),
				s.getEnumInfo(4),
			},
		)
	case 6:
		e = service.CreateEnumInfo(
			"CubeMapImageTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP_NEGATIVE_X",
					Value: 34070,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP_NEGATIVE_Y",
					Value: 34072,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP_NEGATIVE_Z",
					Value: 34074,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP_POSITIVE_X",
					Value: 34069,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP_POSITIVE_Y",
					Value: 34071,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_CUBE_MAP_POSITIVE_Z",
					Value: 34073,
				},
			},
			service.EnumInfoArray{},
		)
	case 7:
		e = service.CreateEnumInfo(
			"Texture2DImageTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_2D",
					Value: 3553,
				},
			},
			service.EnumInfoArray{},
		)
	case 8:
		e = service.CreateEnumInfo(
			"TextureImageTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(6),
				s.getEnumInfo(7),
			},
		)
	case 9:
		e = service.CreateEnumInfo(
			"BaseTexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ALPHA",
					Value: 6406,
				},
				service.EnumEntry{
					Name:  "GL_RGB",
					Value: 6407,
				},
				service.EnumEntry{
					Name:  "GL_RGBA",
					Value: 6408,
				},
			},
			service.EnumInfoArray{},
		)
	case 10:
		e = service.CreateEnumInfo(
			"TexelFormat_GLES_1_1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_LUMINANCE",
					Value: 6409,
				},
				service.EnumEntry{
					Name:  "GL_LUMINANCE_ALPHA",
					Value: 6410,
				},
			},
			service.EnumInfoArray{
				s.getEnumInfo(9),
			},
		)
	case 11:
		e = service.CreateEnumInfo(
			"TexelFormat_GLES_3_0",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RED",
					Value: 6403,
				},
				service.EnumEntry{
					Name:  "GL_RED_INTEGER",
					Value: 36244,
				},
				service.EnumEntry{
					Name:  "GL_RG",
					Value: 33319,
				},
				service.EnumEntry{
					Name:  "GL_RG_INTEGER",
					Value: 33320,
				},
				service.EnumEntry{
					Name:  "GL_RGB_INTEGER",
					Value: 36248,
				},
				service.EnumEntry{
					Name:  "GL_RGBA_INTEGER",
					Value: 36249,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_COMPONENT",
					Value: 6402,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_COMPONENT16",
					Value: 33189,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_STENCIL",
					Value: 34041,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH24_STENCIL8",
					Value: 35056,
				},
			},
			service.EnumInfoArray{},
		)
	case 12:
		e = service.CreateEnumInfo(
			"TexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(10),
				s.getEnumInfo(11),
			},
		)
	case 13:
		e = service.CreateEnumInfo(
			"RenderbufferFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RGBA4",
					Value: 32854,
				},
				service.EnumEntry{
					Name:  "GL_RGB5_A1",
					Value: 32855,
				},
				service.EnumEntry{
					Name:  "GL_RGB565",
					Value: 36194,
				},
				service.EnumEntry{
					Name:  "GL_RGBA8",
					Value: 32856,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_COMPONENT16",
					Value: 33189,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_INDEX8",
					Value: 36168,
				},
			},
			service.EnumInfoArray{},
		)
	case 14:
		e = service.CreateEnumInfo(
			"Type_OES_vertex_half_float",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_HALF_FLOAT_OES",
					Value: 36193,
				},
			},
			service.EnumInfoArray{},
		)
	case 15:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ETC1_RGB8_OES",
					Value: 36196,
				},
			},
			service.EnumInfoArray{},
		)
	case 16:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_AMD_compressed_ATC_texture",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ATC_RGB_AMD",
					Value: 35986,
				},
				service.EnumEntry{
					Name:  "GL_ATC_RGBA_EXPLICIT_ALPHA_AMD",
					Value: 35987,
				},
				service.EnumEntry{
					Name:  "GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD",
					Value: 34798,
				},
			},
			service.EnumInfoArray{},
		)
	case 17:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(15),
				s.getEnumInfo(16),
			},
		)
	case 18:
		e = service.CreateEnumInfo(
			"ImageTexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(12),
				s.getEnumInfo(17),
			},
		)
	case 19:
		e = service.CreateEnumInfo(
			"TexelType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_UNSIGNED_BYTE",
					Value: 5121,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_SHORT",
					Value: 5123,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_INT",
					Value: 5125,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT",
					Value: 5126,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_SHORT_4_4_4_4",
					Value: 32819,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_SHORT_5_5_5_1",
					Value: 32820,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_SHORT_5_6_5",
					Value: 33635,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_INT_24_8",
					Value: 34042,
				},
			},
			service.EnumInfoArray{},
		)
	case 20:
		e = service.CreateEnumInfo(
			"FramebufferAttachment",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COLOR_ATTACHMENT0",
					Value: 36064,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_ATTACHMENT",
					Value: 36096,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_ATTACHMENT",
					Value: 36128,
				},
			},
			service.EnumInfoArray{},
		)
	case 21:
		e = service.CreateEnumInfo(
			"FramebufferAttachmentType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_NONE",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER",
					Value: 36161,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE",
					Value: 5890,
				},
			},
			service.EnumInfoArray{},
		)
	case 22:
		e = service.CreateEnumInfo(
			"FramebufferTarget_GLES_2_0",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER",
					Value: 36160,
				},
			},
			service.EnumInfoArray{},
		)
	case 23:
		e = service.CreateEnumInfo(
			"FramebufferTarget_GLES_3_1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_READ_FRAMEBUFFER",
					Value: 36008,
				},
				service.EnumEntry{
					Name:  "GL_DRAW_FRAMEBUFFER",
					Value: 36009,
				},
			},
			service.EnumInfoArray{},
		)
	case 24:
		e = service.CreateEnumInfo(
			"FramebufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(22),
				s.getEnumInfo(23),
			},
		)
	case 25:
		e = service.CreateEnumInfo(
			"FramebufferAttachmentParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE",
					Value: 36048,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME",
					Value: 36049,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL",
					Value: 36050,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE",
					Value: 36051,
				},
			},
			service.EnumInfoArray{},
		)
	case 26:
		e = service.CreateEnumInfo(
			"FramebufferStatus",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_COMPLETE",
					Value: 36053,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT",
					Value: 36054,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT",
					Value: 36055,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS",
					Value: 36057,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_UNSUPPORTED",
					Value: 36061,
				},
			},
			service.EnumInfoArray{},
		)
	case 27:
		e = service.CreateEnumInfo(
			"RenderbufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER",
					Value: 36161,
				},
			},
			service.EnumInfoArray{},
		)
	case 28:
		e = service.CreateEnumInfo(
			"RenderbufferParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_WIDTH",
					Value: 36162,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_HEIGHT",
					Value: 36163,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_INTERNAL_FORMAT",
					Value: 36164,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_RED_SIZE",
					Value: 36176,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_GREEN_SIZE",
					Value: 36177,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_BLUE_SIZE",
					Value: 36178,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_ALPHA_SIZE",
					Value: 36179,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_DEPTH_SIZE",
					Value: 36180,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_STENCIL_SIZE",
					Value: 36181,
				},
			},
			service.EnumInfoArray{},
		)
	case 29:
		e = service.CreateEnumInfo(
			"BufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ARRAY_BUFFER",
					Value: 34962,
				},
				service.EnumEntry{
					Name:  "GL_ELEMENT_ARRAY_BUFFER",
					Value: 34963,
				},
			},
			service.EnumInfoArray{},
		)
	case 30:
		e = service.CreateEnumInfo(
			"BufferParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_BUFFER_SIZE",
					Value: 34660,
				},
				service.EnumEntry{
					Name:  "GL_BUFFER_USAGE",
					Value: 34661,
				},
			},
			service.EnumInfoArray{},
		)
	case 31:
		e = service.CreateEnumInfo(
			"TextureUnit",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE0",
					Value: 33984,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE1",
					Value: 33985,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE2",
					Value: 33986,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE3",
					Value: 33987,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE4",
					Value: 33988,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE5",
					Value: 33989,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE6",
					Value: 33990,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE7",
					Value: 33991,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE8",
					Value: 33992,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE9",
					Value: 33993,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE10",
					Value: 33994,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE11",
					Value: 33995,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE12",
					Value: 33996,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE13",
					Value: 33997,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE14",
					Value: 33998,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE15",
					Value: 33999,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE16",
					Value: 34000,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE17",
					Value: 34001,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE18",
					Value: 34002,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE19",
					Value: 34003,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE20",
					Value: 34004,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE21",
					Value: 34005,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE22",
					Value: 34006,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE23",
					Value: 34007,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE24",
					Value: 34008,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE25",
					Value: 34009,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE26",
					Value: 34010,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE27",
					Value: 34011,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE28",
					Value: 34012,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE29",
					Value: 34013,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE30",
					Value: 34014,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE31",
					Value: 34015,
				},
			},
			service.EnumInfoArray{},
		)
	case 32:
		e = service.CreateEnumInfo(
			"BufferUsage",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_DYNAMIC_DRAW",
					Value: 35048,
				},
				service.EnumEntry{
					Name:  "GL_STATIC_DRAW",
					Value: 35044,
				},
				service.EnumEntry{
					Name:  "GL_STREAM_DRAW",
					Value: 35040,
				},
			},
			service.EnumInfoArray{},
		)
	case 33:
		e = service.CreateEnumInfo(
			"ShaderType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_VERTEX_SHADER",
					Value: 35633,
				},
				service.EnumEntry{
					Name:  "GL_FRAGMENT_SHADER",
					Value: 35632,
				},
			},
			service.EnumInfoArray{},
		)
	case 34:
		e = service.CreateEnumInfo(
			"StateVariable_GLES_2_0",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ACTIVE_TEXTURE",
					Value: 34016,
				},
				service.EnumEntry{
					Name:  "GL_ALIASED_LINE_WIDTH_RANGE",
					Value: 33902,
				},
				service.EnumEntry{
					Name:  "GL_ALIASED_POINT_SIZE_RANGE",
					Value: 33901,
				},
				service.EnumEntry{
					Name:  "GL_ALPHA_BITS",
					Value: 3413,
				},
				service.EnumEntry{
					Name:  "GL_ARRAY_BUFFER_BINDING",
					Value: 34964,
				},
				service.EnumEntry{
					Name:  "GL_BLEND",
					Value: 3042,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_COLOR",
					Value: 32773,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_DST_ALPHA",
					Value: 32970,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_DST_RGB",
					Value: 32968,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_EQUATION_ALPHA",
					Value: 34877,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_EQUATION_RGB",
					Value: 32777,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_SRC_ALPHA",
					Value: 32971,
				},
				service.EnumEntry{
					Name:  "GL_BLEND_SRC_RGB",
					Value: 32969,
				},
				service.EnumEntry{
					Name:  "GL_BLUE_BITS",
					Value: 3412,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_CLEAR_VALUE",
					Value: 3106,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_WRITEMASK",
					Value: 3107,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_TEXTURE_FORMATS",
					Value: 34467,
				},
				service.EnumEntry{
					Name:  "GL_CULL_FACE",
					Value: 2884,
				},
				service.EnumEntry{
					Name:  "GL_CULL_FACE_MODE",
					Value: 2885,
				},
				service.EnumEntry{
					Name:  "GL_CURRENT_PROGRAM",
					Value: 35725,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BITS",
					Value: 3414,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_CLEAR_VALUE",
					Value: 2931,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_FUNC",
					Value: 2932,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_RANGE",
					Value: 2928,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_TEST",
					Value: 2929,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_WRITEMASK",
					Value: 2930,
				},
				service.EnumEntry{
					Name:  "GL_DITHER",
					Value: 3024,
				},
				service.EnumEntry{
					Name:  "GL_ELEMENT_ARRAY_BUFFER_BINDING",
					Value: 34965,
				},
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER_BINDING",
					Value: 36006,
				},
				service.EnumEntry{
					Name:  "GL_FRONT_FACE",
					Value: 2886,
				},
				service.EnumEntry{
					Name:  "GL_GENERATE_MIPMAP_HINT",
					Value: 33170,
				},
				service.EnumEntry{
					Name:  "GL_GREEN_BITS",
					Value: 3411,
				},
				service.EnumEntry{
					Name:  "GL_IMPLEMENTATION_COLOR_READ_FORMAT",
					Value: 35739,
				},
				service.EnumEntry{
					Name:  "GL_IMPLEMENTATION_COLOR_READ_TYPE",
					Value: 35738,
				},
				service.EnumEntry{
					Name:  "GL_LINE_WIDTH",
					Value: 2849,
				},
				service.EnumEntry{
					Name:  "GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS",
					Value: 35661,
				},
				service.EnumEntry{
					Name:  "GL_MAX_CUBE_MAP_TEXTURE_SIZE",
					Value: 34076,
				},
				service.EnumEntry{
					Name:  "GL_MAX_FRAGMENT_UNIFORM_VECTORS",
					Value: 36349,
				},
				service.EnumEntry{
					Name:  "GL_MAX_RENDERBUFFER_SIZE",
					Value: 34024,
				},
				service.EnumEntry{
					Name:  "GL_MAX_TEXTURE_IMAGE_UNITS",
					Value: 34930,
				},
				service.EnumEntry{
					Name:  "GL_MAX_TEXTURE_SIZE",
					Value: 3379,
				},
				service.EnumEntry{
					Name:  "GL_MAX_VARYING_VECTORS",
					Value: 36348,
				},
				service.EnumEntry{
					Name:  "GL_MAX_VERTEX_ATTRIBS",
					Value: 34921,
				},
				service.EnumEntry{
					Name:  "GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS",
					Value: 35660,
				},
				service.EnumEntry{
					Name:  "GL_MAX_VERTEX_UNIFORM_VECTORS",
					Value: 36347,
				},
				service.EnumEntry{
					Name:  "GL_MAX_VIEWPORT_DIMS",
					Value: 3386,
				},
				service.EnumEntry{
					Name:  "GL_NUM_COMPRESSED_TEXTURE_FORMATS",
					Value: 34466,
				},
				service.EnumEntry{
					Name:  "GL_NUM_SHADER_BINARY_FORMATS",
					Value: 36345,
				},
				service.EnumEntry{
					Name:  "GL_PACK_ALIGNMENT",
					Value: 3333,
				},
				service.EnumEntry{
					Name:  "GL_POLYGON_OFFSET_FACTOR",
					Value: 32824,
				},
				service.EnumEntry{
					Name:  "GL_POLYGON_OFFSET_FILL",
					Value: 32823,
				},
				service.EnumEntry{
					Name:  "GL_POLYGON_OFFSET_UNITS",
					Value: 10752,
				},
				service.EnumEntry{
					Name:  "GL_RED_BITS",
					Value: 3410,
				},
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_BINDING",
					Value: 36007,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_ALPHA_TO_COVERAGE",
					Value: 32926,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_BUFFERS",
					Value: 32936,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_COVERAGE",
					Value: 32928,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_COVERAGE_INVERT",
					Value: 32939,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_COVERAGE_VALUE",
					Value: 32938,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLES",
					Value: 32937,
				},
				service.EnumEntry{
					Name:  "GL_SCISSOR_BOX",
					Value: 3088,
				},
				service.EnumEntry{
					Name:  "GL_SCISSOR_TEST",
					Value: 3089,
				},
				service.EnumEntry{
					Name:  "GL_SHADER_BINARY_FORMATS",
					Value: 36344,
				},
				service.EnumEntry{
					Name:  "GL_SHADER_COMPILER",
					Value: 36346,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_FAIL",
					Value: 34817,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_FUNC",
					Value: 34816,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_PASS_DEPTH_FAIL",
					Value: 34818,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_PASS_DEPTH_PASS",
					Value: 34819,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_REF",
					Value: 36003,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_VALUE_MASK",
					Value: 36004,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BACK_WRITEMASK",
					Value: 36005,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BITS",
					Value: 3415,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_CLEAR_VALUE",
					Value: 2961,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_FAIL",
					Value: 2964,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_FUNC",
					Value: 2962,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_PASS_DEPTH_FAIL",
					Value: 2965,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_PASS_DEPTH_PASS",
					Value: 2966,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_REF",
					Value: 2967,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_TEST",
					Value: 2960,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_VALUE_MASK",
					Value: 2963,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_WRITEMASK",
					Value: 2968,
				},
				service.EnumEntry{
					Name:  "GL_SUBPIXEL_BITS",
					Value: 3408,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_BINDING_2D",
					Value: 32873,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_BINDING_CUBE_MAP",
					Value: 34068,
				},
				service.EnumEntry{
					Name:  "GL_UNPACK_ALIGNMENT",
					Value: 3317,
				},
				service.EnumEntry{
					Name:  "GL_VIEWPORT",
					Value: 2978,
				},
			},
			service.EnumInfoArray{},
		)
	case 35:
		e = service.CreateEnumInfo(
			"StateVariable_GLES_3_1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_READ_FRAMEBUFFER_BINDING",
					Value: 36010,
				},
			},
			service.EnumInfoArray{},
		)
	case 36:
		e = service.CreateEnumInfo(
			"StateVariable_EXT_texture_filter_anisotropic",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT",
					Value: 34047,
				},
			},
			service.EnumInfoArray{},
		)
	case 37:
		e = service.CreateEnumInfo(
			"StateVariable_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_GPU_DISJOINT_EXT",
					Value: 36795,
				},
			},
			service.EnumInfoArray{},
		)
	case 38:
		e = service.CreateEnumInfo(
			"StateVariable",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(34),
				s.getEnumInfo(35),
				s.getEnumInfo(36),
				s.getEnumInfo(37),
			},
		)
	case 39:
		e = service.CreateEnumInfo(
			"FaceMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FRONT",
					Value: 1028,
				},
				service.EnumEntry{
					Name:  "GL_BACK",
					Value: 1029,
				},
				service.EnumEntry{
					Name:  "GL_FRONT_AND_BACK",
					Value: 1032,
				},
			},
			service.EnumInfoArray{},
		)
	case 40:
		e = service.CreateEnumInfo(
			"ArrayType_GLES_1_1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_VERTEX_ARRAY",
					Value: 32884,
				},
				service.EnumEntry{
					Name:  "GL_NORMAL_ARRAY",
					Value: 32885,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_ARRAY",
					Value: 32886,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_COORD_ARRAY",
					Value: 32888,
				},
			},
			service.EnumInfoArray{},
		)
	case 41:
		e = service.CreateEnumInfo(
			"ArrayType_OES_point_size_array",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_POINT_SIZE_ARRAY_OES",
					Value: 35740,
				},
			},
			service.EnumInfoArray{},
		)
	case 42:
		e = service.CreateEnumInfo(
			"ArrayType",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(40),
				s.getEnumInfo(41),
			},
		)
	case 43:
		e = service.CreateEnumInfo(
			"Capability",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_BLEND",
					Value: 3042,
				},
				service.EnumEntry{
					Name:  "GL_CULL_FACE",
					Value: 2884,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_TEST",
					Value: 2929,
				},
				service.EnumEntry{
					Name:  "GL_DITHER",
					Value: 3024,
				},
				service.EnumEntry{
					Name:  "GL_POLYGON_OFFSET_FILL",
					Value: 32823,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_ALPHA_TO_COVERAGE",
					Value: 32926,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLE_COVERAGE",
					Value: 32928,
				},
				service.EnumEntry{
					Name:  "GL_SCISSOR_TEST",
					Value: 3089,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_TEST",
					Value: 2960,
				},
			},
			service.EnumInfoArray{
				s.getEnumInfo(42),
			},
		)
	case 44:
		e = service.CreateEnumInfo(
			"StringConstant",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_EXTENSIONS",
					Value: 7939,
				},
				service.EnumEntry{
					Name:  "GL_RENDERER",
					Value: 7937,
				},
				service.EnumEntry{
					Name:  "GL_VENDOR",
					Value: 7936,
				},
				service.EnumEntry{
					Name:  "GL_VERSION",
					Value: 7938,
				},
			},
			service.EnumInfoArray{},
		)
	case 45:
		e = service.CreateEnumInfo(
			"VertexAttribType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_BYTE",
					Value: 5120,
				},
				service.EnumEntry{
					Name:  "GL_FIXED",
					Value: 5132,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT",
					Value: 5126,
				},
				service.EnumEntry{
					Name:  "GL_SHORT",
					Value: 5122,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_BYTE",
					Value: 5121,
				},
				service.EnumEntry{
					Name:  "GL_UNSIGNED_SHORT",
					Value: 5123,
				},
			},
			service.EnumInfoArray{
				s.getEnumInfo(14),
			},
		)
	case 46:
		e = service.CreateEnumInfo(
			"ShaderAttribType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FLOAT",
					Value: 5126,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_VEC2",
					Value: 35664,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_VEC3",
					Value: 35665,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_VEC4",
					Value: 35666,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_MAT2",
					Value: 35674,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_MAT3",
					Value: 35675,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_MAT4",
					Value: 35676,
				},
			},
			service.EnumInfoArray{},
		)
	case 47:
		e = service.CreateEnumInfo(
			"ShaderUniformType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FLOAT",
					Value: 5126,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_VEC2",
					Value: 35664,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_VEC3",
					Value: 35665,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_VEC4",
					Value: 35666,
				},
				service.EnumEntry{
					Name:  "GL_INT",
					Value: 5124,
				},
				service.EnumEntry{
					Name:  "GL_INT_VEC2",
					Value: 35667,
				},
				service.EnumEntry{
					Name:  "GL_INT_VEC3",
					Value: 35668,
				},
				service.EnumEntry{
					Name:  "GL_INT_VEC4",
					Value: 35669,
				},
				service.EnumEntry{
					Name:  "GL_BOOL",
					Value: 35670,
				},
				service.EnumEntry{
					Name:  "GL_BOOL_VEC2",
					Value: 35671,
				},
				service.EnumEntry{
					Name:  "GL_BOOL_VEC3",
					Value: 35672,
				},
				service.EnumEntry{
					Name:  "GL_BOOL_VEC4",
					Value: 35673,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_MAT2",
					Value: 35674,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_MAT3",
					Value: 35675,
				},
				service.EnumEntry{
					Name:  "GL_FLOAT_MAT4",
					Value: 35676,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLER_2D",
					Value: 35678,
				},
				service.EnumEntry{
					Name:  "GL_SAMPLER_CUBE",
					Value: 35680,
				},
			},
			service.EnumInfoArray{},
		)
	case 48:
		e = service.CreateEnumInfo(
			"Error",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_NO_ERROR",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "GL_INVALID_ENUM",
					Value: 1280,
				},
				service.EnumEntry{
					Name:  "GL_INVALID_VALUE",
					Value: 1281,
				},
				service.EnumEntry{
					Name:  "GL_INVALID_OPERATION",
					Value: 1282,
				},
				service.EnumEntry{
					Name:  "GL_INVALID_FRAMEBUFFER_OPERATION",
					Value: 1286,
				},
				service.EnumEntry{
					Name:  "GL_OUT_OF_MEMORY",
					Value: 1285,
				},
			},
			service.EnumInfoArray{},
		)
	case 49:
		e = service.CreateEnumInfo(
			"HintTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_GENERATE_MIPMAP_HINT",
					Value: 33170,
				},
			},
			service.EnumInfoArray{},
		)
	case 50:
		e = service.CreateEnumInfo(
			"HintMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_DONT_CARE",
					Value: 4352,
				},
				service.EnumEntry{
					Name:  "GL_FASTEST",
					Value: 4353,
				},
				service.EnumEntry{
					Name:  "GL_NICEST",
					Value: 4354,
				},
			},
			service.EnumInfoArray{},
		)
	case 51:
		e = service.CreateEnumInfo(
			"DiscardFramebufferAttachment",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COLOR_EXT",
					Value: 6144,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_EXT",
					Value: 6145,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_EXT",
					Value: 6146,
				},
			},
			service.EnumInfoArray{},
		)
	case 52:
		e = service.CreateEnumInfo(
			"ProgramParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_DELETE_STATUS",
					Value: 35712,
				},
				service.EnumEntry{
					Name:  "GL_LINK_STATUS",
					Value: 35714,
				},
				service.EnumEntry{
					Name:  "GL_VALIDATE_STATUS",
					Value: 35715,
				},
				service.EnumEntry{
					Name:  "GL_INFO_LOG_LENGTH",
					Value: 35716,
				},
				service.EnumEntry{
					Name:  "GL_ATTACHED_SHADERS",
					Value: 35717,
				},
				service.EnumEntry{
					Name:  "GL_ACTIVE_ATTRIBUTES",
					Value: 35721,
				},
				service.EnumEntry{
					Name:  "GL_ACTIVE_ATTRIBUTE_MAX_LENGTH",
					Value: 35722,
				},
				service.EnumEntry{
					Name:  "GL_ACTIVE_UNIFORMS",
					Value: 35718,
				},
				service.EnumEntry{
					Name:  "GL_ACTIVE_UNIFORM_MAX_LENGTH",
					Value: 35719,
				},
			},
			service.EnumInfoArray{},
		)
	case 53:
		e = service.CreateEnumInfo(
			"ShaderParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_SHADER_TYPE",
					Value: 35663,
				},
				service.EnumEntry{
					Name:  "GL_DELETE_STATUS",
					Value: 35712,
				},
				service.EnumEntry{
					Name:  "GL_COMPILE_STATUS",
					Value: 35713,
				},
				service.EnumEntry{
					Name:  "GL_INFO_LOG_LENGTH",
					Value: 35716,
				},
				service.EnumEntry{
					Name:  "GL_SHADER_SOURCE_LENGTH",
					Value: 35720,
				},
			},
			service.EnumInfoArray{},
		)
	case 54:
		e = service.CreateEnumInfo(
			"PixelStoreParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_PACK_ALIGNMENT",
					Value: 3333,
				},
				service.EnumEntry{
					Name:  "GL_UNPACK_ALIGNMENT",
					Value: 3317,
				},
			},
			service.EnumInfoArray{},
		)
	case 55:
		e = service.CreateEnumInfo(
			"TextureParameter_FilterMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_MIN_FILTER",
					Value: 10241,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_MAG_FILTER",
					Value: 10240,
				},
			},
			service.EnumInfoArray{},
		)
	case 56:
		e = service.CreateEnumInfo(
			"TextureParameter_WrapMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_WRAP_S",
					Value: 10242,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_WRAP_T",
					Value: 10243,
				},
			},
			service.EnumInfoArray{},
		)
	case 57:
		e = service.CreateEnumInfo(
			"TextureParameter_EXT_texture_filter_anisotropic",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_MAX_ANISOTROPY_EXT",
					Value: 34046,
				},
			},
			service.EnumInfoArray{},
		)
	case 58:
		e = service.CreateEnumInfo(
			"TextureParameter_SwizzleMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_SWIZZLE_R",
					Value: 36418,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_SWIZZLE_G",
					Value: 36419,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_SWIZZLE_B",
					Value: 36420,
				},
				service.EnumEntry{
					Name:  "GL_TEXTURE_SWIZZLE_A",
					Value: 36421,
				},
			},
			service.EnumInfoArray{},
		)
	case 59:
		e = service.CreateEnumInfo(
			"TextureParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(55),
				s.getEnumInfo(56),
				s.getEnumInfo(58),
				s.getEnumInfo(57),
			},
		)
	case 60:
		e = service.CreateEnumInfo(
			"TextureFilterMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_NEAREST",
					Value: 9728,
				},
				service.EnumEntry{
					Name:  "GL_LINEAR",
					Value: 9729,
				},
				service.EnumEntry{
					Name:  "GL_NEAREST_MIPMAP_NEAREST",
					Value: 9984,
				},
				service.EnumEntry{
					Name:  "GL_LINEAR_MIPMAP_NEAREST",
					Value: 9985,
				},
				service.EnumEntry{
					Name:  "GL_NEAREST_MIPMAP_LINEAR",
					Value: 9986,
				},
				service.EnumEntry{
					Name:  "GL_LINEAR_MIPMAP_LINEAR",
					Value: 9987,
				},
			},
			service.EnumInfoArray{},
		)
	case 61:
		e = service.CreateEnumInfo(
			"TextureWrapMode",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_CLAMP_TO_EDGE",
					Value: 33071,
				},
				service.EnumEntry{
					Name:  "GL_MIRRORED_REPEAT",
					Value: 33648,
				},
				service.EnumEntry{
					Name:  "GL_REPEAT",
					Value: 10497,
				},
			},
			service.EnumInfoArray{},
		)
	case 62:
		e = service.CreateEnumInfo(
			"TexelComponent",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RED",
					Value: 6403,
				},
				service.EnumEntry{
					Name:  "GL_GREEN",
					Value: 6404,
				},
				service.EnumEntry{
					Name:  "GL_BLUE",
					Value: 6405,
				},
				service.EnumEntry{
					Name:  "GL_ALPHA",
					Value: 6406,
				},
			},
			service.EnumInfoArray{},
		)
	case 63:
		e = service.CreateEnumInfo(
			"BlendFactor",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ZERO",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "GL_ONE",
					Value: 1,
				},
				service.EnumEntry{
					Name:  "GL_SRC_COLOR",
					Value: 768,
				},
				service.EnumEntry{
					Name:  "GL_ONE_MINUS_SRC_COLOR",
					Value: 769,
				},
				service.EnumEntry{
					Name:  "GL_DST_COLOR",
					Value: 774,
				},
				service.EnumEntry{
					Name:  "GL_ONE_MINUS_DST_COLOR",
					Value: 775,
				},
				service.EnumEntry{
					Name:  "GL_SRC_ALPHA",
					Value: 770,
				},
				service.EnumEntry{
					Name:  "GL_ONE_MINUS_SRC_ALPHA",
					Value: 771,
				},
				service.EnumEntry{
					Name:  "GL_DST_ALPHA",
					Value: 772,
				},
				service.EnumEntry{
					Name:  "GL_ONE_MINUS_DST_ALPHA",
					Value: 773,
				},
				service.EnumEntry{
					Name:  "GL_CONSTANT_COLOR",
					Value: 32769,
				},
				service.EnumEntry{
					Name:  "GL_ONE_MINUS_CONSTANT_COLOR",
					Value: 32770,
				},
				service.EnumEntry{
					Name:  "GL_CONSTANT_ALPHA",
					Value: 32771,
				},
				service.EnumEntry{
					Name:  "GL_ONE_MINUS_CONSTANT_ALPHA",
					Value: 32772,
				},
				service.EnumEntry{
					Name:  "GL_SRC_ALPHA_SATURATE",
					Value: 776,
				},
			},
			service.EnumInfoArray{},
		)
	case 64:
		e = service.CreateEnumInfo(
			"PrecisionType",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_LOW_FLOAT",
					Value: 36336,
				},
				service.EnumEntry{
					Name:  "GL_MEDIUM_FLOAT",
					Value: 36337,
				},
				service.EnumEntry{
					Name:  "GL_HIGH_FLOAT",
					Value: 36338,
				},
				service.EnumEntry{
					Name:  "GL_LOW_INT",
					Value: 36339,
				},
				service.EnumEntry{
					Name:  "GL_MEDIUM_INT",
					Value: 36340,
				},
				service.EnumEntry{
					Name:  "GL_HIGH_INT",
					Value: 36341,
				},
			},
			service.EnumInfoArray{},
		)
	case 65:
		e = service.CreateEnumInfo(
			"TestFunction",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_NEVER",
					Value: 512,
				},
				service.EnumEntry{
					Name:  "GL_LESS",
					Value: 513,
				},
				service.EnumEntry{
					Name:  "GL_EQUAL",
					Value: 514,
				},
				service.EnumEntry{
					Name:  "GL_LEQUAL",
					Value: 515,
				},
				service.EnumEntry{
					Name:  "GL_GREATER",
					Value: 516,
				},
				service.EnumEntry{
					Name:  "GL_NOTEQUAL",
					Value: 517,
				},
				service.EnumEntry{
					Name:  "GL_GEQUAL",
					Value: 518,
				},
				service.EnumEntry{
					Name:  "GL_ALWAYS",
					Value: 519,
				},
			},
			service.EnumInfoArray{},
		)
	case 66:
		e = service.CreateEnumInfo(
			"StencilAction",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_KEEP",
					Value: 7680,
				},
				service.EnumEntry{
					Name:  "GL_ZERO",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "GL_REPLACE",
					Value: 7681,
				},
				service.EnumEntry{
					Name:  "GL_INCR",
					Value: 7682,
				},
				service.EnumEntry{
					Name:  "GL_INCR_WRAP",
					Value: 34055,
				},
				service.EnumEntry{
					Name:  "GL_DECR",
					Value: 7683,
				},
				service.EnumEntry{
					Name:  "GL_DECR_WRAP",
					Value: 34056,
				},
				service.EnumEntry{
					Name:  "GL_INVERT",
					Value: 5386,
				},
			},
			service.EnumInfoArray{},
		)
	case 67:
		e = service.CreateEnumInfo(
			"FaceOrientation",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_CW",
					Value: 2304,
				},
				service.EnumEntry{
					Name:  "GL_CCW",
					Value: 2305,
				},
			},
			service.EnumInfoArray{},
		)
	case 68:
		e = service.CreateEnumInfo(
			"BlendEquation",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FUNC_ADD",
					Value: 32774,
				},
				service.EnumEntry{
					Name:  "GL_FUNC_SUBTRACT",
					Value: 32778,
				},
				service.EnumEntry{
					Name:  "GL_FUNC_REVERSE_SUBTRACT",
					Value: 32779,
				},
			},
			service.EnumInfoArray{},
		)
	case 69:
		e = service.CreateEnumInfo(
			"MapBufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ARRAY_BUFFER",
					Value: 34962,
				},
				service.EnumEntry{
					Name:  "GL_COPY_READ_BUFFER",
					Value: 36662,
				},
				service.EnumEntry{
					Name:  "GL_COPY_WRITE_BUFFER",
					Value: 36663,
				},
				service.EnumEntry{
					Name:  "GL_ELEMENT_ARRAY_BUFFER",
					Value: 34963,
				},
				service.EnumEntry{
					Name:  "GL_PIXEL_PACK_BUFFER",
					Value: 35051,
				},
				service.EnumEntry{
					Name:  "GL_PIXEL_UNPACK_BUFFER",
					Value: 35052,
				},
				service.EnumEntry{
					Name:  "GL_TRANSFORM_FEEDBACK_BUFFER",
					Value: 35982,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BUFFER",
					Value: 35345,
				},
			},
			service.EnumInfoArray{},
		)
	case 70:
		e = service.CreateEnumInfo(
			"ImageTargetTexture_OES_EGL_image",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_2D",
					Value: 3553,
				},
			},
			service.EnumInfoArray{},
		)
	case 71:
		e = service.CreateEnumInfo(
			"ImageTargetTexture_OES_EGL_image_external",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_EXTERNAL_OES",
					Value: 36197,
				},
			},
			service.EnumInfoArray{},
		)
	case 72:
		e = service.CreateEnumInfo(
			"ImageTargetTexture",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(70),
				s.getEnumInfo(71),
			},
		)
	case 73:
		e = service.CreateEnumInfo(
			"ImageTargetRenderbufferStorage",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_OES",
					Value: 36161,
				},
			},
			service.EnumInfoArray{},
		)
	case 74:
		e = service.CreateEnumInfo(
			"ResetStatus",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_NO_ERROR",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "GL_GUILTY_CONTEXT_RESET_EXT",
					Value: 33363,
				},
				service.EnumEntry{
					Name:  "GL_INNOCENT_CONTEXT_RESET_EXT",
					Value: 33364,
				},
				service.EnumEntry{
					Name:  "GL_UNKNOWN_CONTEXT_RESET_EXT",
					Value: 33365,
				},
			},
			service.EnumInfoArray{},
		)
	case 75:
		e = service.CreateEnumInfo(
			"TextureKind",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "UNDEFINED",
					Value: 0,
				},
				service.EnumEntry{
					Name:  "TEXTURE2D",
					Value: 1,
				},
				service.EnumEntry{
					Name:  "CUBEMAP",
					Value: 2,
				},
			},
			service.EnumInfoArray{},
		)
	case 76:
		e = service.CreateEnumInfo(
			"QueryParameter_GLES_3",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_CURRENT_QUERY",
					Value: 34917,
				},
			},
			service.EnumInfoArray{},
		)
	case 77:
		e = service.CreateEnumInfo(
			"QueryParameter_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_QUERY_COUNTER_BITS_EXT",
					Value: 34916,
				},
			},
			service.EnumInfoArray{},
		)
	case 78:
		e = service.CreateEnumInfo(
			"QueryParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(76),
				s.getEnumInfo(77),
			},
		)
	case 79:
		e = service.CreateEnumInfo(
			"QueryObjectParameter_GLES_3",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_QUERY_RESULT",
					Value: 34918,
				},
				service.EnumEntry{
					Name:  "GL_QUERY_RESULT_AVAILABLE",
					Value: 34919,
				},
			},
			service.EnumInfoArray{},
		)
	case 80:
		e = service.CreateEnumInfo(
			"QueryObjectParameter_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{},
		)
	case 81:
		e = service.CreateEnumInfo(
			"QueryObjectParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(79),
				s.getEnumInfo(80),
			},
		)
	case 82:
		e = service.CreateEnumInfo(
			"QueryTarget_GLES_3",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ANY_SAMPLES_PASSED",
					Value: 35887,
				},
				service.EnumEntry{
					Name:  "GL_ANY_SAMPLES_PASSED_CONSERVATIVE",
					Value: 36202,
				},
				service.EnumEntry{
					Name:  "GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN",
					Value: 35976,
				},
			},
			service.EnumInfoArray{},
		)
	case 83:
		e = service.CreateEnumInfo(
			"QueryTarget_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TIME_ELAPSED_EXT",
					Value: 35007,
				},
				service.EnumEntry{
					Name:  "GL_TIMESTAMP_EXT",
					Value: 36392,
				},
			},
			service.EnumInfoArray{},
		)
	case 84:
		e = service.CreateEnumInfo(
			"QueryTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoArray{
				s.getEnumInfo(82),
				s.getEnumInfo(83),
			},
		)
	case 85:
		e = service.CreateEnumInfo(
			"TilePreserveMaskQCOM",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT0_QCOM",
					Value: 1,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT1_QCOM",
					Value: 2,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT2_QCOM",
					Value: 4,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT3_QCOM",
					Value: 8,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT4_QCOM",
					Value: 16,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT5_QCOM",
					Value: 32,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT6_QCOM",
					Value: 64,
				},
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT7_QCOM",
					Value: 128,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT0_QCOM",
					Value: 256,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT1_QCOM",
					Value: 512,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT2_QCOM",
					Value: 1024,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT3_QCOM",
					Value: 2048,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT4_QCOM",
					Value: 4096,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT5_QCOM",
					Value: 8192,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT6_QCOM",
					Value: 16384,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT7_QCOM",
					Value: 32768,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT0_QCOM",
					Value: 65536,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT1_QCOM",
					Value: 131072,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT2_QCOM",
					Value: 262144,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT3_QCOM",
					Value: 524288,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT4_QCOM",
					Value: 1048576,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT5_QCOM",
					Value: 2097152,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT6_QCOM",
					Value: 4194304,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT7_QCOM",
					Value: 8388608,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT0_QCOM",
					Value: 16777216,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT1_QCOM",
					Value: 33554432,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT2_QCOM",
					Value: 67108864,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT3_QCOM",
					Value: 134217728,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT4_QCOM",
					Value: 268435456,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT5_QCOM",
					Value: 536870912,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT6_QCOM",
					Value: 1073741824,
				},
				service.EnumEntry{
					Name:  "GL_MULTISAMPLE_BUFFER_BIT7_QCOM",
					Value: 2147483648,
				},
			},
			service.EnumInfoArray{},
		)
	case 86:
		e = service.CreateEnumInfo(
			"ClearMask",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COLOR_BUFFER_BIT",
					Value: 16384,
				},
				service.EnumEntry{
					Name:  "GL_DEPTH_BUFFER_BIT",
					Value: 256,
				},
				service.EnumEntry{
					Name:  "GL_STENCIL_BUFFER_BIT",
					Value: 1024,
				},
			},
			service.EnumInfoArray{},
		)
	case 87:
		e = service.CreateEnumInfo(
			"MapBufferRangeAccess",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_MAP_READ_BIT",
					Value: 1,
				},
				service.EnumEntry{
					Name:  "GL_MAP_WRITE_BIT",
					Value: 2,
				},
				service.EnumEntry{
					Name:  "GL_MAP_INVALIDATE_RANGE_BIT",
					Value: 4,
				},
				service.EnumEntry{
					Name:  "GL_MAP_INVALIDATE_BUFFER_BIT",
					Value: 8,
				},
				service.EnumEntry{
					Name:  "GL_MAP_FLUSH_EXPLICIT_BIT",
					Value: 16,
				},
				service.EnumEntry{
					Name:  "GL_MAP_UNSYNCHRONIZED_BIT",
					Value: 32,
				},
			},
			service.EnumInfoArray{},
		)
	}
	s.enums[id] = e
	return e
}
func (s schemaBuilder) getClassInfo(id int) *service.ClassInfo {
	e, f := s.classes[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateClassInfo(
			"Color",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Red",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Green",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Blue",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Alpha",
					Type: schema.Float,
				},
			},
			service.ClassInfoArray{},
		)
	case 1:
		e = service.CreateClassInfo(
			"Rect",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Width",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Height",
					Type: schema.S32,
				},
			},
			service.ClassInfoArray{},
		)
	case 2:
		e = service.CreateClassInfo(
			"Image",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Width",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Height",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Data",
					Type: schema.Memory,
				},
				&service.FieldInfo{
					Name: "Size",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Format",
					Type: s.getEnumInfo(18),
				},
			},
			service.ClassInfoArray{},
		)
	case 3:
		e = service.CreateClassInfo(
			"FramebufferAttachable",
			service.TypeKindClass,
			service.FieldInfoArray{},
			service.ClassInfoArray{},
		)
	case 4:
		e = service.CreateClassInfo(
			"Renderbuffer",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Width",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Height",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Data",
					Type: schema.Memory,
				},
				&service.FieldInfo{
					Name: "Format",
					Type: s.getEnumInfo(13),
				},
			},
			service.ClassInfoArray{
				s.getClassInfo(3),
			},
		)
	case 5:
		e = service.CreateClassInfo(
			"Texture",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Kind",
					Type: s.getEnumInfo(75),
				},
				&service.FieldInfo{
					Name: "Format",
					Type: s.getEnumInfo(18),
				},
				&service.FieldInfo{
					Name: "Texture2D",
					Type: s.getMapInfo(9),
				},
				&service.FieldInfo{
					Name: "Cubemap",
					Type: s.getMapInfo(4),
				},
				&service.FieldInfo{
					Name: "MagFilter",
					Type: s.getEnumInfo(60),
				},
				&service.FieldInfo{
					Name: "MinFilter",
					Type: s.getEnumInfo(60),
				},
				&service.FieldInfo{
					Name: "WrapS",
					Type: s.getEnumInfo(61),
				},
				&service.FieldInfo{
					Name: "WrapT",
					Type: s.getEnumInfo(61),
				},
				&service.FieldInfo{
					Name: "SwizzleR",
					Type: s.getEnumInfo(62),
				},
				&service.FieldInfo{
					Name: "SwizzleG",
					Type: s.getEnumInfo(62),
				},
				&service.FieldInfo{
					Name: "SwizzleB",
					Type: s.getEnumInfo(62),
				},
				&service.FieldInfo{
					Name: "SwizzleA",
					Type: s.getEnumInfo(62),
				},
				&service.FieldInfo{
					Name: "MaxAnisotropy",
					Type: schema.Float,
				},
			},
			service.ClassInfoArray{
				s.getClassInfo(3),
			},
		)
	case 6:
		e = service.CreateClassInfo(
			"CubemapLevel",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Faces",
					Type: s.getMapInfo(8),
				},
			},
			service.ClassInfoArray{},
		)
	case 7:
		e = service.CreateClassInfo(
			"FramebufferAttachmentInfo",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Object",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(21),
				},
				&service.FieldInfo{
					Name: "TextureLevel",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "CubeMapFace",
					Type: s.getEnumInfo(6),
				},
			},
			service.ClassInfoArray{},
		)
	case 8:
		e = service.CreateClassInfo(
			"Framebuffer",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Attachments",
					Type: s.getMapInfo(5),
				},
			},
			service.ClassInfoArray{},
		)
	case 9:
		e = service.CreateClassInfo(
			"Buffer",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Data",
					Type: schema.Memory,
				},
				&service.FieldInfo{
					Name: "Size",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Usage",
					Type: s.getEnumInfo(32),
				},
			},
			service.ClassInfoArray{},
		)
	case 10:
		e = service.CreateClassInfo(
			"Shader",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Binary",
					Type: schema.Memory,
				},
				&service.FieldInfo{
					Name: "Compiled",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "Deletable",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "InfoLog",
					Type: schema.String,
				},
				&service.FieldInfo{
					Name: "Source",
					Type: s.getArrayInfo(10),
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(33),
				},
			},
			service.ClassInfoArray{},
		)
	case 11:
		e = service.CreateClassInfo(
			"VertexAttribute",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Name",
					Type: schema.String,
				},
				&service.FieldInfo{
					Name: "VectorCount",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(46),
				},
			},
			service.ClassInfoArray{},
		)
	case 12:
		e = service.CreateClassInfo(
			"Vec2i",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.S32,
				},
			},
			service.ClassInfoArray{},
		)
	case 13:
		e = service.CreateClassInfo(
			"Vec3i",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Z",
					Type: schema.S32,
				},
			},
			service.ClassInfoArray{},
		)
	case 14:
		e = service.CreateClassInfo(
			"Vec4i",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Z",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "W",
					Type: schema.S32,
				},
			},
			service.ClassInfoArray{},
		)
	case 15:
		e = service.CreateClassInfo(
			"Vec2f",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.Float,
				},
			},
			service.ClassInfoArray{},
		)
	case 16:
		e = service.CreateClassInfo(
			"Vec3f",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Z",
					Type: schema.Float,
				},
			},
			service.ClassInfoArray{},
		)
	case 17:
		e = service.CreateClassInfo(
			"Vec4f",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "X",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Y",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Z",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "W",
					Type: schema.Float,
				},
			},
			service.ClassInfoArray{},
		)
	case 18:
		e = service.CreateClassInfo(
			"Mat2f",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Col0",
					Type: s.getClassInfo(15),
				},
				&service.FieldInfo{
					Name: "Col1",
					Type: s.getClassInfo(15),
				},
			},
			service.ClassInfoArray{},
		)
	case 19:
		e = service.CreateClassInfo(
			"Mat3f",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Col0",
					Type: s.getClassInfo(16),
				},
				&service.FieldInfo{
					Name: "Col1",
					Type: s.getClassInfo(16),
				},
				&service.FieldInfo{
					Name: "Col2",
					Type: s.getClassInfo(16),
				},
			},
			service.ClassInfoArray{},
		)
	case 20:
		e = service.CreateClassInfo(
			"Mat4f",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Col0",
					Type: s.getClassInfo(17),
				},
				&service.FieldInfo{
					Name: "Col1",
					Type: s.getClassInfo(17),
				},
				&service.FieldInfo{
					Name: "Col2",
					Type: s.getClassInfo(17),
				},
				&service.FieldInfo{
					Name: "Col3",
					Type: s.getClassInfo(17),
				},
			},
			service.ClassInfoArray{},
		)
	case 21:
		e = service.CreateClassInfo(
			"UniformValue",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "F32",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "Vec2f",
					Type: s.getClassInfo(15),
				},
				&service.FieldInfo{
					Name: "Vec3f",
					Type: s.getClassInfo(16),
				},
				&service.FieldInfo{
					Name: "Vec4f",
					Type: s.getClassInfo(17),
				},
				&service.FieldInfo{
					Name: "S32",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Vec2i",
					Type: s.getClassInfo(12),
				},
				&service.FieldInfo{
					Name: "Vec3i",
					Type: s.getClassInfo(13),
				},
				&service.FieldInfo{
					Name: "Vec4i",
					Type: s.getClassInfo(14),
				},
				&service.FieldInfo{
					Name: "Mat2f",
					Type: s.getClassInfo(18),
				},
				&service.FieldInfo{
					Name: "Mat3f",
					Type: s.getClassInfo(19),
				},
				&service.FieldInfo{
					Name: "Mat4f",
					Type: s.getClassInfo(20),
				},
			},
			service.ClassInfoArray{},
		)
	case 22:
		e = service.CreateClassInfo(
			"Uniform",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Name",
					Type: schema.String,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(47),
				},
				&service.FieldInfo{
					Name: "Value",
					Type: s.getClassInfo(21),
				},
			},
			service.ClassInfoArray{},
		)
	case 23:
		e = service.CreateClassInfo(
			"Program",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Shaders",
					Type: s.getMapInfo(15),
				},
				&service.FieldInfo{
					Name: "Linked",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "Binary",
					Type: schema.Memory,
				},
				&service.FieldInfo{
					Name: "AttributeBindings",
					Type: s.getMapInfo(0),
				},
				&service.FieldInfo{
					Name: "Attributes",
					Type: s.getMapInfo(24),
				},
				&service.FieldInfo{
					Name: "Uniforms",
					Type: s.getMapInfo(21),
				},
				&service.FieldInfo{
					Name: "InfoLog",
					Type: schema.String,
				},
			},
			service.ClassInfoArray{},
		)
	case 24:
		e = service.CreateClassInfo(
			"VertexArray",
			service.TypeKindClass,
			service.FieldInfoArray{},
			service.ClassInfoArray{},
		)
	case 25:
		e = service.CreateClassInfo(
			"VertexAttributeArray",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Enabled",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "Size",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(45),
				},
				&service.FieldInfo{
					Name: "Normalized",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "Stride",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Buffer",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Pointer",
					Type: schema.Pointer,
				},
			},
			service.ClassInfoArray{},
		)
	case 26:
		e = service.CreateClassInfo(
			"Query",
			service.TypeKindClass,
			service.FieldInfoArray{},
			service.ClassInfoArray{},
		)
	case 27:
		e = service.CreateClassInfo(
			"BlendState",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "SrcRgbBlendFactor",
					Type: s.getEnumInfo(63),
				},
				&service.FieldInfo{
					Name: "SrcAlphaBlendFactor",
					Type: s.getEnumInfo(63),
				},
				&service.FieldInfo{
					Name: "DstRgbBlendFactor",
					Type: s.getEnumInfo(63),
				},
				&service.FieldInfo{
					Name: "DstAlphaBlendFactor",
					Type: s.getEnumInfo(63),
				},
				&service.FieldInfo{
					Name: "BlendEquationRgb",
					Type: s.getEnumInfo(68),
				},
				&service.FieldInfo{
					Name: "BlendEquationAlpha",
					Type: s.getEnumInfo(68),
				},
				&service.FieldInfo{
					Name: "BlendColor",
					Type: s.getClassInfo(0),
				},
			},
			service.ClassInfoArray{},
		)
	case 28:
		e = service.CreateClassInfo(
			"RasterizerState",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "DepthMask",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "DepthTestFunction",
					Type: s.getEnumInfo(65),
				},
				&service.FieldInfo{
					Name: "DepthNear",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "DepthFar",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "ColorMaskRed",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "ColorMaskGreen",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "ColorMaskBlue",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "ColorMaskAlpha",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "StencilMask",
					Type: s.getMapInfo(20),
				},
				&service.FieldInfo{
					Name: "Viewport",
					Type: s.getClassInfo(1),
				},
				&service.FieldInfo{
					Name: "Scissor",
					Type: s.getClassInfo(1),
				},
				&service.FieldInfo{
					Name: "FrontFace",
					Type: s.getEnumInfo(67),
				},
				&service.FieldInfo{
					Name: "CullFace",
					Type: s.getEnumInfo(39),
				},
				&service.FieldInfo{
					Name: "LineWidth",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "PolygonOffsetFactor",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "PolygonOffsetUnits",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "SampleCoverageValue",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "SampleCoverageInvert",
					Type: schema.Bool,
				},
			},
			service.ClassInfoArray{},
		)
	case 29:
		e = service.CreateClassInfo(
			"ClearState",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "ClearColor",
					Type: s.getClassInfo(0),
				},
				&service.FieldInfo{
					Name: "ClearDepth",
					Type: schema.Float,
				},
				&service.FieldInfo{
					Name: "ClearStencil",
					Type: schema.S32,
				},
			},
			service.ClassInfoArray{},
		)
	case 30:
		e = service.CreateClassInfo(
			"InternalState",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "NilBuffer",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "NilTexture",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "NilRenderbuffer",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Backbuffer",
					Type: schema.U32,
				},
			},
			service.ClassInfoArray{},
		)
	case 31:
		e = service.CreateClassInfo(
			"Objects",
			service.TypeKindClass,
			service.FieldInfoArray{
				&service.FieldInfo{
					Name: "Renderbuffers",
					Type: s.getMapInfo(13),
				},
				&service.FieldInfo{
					Name: "Textures",
					Type: s.getMapInfo(19),
				},
				&service.FieldInfo{
					Name: "Framebuffers",
					Type: s.getMapInfo(7),
				},
				&service.FieldInfo{
					Name: "Buffers",
					Type: s.getMapInfo(3),
				},
				&service.FieldInfo{
					Name: "Shaders",
					Type: s.getMapInfo(16),
				},
				&service.FieldInfo{
					Name: "Programs",
					Type: s.getMapInfo(10),
				},
				&service.FieldInfo{
					Name: "VertexArrays",
					Type: s.getMapInfo(22),
				},
				&service.FieldInfo{
					Name: "Queries",
					Type: s.getMapInfo(11),
				},
			},
			service.ClassInfoArray{},
		)
	}
	s.classes[id] = e
	return e
}
