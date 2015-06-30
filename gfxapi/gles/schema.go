////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
	"android.googlesource.com/platform/tools/gpu/service"
)

func init() {
	s := schemaBuilder{
		staticArrays: make(map[int]*service.StaticArrayInfo),
		maps:         make(map[int]*service.MapInfo),
		enums:        make(map[int]*service.EnumInfo),
		structs:      make(map[int]*service.StructInfo),
		classes:      make(map[int]*service.ClassInfo),
	}
	_ = s
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 0,
		Name: "eglInitialize",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "dpy",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "major",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "minor",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglInitialize.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 1,
		Name: "eglCreateContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "config",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "share_context",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attrib_list",
				Type: schema.Pointer,
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
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 2,
		Name: "eglMakeCurrent",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "draw",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "read",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "context",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 3,
		Name: "eglSwapBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "surface",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     true,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 4,
		Name: "eglQuerySurface",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "surface",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attribute",
				Type: schema.Int,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 5,
		Name: "glXCreateContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "dpy",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "vis",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shareList",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "direct",
				Type: schema.Bool,
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
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 6,
		Name: "glXCreateNewContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "fbconfig",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shared",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "direct",
				Type: schema.Bool,
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
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 7,
		Name: "glXMakeContextCurrent",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "draw",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "read",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "ctx",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 8,
		Name: "glXMakeCurrent",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "drawable",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "ctx",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 9,
		Name: "glXSwapBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "drawable",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     true,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 10,
		Name: "glXQueryDrawable",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "display",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "draw",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attribute",
				Type: schema.Int,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 11,
		Name: "wglCreateContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "hdc",
				Type: schema.Pointer,
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
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374379(v=vs.85).aspx]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 12,
		Name: "wglCreateContextAttribsARB",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "hdc",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "hShareContext",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attribList",
				Type: schema.Pointer,
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
		DocumentationUrl: "[http://www.opengl.org/registry/specs/ARB/wgl_create_context.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 13,
		Name: "wglMakeCurrent",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "hdc",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "hglrc",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374387(v=vs.85).aspx]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 14,
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
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/dd369060(v=vs.85)]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 15,
		Name: "CGLCreateContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "pix",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "share",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "ctx",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGL_OpenGL/index.html#//apple_ref/c/func/CGLCreateContext]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 16,
		Name: "CGLSetCurrentContext",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "ctx",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 17,
		Name: "CGLGetSurface",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "ctx",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "cid",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "wid",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "sid",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 18,
		Name: "CGSGetSurfaceBounds",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "cid",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "wid",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "sid",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "bounds",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 19,
		Name: "CGLFlushDrawable",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "ctx",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     true,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 20,
		Name: "glEnableClientState",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(46),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 21,
		Name: "glDisableClientState",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(46),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 22,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary_format",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "binary",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 23,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 24,
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
				Type: s.getEnumInfo(91),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 25,
		Name: "glEndTilingQCOM",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "preserve_mask",
				Type: s.getEnumInfo(91),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 26,
		Name: "glDiscardFramebufferEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "numAttachments",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attachments",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_discard_framebuffer.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 27,
		Name: "glInsertEventMarkerEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "marker",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 28,
		Name: "glPushGroupMarkerEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "length",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "marker",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             29,
		Name:             "glPopGroupMarkerEXT",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 30,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 31,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 32,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 33,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 34,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 35,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 36,
		Name: "glGenVertexArraysOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "arrays",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 37,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 38,
		Name: "glDeleteVertexArraysOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "arrays",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 39,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 40,
		Name: "glEGLImageTargetTexture2DOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(76),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 41,
		Name: "glEGLImageTargetRenderbufferStorageOES",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(77),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 42,
		Name: "glGetGraphicsResetStatusEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(78),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_robustness.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 43,
		Name: "glBindAttribLocation",
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 44,
		Name: "glBlendFunc",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "src_factor",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dst_factor",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFunc.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 45,
		Name: "glBlendFuncSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "src_factor_rgb",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dst_factor_rgb",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "src_factor_alpha",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "dst_factor_alpha",
				Type: s.getEnumInfo(67),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFuncSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 46,
		Name: "glBlendEquation",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "equation",
				Type: s.getEnumInfo(72),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquation.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 47,
		Name: "glBlendEquationSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "rgb",
				Type: s.getEnumInfo(72),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "alpha",
				Type: s.getEnumInfo(72),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquationSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 48,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 49,
		Name: "glEnableVertexAttribArray",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 50,
		Name: "glDisableVertexAttribArray",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 51,
		Name: "glVertexAttribPointer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "size",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(49),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 52,
		Name: "glGetActiveAttrib",
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "vector_count",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 53,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "vector_count",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "type",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 54,
		Name: "glGetError",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(52),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetError.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 55,
		Name: "glGetProgramiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(56),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgram.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 56,
		Name: "glGetShaderiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(57),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderiv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 57,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 58,
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
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttribLocation.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 59,
		Name: "glPixelStorei",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(58),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 60,
		Name: "glTexParameteri",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(63),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 61,
		Name: "glTexParameterf",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(63),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 62,
		Name: "glGetTexParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 63,
		Name: "glGetTexParameterfv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(5),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(63),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 64,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 65,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 66,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 67,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 68,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 69,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 70,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 71,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 72,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 73,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 74,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 75,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 76,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 77,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 78,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 79,
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
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 80,
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
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 81,
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
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 82,
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
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 83,
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
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 84,
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
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 85,
		Name: "glVertexAttrib1f",
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
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 86,
		Name: "glVertexAttrib2f",
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 87,
		Name: "glVertexAttrib3f",
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 88,
		Name: "glVertexAttrib4f",
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 89,
		Name: "glVertexAttrib1fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 90,
		Name: "glVertexAttrib2fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 91,
		Name: "glVertexAttrib3fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 92,
		Name: "glVertexAttrib4fv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "location",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 93,
		Name: "glGetShaderPrecisionFormat",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "shader_type",
				Type: s.getEnumInfo(37),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "precision_type",
				Type: s.getEnumInfo(68),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "range",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "precision",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 94,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 95,
		Name: "glDepthFunc",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "function",
				Type: s.getEnumInfo(69),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthFunc.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 96,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 97,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 98,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 99,
		Name: "glStencilMaskSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "face",
				Type: s.getEnumInfo(43),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 100,
		Name: "glStencilFuncSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "face",
				Type: s.getEnumInfo(43),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "function",
				Type: s.getEnumInfo(69),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 101,
		Name: "glStencilOpSeparate",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "face",
				Type: s.getEnumInfo(43),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_fail",
				Type: s.getEnumInfo(70),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_pass_depth_fail",
				Type: s.getEnumInfo(70),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "stencil_pass_depth_pass",
				Type: s.getEnumInfo(70),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilOpSeparate.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 102,
		Name: "glFrontFace",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "orientation",
				Type: s.getEnumInfo(71),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFrontFace.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 103,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 104,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 105,
		Name: "glActiveTexture",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "unit",
				Type: s.getEnumInfo(35),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glActiveTexture.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 106,
		Name: "glGenTextures",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "textures",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenTextures.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 107,
		Name: "glDeleteTextures",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "textures",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteTextures.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 108,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 109,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 110,
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
				Type: s.getEnumInfo(24),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 111,
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
				Type: s.getEnumInfo(24),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 112,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 113,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 114,
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
				Type: s.getEnumInfo(22),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 115,
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
				Type: s.getEnumInfo(22),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 116,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 117,
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
				Type: s.getEnumInfo(24),
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReadPixels.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 118,
		Name: "glGenFramebuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffers",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenFramebuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 119,
		Name: "glBindFramebuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 120,
		Name: "glCheckFramebufferStatus",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(31),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 121,
		Name: "glDeleteFramebuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffers",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteFramebuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 122,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 123,
		Name: "glGenRenderbuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffers",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenRenderbuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 124,
		Name: "glBindRenderbuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(32),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 125,
		Name: "glRenderbufferStorage",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(32),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 126,
		Name: "glDeleteRenderbuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffers",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 127,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 128,
		Name: "glGetRenderbufferParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(32),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(33),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 129,
		Name: "glGenBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffers",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 130,
		Name: "glBindBuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 131,
		Name: "glBufferData",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
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
				Type: s.getEnumInfo(36),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferData.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 132,
		Name: "glBufferSubData",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 133,
		Name: "glDeleteBuffers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "buffers",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteBuffers.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 134,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 135,
		Name: "glGetBufferParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(34),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetBufferParameteriv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 136,
		Name: "glCreateShader",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "type",
				Type: s.getEnumInfo(37),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 137,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 138,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "length",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderSource.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 139,
		Name: "glShaderBinary",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shaders",
				Type: schema.Pointer,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 140,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "info",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderInfoLog.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 141,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "source",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderSource.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             142,
		Name:             "glReleaseShaderCompiler",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 143,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 144,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 145,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 146,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 147,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 148,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 149,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "shaders",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttachedShaders.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 150,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 151,
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "info",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgramInfoLog.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 152,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 153,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 154,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 155,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 156,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 157,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 158,
		Name: "glClear",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "mask",
				Type: s.getEnumInfo(92),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClear.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 159,
		Name: "glCullFace",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "mode",
				Type: s.getEnumInfo(43),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCullFace.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 160,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 161,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 162,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 163,
		Name: "glHint",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(53),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "mode",
				Type: s.getEnumInfo(54),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glHint.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 164,
		Name: "glFramebufferRenderbuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer_target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffer_attachment",
				Type: s.getEnumInfo(25),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "renderbuffer_target",
				Type: s.getEnumInfo(32),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 165,
		Name: "glFramebufferTexture2D",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer_target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "framebuffer_attachment",
				Type: s.getEnumInfo(25),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 166,
		Name: "glGetFramebufferAttachmentParameteriv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "framebuffer_target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attachment",
				Type: s.getEnumInfo(25),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(30),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 167,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 168,
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
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             169,
		Name:             "glFlush",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFlush.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             170,
		Name:             "glFinish",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFinish.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 171,
		Name: "glGetBooleanv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(42),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 172,
		Name: "glGetFloatv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(42),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 173,
		Name: "glGetIntegerv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(42),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "values",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 174,
		Name: "glGetString",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "param",
				Type: s.getEnumInfo(48),
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetString.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 175,
		Name: "glEnable",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "capability",
				Type: s.getEnumInfo(47),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnable.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 176,
		Name: "glDisable",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "capability",
				Type: s.getEnumInfo(47),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisable.xml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 177,
		Name: "glIsEnabled",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "capability",
				Type: s.getEnumInfo(47),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 178,
		Name: "glFenceSync",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "condition",
				Type: s.getEnumInfo(94),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "syncFlags",
				Type: s.getEnumInfo(96),
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glFenceSync.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 179,
		Name: "glDeleteSync",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "sync",
				Type: schema.U64,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteSync.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 180,
		Name: "glWaitSync",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "sync",
				Type: schema.U64,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "syncFlags",
				Type: s.getEnumInfo(96),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "timeout",
				Type: schema.U64,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glWaitSync.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 181,
		Name: "glClientWaitSync",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "sync",
				Type: schema.U64,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "syncFlags",
				Type: s.getEnumInfo(96),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "timeout",
				Type: schema.U64,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "result",
				Type: s.getEnumInfo(95),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glClientWaitSync.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 182,
		Name: "glMapBufferRange",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
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
				Type: s.getEnumInfo(93),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 183,
		Name: "glUnmapBuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(73),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 184,
		Name: "glInvalidateFramebuffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(29),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "attachments",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glInvalidateFramebuffer.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 185,
		Name: "glRenderbufferStorageMultisample",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(32),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 186,
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
				Type: s.getEnumInfo(92),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "filter",
				Type: s.getEnumInfo(64),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBlitFramebuffer.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 187,
		Name: "glGenQueries",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenQueries.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 188,
		Name: "glBeginQuery",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 189,
		Name: "glEndQuery",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glEndQuery.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 190,
		Name: "glDeleteQueries",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteQueries.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 191,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 192,
		Name: "glGetQueryiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(82),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryiv.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 193,
		Name: "glGetQueryObjectuiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryObjectuiv.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 194,
		Name: "glGetActiveUniformBlockName",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "uniform_block_index",
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
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "name",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniformBlockName.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 195,
		Name: "glGetActiveUniformBlockiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "uniform_block_index",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter_name",
				Type: s.getEnumInfo(89),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameters",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniformBlockiv.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 196,
		Name: "glUniformBlockBinding",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "uniform_block_index",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "uniform_block_binding",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glUniformBlockBinding.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 197,
		Name: "glGetActiveUniformsiv",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "program",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "uniform_count",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "uniform_indices",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter_name",
				Type: s.getEnumInfo(89),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameters",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniformsiv.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 198,
		Name: "glBindBufferBase",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(90),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "index",
				Type: schema.U32,
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBindBufferBase.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 199,
		Name: "glGenVertexArrays",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "arrays",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenVertexArrays.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 200,
		Name: "glBindVertexArray",
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
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBindVertexArray.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 201,
		Name: "glDeleteVertexArrays",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "arrays",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteVertexArrays.xhtml]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 202,
		Name: "glGetQueryObjecti64v",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 203,
		Name: "glGetQueryObjectui64v",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 204,
		Name: "glGenQueriesEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 205,
		Name: "glBeginQueryEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 206,
		Name: "glEndQueryEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 207,
		Name: "glDeleteQueriesEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "queries",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 208,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 209,
		Name: "glQueryCounterEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 210,
		Name: "glGetQueryivEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "target",
				Type: s.getEnumInfo(88),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(82),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 211,
		Name: "glGetQueryObjectivEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 212,
		Name: "glGetQueryObjectuivEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 213,
		Name: "glGetQueryObjecti64vEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 214,
		Name: "glGetQueryObjectui64vEXT",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "query",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "parameter",
				Type: s.getEnumInfo(85),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "value",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 215,
		Name: "architecture",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "pointer_alignment",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "pointer_size",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "integer_size",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "little_endian",
				Type: schema.Bool,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 216,
		Name: "replayCreateRenderer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "id",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 217,
		Name: "replayBindRenderer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "id",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 218,
		Name: "backbufferInfo",
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
			service.ParameterInfo{
				Name: "resetViewportScissor",
				Type: schema.Bool,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 219,
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
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 220,
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
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             221,
		Name:             "flushPostBuffer",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
}

type schemaBuilder struct {
	staticArrays map[int]*service.StaticArrayInfo
	maps         map[int]*service.MapInfo
	enums        map[int]*service.EnumInfo
	structs      map[int]*service.StructInfo
	classes      map[int]*service.ClassInfo
}

func (s schemaBuilder) getStaticArrayInfo(id int) *service.StaticArrayInfo {
	e, f := s.staticArrays[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateStaticArrayInfo("S32ː2ᵃ", service.TypeKindStaticArray, schema.S32, 2)
	case 1:
		e = service.CreateStaticArrayInfo("S32ː3ᵃ", service.TypeKindStaticArray, schema.S32, 3)
	case 2:
		e = service.CreateStaticArrayInfo("S32ː4ᵃ", service.TypeKindStaticArray, schema.S32, 4)
	case 3:
		e = service.CreateStaticArrayInfo("F32ː2ᵃ", service.TypeKindStaticArray, schema.Float, 2)
	case 4:
		e = service.CreateStaticArrayInfo("F32ː3ᵃ", service.TypeKindStaticArray, schema.Float, 3)
	case 5:
		e = service.CreateStaticArrayInfo("F32ː4ᵃ", service.TypeKindStaticArray, schema.Float, 4)
	case 6:
		e = service.CreateStaticArrayInfo("Vec2fː2ᵃ", service.TypeKindStaticArray, s.getStaticArrayInfo(3), 2)
	case 7:
		e = service.CreateStaticArrayInfo("Vec3fː3ᵃ", service.TypeKindStaticArray, s.getStaticArrayInfo(4), 3)
	case 8:
		e = service.CreateStaticArrayInfo("Vec4fː4ᵃ", service.TypeKindStaticArray, s.getStaticArrayInfo(5), 4)
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
		e = service.CreateMapInfo("AttributeLocationːVertexAttributeArrayʳᵐ", service.TypeKindMap, schema.S32, schema.Int /* TODO: Reference */)
	case 1:
		e = service.CreateMapInfo("BufferIdːBufferʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 2:
		e = service.CreateMapInfo("BufferTargetːBufferIdᵐ", service.TypeKindMap, s.getEnumInfo(73), schema.U32)
	case 3:
		e = service.CreateMapInfo("CGLContextObjːContextʳᵐ", service.TypeKindMap, schema.Pointer, schema.Int /* TODO: Reference */)
	case 4:
		e = service.CreateMapInfo("Capabilityːboolᵐ", service.TypeKindMap, s.getEnumInfo(47), schema.Bool)
	case 5:
		e = service.CreateMapInfo("CubeMapImageTargetːImageᵐ", service.TypeKindMap, s.getEnumInfo(6), s.getClassInfo(2))
	case 6:
		e = service.CreateMapInfo("EGLContextːContextʳᵐ", service.TypeKindMap, schema.Pointer, schema.Int /* TODO: Reference */)
	case 7:
		e = service.CreateMapInfo("FaceModeːu32ᵐ", service.TypeKindMap, s.getEnumInfo(43), schema.U32)
	case 8:
		e = service.CreateMapInfo("FramebufferAttachmentːFramebufferAttachmentInfoᵐ", service.TypeKindMap, s.getEnumInfo(25), s.getClassInfo(6))
	case 9:
		e = service.CreateMapInfo("FramebufferIdːFramebufferʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 10:
		e = service.CreateMapInfo("FramebufferTargetːFramebufferIdᵐ", service.TypeKindMap, s.getEnumInfo(29), schema.U32)
	case 11:
		e = service.CreateMapInfo("GLXContextːContextʳᵐ", service.TypeKindMap, schema.Pointer, schema.Int /* TODO: Reference */)
	case 12:
		e = service.CreateMapInfo("HGLRCːContextʳᵐ", service.TypeKindMap, schema.Pointer, schema.Int /* TODO: Reference */)
	case 13:
		e = service.CreateMapInfo("PixelStoreParameterːs32ᵐ", service.TypeKindMap, s.getEnumInfo(58), schema.S32)
	case 14:
		e = service.CreateMapInfo("ProgramIdːProgramʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 15:
		e = service.CreateMapInfo("QueryIdːQueryʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 16:
		e = service.CreateMapInfo("RenderbufferIdːRenderbufferʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 17:
		e = service.CreateMapInfo("RenderbufferTargetːRenderbufferIdᵐ", service.TypeKindMap, s.getEnumInfo(32), schema.U32)
	case 18:
		e = service.CreateMapInfo("S32ːCubemapLevelᵐ", service.TypeKindMap, schema.S32, s.getClassInfo(5))
	case 19:
		e = service.CreateMapInfo("S32ːImageᵐ", service.TypeKindMap, schema.S32, s.getClassInfo(2))
	case 20:
		e = service.CreateMapInfo("S32ːVertexAttributeᵐ", service.TypeKindMap, schema.S32, s.getClassInfo(10))
	case 21:
		e = service.CreateMapInfo("ShaderIdːShaderʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 22:
		e = service.CreateMapInfo("ShaderTypeːShaderIdᵐ", service.TypeKindMap, s.getEnumInfo(37), schema.U32)
	case 23:
		e = service.CreateMapInfo("StringːAttributeLocationᵐ", service.TypeKindMap, schema.String, schema.S32)
	case 24:
		e = service.CreateMapInfo("TextureIdːTextureʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 25:
		e = service.CreateMapInfo("TextureTargetːTextureIdᵐ", service.TypeKindMap, s.getEnumInfo(5), schema.U32)
	case 26:
		e = service.CreateMapInfo("TextureUnitːTextureTargetːTextureIdᵐᵐ", service.TypeKindMap, s.getEnumInfo(35), s.getMapInfo(25))
	case 27:
		e = service.CreateMapInfo("ThreadIDːContextʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
	case 28:
		e = service.CreateMapInfo("UniformLocationːUniformᵐ", service.TypeKindMap, schema.S32, s.getClassInfo(11))
	case 29:
		e = service.CreateMapInfo("VertexArrayIdːVertexArrayʳᵐ", service.TypeKindMap, schema.U32, schema.Int /* TODO: Reference */)
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
			service.EnumInfoPtrArray{},
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
			service.EnumInfoPtrArray{},
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
			service.EnumInfoPtrArray{},
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
			service.EnumInfoPtrArray{},
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
			service.EnumInfoPtrArray{},
		)
	case 5:
		e = service.CreateEnumInfo(
			"TextureTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
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
			service.EnumInfoPtrArray{},
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
			service.EnumInfoPtrArray{},
		)
	case 8:
		e = service.CreateEnumInfo(
			"TextureImageTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
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
			service.EnumInfoPtrArray{},
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
			service.EnumInfoPtrArray{
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
			service.EnumInfoPtrArray{},
		)
	case 12:
		e = service.CreateEnumInfo(
			"TexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
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
			service.EnumInfoPtrArray{},
		)
	case 14:
		e = service.CreateEnumInfo(
			"Type_ARB_half_float_vertex",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_HALF_FLOAT_ARB",
					Value: 5131,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 15:
		e = service.CreateEnumInfo(
			"Type_OES_vertex_half_float",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_HALF_FLOAT_OES",
					Value: 36193,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 16:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ETC1_RGB8_OES",
					Value: 36196,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 17:
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
			service.EnumInfoPtrArray{},
		)
	case 18:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_EXT_texture_compression_dxt1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGB_S3TC_DXT1_EXT",
					Value: 33776,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_S3TC_DXT1_EXT",
					Value: 33777,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 19:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_EXT_texture_compression_s3tc",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_S3TC_DXT3_EXT",
					Value: 33778,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_S3TC_DXT5_EXT",
					Value: 33779,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 20:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_KHR_texture_compression_astc_ldr",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_4x4_KHR",
					Value: 37808,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_5x4_KHR",
					Value: 37809,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_5x5_KHR",
					Value: 37810,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_6x5_KHR",
					Value: 37811,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_6x6_KHR",
					Value: 37812,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_8x5_KHR",
					Value: 37813,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_8x6_KHR",
					Value: 37814,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_8x8_KHR",
					Value: 37815,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_10x5_KHR",
					Value: 37816,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_10x6_KHR",
					Value: 37817,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_10x8_KHR",
					Value: 37818,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_10x10_KHR",
					Value: 37819,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_12x10_KHR",
					Value: 37820,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_RGBA_ASTC_12x12_KHR",
					Value: 37821,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR",
					Value: 37840,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR",
					Value: 37841,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR",
					Value: 37842,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR",
					Value: 37843,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR",
					Value: 37844,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR",
					Value: 37845,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR",
					Value: 37846,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR",
					Value: 37847,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR",
					Value: 37848,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR",
					Value: 37849,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR",
					Value: 37850,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR",
					Value: 37851,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR",
					Value: 37852,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR",
					Value: 37853,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 21:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat_NV_texture_compression_latc",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_COMPRESSED_LUMINANCE_LATC1_NV",
					Value: 35952,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_NV",
					Value: 35953,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_NV",
					Value: 35954,
				},
				service.EnumEntry{
					Name:  "GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_NV",
					Value: 35955,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 22:
		e = service.CreateEnumInfo(
			"CompressedTexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(16),
				s.getEnumInfo(17),
				s.getEnumInfo(18),
				s.getEnumInfo(19),
				s.getEnumInfo(20),
				s.getEnumInfo(21),
			},
		)
	case 23:
		e = service.CreateEnumInfo(
			"ImageTexelFormat",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(12),
				s.getEnumInfo(22),
			},
		)
	case 24:
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
			service.EnumInfoPtrArray{},
		)
	case 25:
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
			service.EnumInfoPtrArray{},
		)
	case 26:
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
			service.EnumInfoPtrArray{},
		)
	case 27:
		e = service.CreateEnumInfo(
			"FramebufferTarget_GLES_2_0",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_FRAMEBUFFER",
					Value: 36160,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 28:
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
			service.EnumInfoPtrArray{},
		)
	case 29:
		e = service.CreateEnumInfo(
			"FramebufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(27),
				s.getEnumInfo(28),
			},
		)
	case 30:
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
			service.EnumInfoPtrArray{},
		)
	case 31:
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
			service.EnumInfoPtrArray{},
		)
	case 32:
		e = service.CreateEnumInfo(
			"RenderbufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER",
					Value: 36161,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 33:
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
			service.EnumInfoPtrArray{},
		)
	case 34:
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
			service.EnumInfoPtrArray{},
		)
	case 35:
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
			service.EnumInfoPtrArray{},
		)
	case 36:
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
			service.EnumInfoPtrArray{},
		)
	case 37:
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
			service.EnumInfoPtrArray{},
		)
	case 38:
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
			service.EnumInfoPtrArray{},
		)
	case 39:
		e = service.CreateEnumInfo(
			"StateVariable_GLES_3_1",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_READ_FRAMEBUFFER_BINDING",
					Value: 36010,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 40:
		e = service.CreateEnumInfo(
			"StateVariable_EXT_texture_filter_anisotropic",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT",
					Value: 34047,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 41:
		e = service.CreateEnumInfo(
			"StateVariable_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_GPU_DISJOINT_EXT",
					Value: 36795,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 42:
		e = service.CreateEnumInfo(
			"StateVariable",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(38),
				s.getEnumInfo(39),
				s.getEnumInfo(40),
				s.getEnumInfo(41),
			},
		)
	case 43:
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
			service.EnumInfoPtrArray{},
		)
	case 44:
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
			service.EnumInfoPtrArray{},
		)
	case 45:
		e = service.CreateEnumInfo(
			"ArrayType_OES_point_size_array",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_POINT_SIZE_ARRAY_OES",
					Value: 35740,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 46:
		e = service.CreateEnumInfo(
			"ArrayType",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(44),
				s.getEnumInfo(45),
			},
		)
	case 47:
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
			service.EnumInfoPtrArray{
				s.getEnumInfo(46),
			},
		)
	case 48:
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
			service.EnumInfoPtrArray{},
		)
	case 49:
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
			service.EnumInfoPtrArray{
				s.getEnumInfo(15),
				s.getEnumInfo(14),
			},
		)
	case 50:
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
			service.EnumInfoPtrArray{},
		)
	case 51:
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
			service.EnumInfoPtrArray{},
		)
	case 52:
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
			service.EnumInfoPtrArray{},
		)
	case 53:
		e = service.CreateEnumInfo(
			"HintTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_GENERATE_MIPMAP_HINT",
					Value: 33170,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 54:
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
			service.EnumInfoPtrArray{},
		)
	case 55:
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
			service.EnumInfoPtrArray{},
		)
	case 56:
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
			service.EnumInfoPtrArray{},
		)
	case 57:
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
			service.EnumInfoPtrArray{},
		)
	case 58:
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
			service.EnumInfoPtrArray{},
		)
	case 59:
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
			service.EnumInfoPtrArray{},
		)
	case 60:
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
			service.EnumInfoPtrArray{},
		)
	case 61:
		e = service.CreateEnumInfo(
			"TextureParameter_EXT_texture_filter_anisotropic",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_MAX_ANISOTROPY_EXT",
					Value: 34046,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 62:
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
			service.EnumInfoPtrArray{},
		)
	case 63:
		e = service.CreateEnumInfo(
			"TextureParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(59),
				s.getEnumInfo(60),
				s.getEnumInfo(62),
				s.getEnumInfo(61),
			},
		)
	case 64:
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
			service.EnumInfoPtrArray{},
		)
	case 65:
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
			service.EnumInfoPtrArray{},
		)
	case 66:
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
			service.EnumInfoPtrArray{},
		)
	case 67:
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
			service.EnumInfoPtrArray{},
		)
	case 68:
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
			service.EnumInfoPtrArray{},
		)
	case 69:
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
			service.EnumInfoPtrArray{},
		)
	case 70:
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
			service.EnumInfoPtrArray{},
		)
	case 71:
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
			service.EnumInfoPtrArray{},
		)
	case 72:
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
			service.EnumInfoPtrArray{},
		)
	case 73:
		e = service.CreateEnumInfo(
			"BufferTarget",
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
			service.EnumInfoPtrArray{},
		)
	case 74:
		e = service.CreateEnumInfo(
			"ImageTargetTexture_OES_EGL_image",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_2D",
					Value: 3553,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 75:
		e = service.CreateEnumInfo(
			"ImageTargetTexture_OES_EGL_image_external",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TEXTURE_EXTERNAL_OES",
					Value: 36197,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 76:
		e = service.CreateEnumInfo(
			"ImageTargetTexture",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(74),
				s.getEnumInfo(75),
			},
		)
	case 77:
		e = service.CreateEnumInfo(
			"ImageTargetRenderbufferStorage",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_RENDERBUFFER_OES",
					Value: 36161,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 78:
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
			service.EnumInfoPtrArray{},
		)
	case 79:
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
			service.EnumInfoPtrArray{},
		)
	case 80:
		e = service.CreateEnumInfo(
			"QueryParameter_GLES_3",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_CURRENT_QUERY",
					Value: 34917,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 81:
		e = service.CreateEnumInfo(
			"QueryParameter_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_QUERY_COUNTER_BITS_EXT",
					Value: 34916,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 82:
		e = service.CreateEnumInfo(
			"QueryParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(80),
				s.getEnumInfo(81),
			},
		)
	case 83:
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
			service.EnumInfoPtrArray{},
		)
	case 84:
		e = service.CreateEnumInfo(
			"QueryObjectParameter_EXT_disjoint_timer_query",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{},
		)
	case 85:
		e = service.CreateEnumInfo(
			"QueryObjectParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(83),
				s.getEnumInfo(84),
			},
		)
	case 86:
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
			service.EnumInfoPtrArray{},
		)
	case 87:
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
			service.EnumInfoPtrArray{},
		)
	case 88:
		e = service.CreateEnumInfo(
			"QueryTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{},
			service.EnumInfoPtrArray{
				s.getEnumInfo(86),
				s.getEnumInfo(87),
			},
		)
	case 89:
		e = service.CreateEnumInfo(
			"UniformBlockParameter",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_BINDING",
					Value: 35391,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_DATA_SIZE",
					Value: 35392,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_NAME_LENGTH",
					Value: 35393,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS",
					Value: 35394,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES",
					Value: 35395,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER",
					Value: 35396,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER",
					Value: 35397,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER",
					Value: 35398,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 90:
		e = service.CreateEnumInfo(
			"IndexedBufferTarget",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_TRANSFORM_FEEDBACK_BUFFER",
					Value: 35982,
				},
				service.EnumEntry{
					Name:  "GL_UNIFORM_BUFFER",
					Value: 35345,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 91:
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
			service.EnumInfoPtrArray{},
		)
	case 92:
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
			service.EnumInfoPtrArray{},
		)
	case 93:
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
			service.EnumInfoPtrArray{},
		)
	case 94:
		e = service.CreateEnumInfo(
			"SyncCondition",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_SYNC_GPU_COMMANDS_COMPLETE",
					Value: 37143,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 95:
		e = service.CreateEnumInfo(
			"ClientWaitSyncSignal",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_ALREADY_SIGNALED",
					Value: 37146,
				},
				service.EnumEntry{
					Name:  "GL_TIMEOUT_EXPIRED",
					Value: 37147,
				},
				service.EnumEntry{
					Name:  "GL_CONDITION_SATISFIED",
					Value: 37148,
				},
				service.EnumEntry{
					Name:  "GL_WAIT_FAILED",
					Value: 37149,
				},
			},
			service.EnumInfoPtrArray{},
		)
	case 96:
		e = service.CreateEnumInfo(
			"SyncFlags",
			service.TypeKindEnum,
			[]service.EnumEntry{
				service.EnumEntry{
					Name:  "GL_SYNC_FLUSH_COMMANDS_BIT",
					Value: 1,
				},
			},
			service.EnumInfoPtrArray{},
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
			service.FieldInfoPtrArray{
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
			service.ClassInfoPtrArray{},
		)
	case 1:
		e = service.CreateClassInfo(
			"Rect",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
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
			service.ClassInfoPtrArray{},
		)
	case 2:
		e = service.CreateClassInfo(
			"Image",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
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
					Type: schema.Int, /* TODO: Slice */
				},
				&service.FieldInfo{
					Name: "Size",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Format",
					Type: s.getEnumInfo(23),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 3:
		e = service.CreateClassInfo(
			"Renderbuffer",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
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
					Type: schema.Int, /* TODO: Slice */
				},
				&service.FieldInfo{
					Name: "Format",
					Type: s.getEnumInfo(13),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 4:
		e = service.CreateClassInfo(
			"Texture",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Kind",
					Type: s.getEnumInfo(79),
				},
				&service.FieldInfo{
					Name: "Format",
					Type: s.getEnumInfo(23),
				},
				&service.FieldInfo{
					Name: "Texture2D",
					Type: s.getMapInfo(19),
				},
				&service.FieldInfo{
					Name: "Cubemap",
					Type: s.getMapInfo(18),
				},
				&service.FieldInfo{
					Name: "MagFilter",
					Type: s.getEnumInfo(64),
				},
				&service.FieldInfo{
					Name: "MinFilter",
					Type: s.getEnumInfo(64),
				},
				&service.FieldInfo{
					Name: "WrapS",
					Type: s.getEnumInfo(65),
				},
				&service.FieldInfo{
					Name: "WrapT",
					Type: s.getEnumInfo(65),
				},
				&service.FieldInfo{
					Name: "SwizzleR",
					Type: s.getEnumInfo(66),
				},
				&service.FieldInfo{
					Name: "SwizzleG",
					Type: s.getEnumInfo(66),
				},
				&service.FieldInfo{
					Name: "SwizzleB",
					Type: s.getEnumInfo(66),
				},
				&service.FieldInfo{
					Name: "SwizzleA",
					Type: s.getEnumInfo(66),
				},
				&service.FieldInfo{
					Name: "MaxAnisotropy",
					Type: schema.Float,
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 5:
		e = service.CreateClassInfo(
			"CubemapLevel",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Faces",
					Type: s.getMapInfo(5),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 6:
		e = service.CreateClassInfo(
			"FramebufferAttachmentInfo",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Object",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(26),
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
			service.ClassInfoPtrArray{},
		)
	case 7:
		e = service.CreateClassInfo(
			"Framebuffer",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Attachments",
					Type: s.getMapInfo(8),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 8:
		e = service.CreateClassInfo(
			"Buffer",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Data",
					Type: schema.Int, /* TODO: Slice */
				},
				&service.FieldInfo{
					Name: "Size",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Usage",
					Type: s.getEnumInfo(36),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 9:
		e = service.CreateClassInfo(
			"Shader",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Binary",
					Type: schema.Int, /* TODO: Slice */
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
					Type: schema.Int, /* TODO: Slice */
				},
				&service.FieldInfo{
					Name: "Source",
					Type: schema.String,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(37),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 10:
		e = service.CreateClassInfo(
			"VertexAttribute",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Name",
					Type: schema.Int, /* TODO: Slice */
				},
				&service.FieldInfo{
					Name: "VectorCount",
					Type: schema.S32,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(50),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 11:
		e = service.CreateClassInfo(
			"Uniform",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Name",
					Type: schema.String,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(51),
				},
				&service.FieldInfo{
					Name: "Value",
					Type: schema.Int, /* TODO: Slice */
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 12:
		e = service.CreateClassInfo(
			"Program",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Shaders",
					Type: s.getMapInfo(22),
				},
				&service.FieldInfo{
					Name: "Linked",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "Binary",
					Type: schema.Int, /* TODO: Slice */
				},
				&service.FieldInfo{
					Name: "AttributeBindings",
					Type: s.getMapInfo(23),
				},
				&service.FieldInfo{
					Name: "Attributes",
					Type: s.getMapInfo(20),
				},
				&service.FieldInfo{
					Name: "Uniforms",
					Type: s.getMapInfo(28),
				},
				&service.FieldInfo{
					Name: "InfoLog",
					Type: schema.Int, /* TODO: Slice */
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 13:
		e = service.CreateClassInfo(
			"VertexArray",
			service.TypeKindClass,
			service.FieldInfoPtrArray{},
			service.ClassInfoPtrArray{},
		)
	case 14:
		e = service.CreateClassInfo(
			"VertexAttributeArray",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Enabled",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "Size",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Type",
					Type: s.getEnumInfo(49),
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
			service.ClassInfoPtrArray{},
		)
	case 15:
		e = service.CreateClassInfo(
			"Query",
			service.TypeKindClass,
			service.FieldInfoPtrArray{},
			service.ClassInfoPtrArray{},
		)
	case 16:
		e = service.CreateClassInfo(
			"BlendState",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "SrcRgbBlendFactor",
					Type: s.getEnumInfo(67),
				},
				&service.FieldInfo{
					Name: "SrcAlphaBlendFactor",
					Type: s.getEnumInfo(67),
				},
				&service.FieldInfo{
					Name: "DstRgbBlendFactor",
					Type: s.getEnumInfo(67),
				},
				&service.FieldInfo{
					Name: "DstAlphaBlendFactor",
					Type: s.getEnumInfo(67),
				},
				&service.FieldInfo{
					Name: "BlendEquationRgb",
					Type: s.getEnumInfo(72),
				},
				&service.FieldInfo{
					Name: "BlendEquationAlpha",
					Type: s.getEnumInfo(72),
				},
				&service.FieldInfo{
					Name: "BlendColor",
					Type: s.getClassInfo(0),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 17:
		e = service.CreateClassInfo(
			"RasterizerState",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "DepthMask",
					Type: schema.Bool,
				},
				&service.FieldInfo{
					Name: "DepthTestFunction",
					Type: s.getEnumInfo(69),
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
					Type: s.getMapInfo(7),
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
					Type: s.getEnumInfo(71),
				},
				&service.FieldInfo{
					Name: "CullFace",
					Type: s.getEnumInfo(43),
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
			service.ClassInfoPtrArray{},
		)
	case 18:
		e = service.CreateClassInfo(
			"ClearState",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
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
			service.ClassInfoPtrArray{},
		)
	case 19:
		e = service.CreateClassInfo(
			"Objects",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Renderbuffers",
					Type: s.getMapInfo(16),
				},
				&service.FieldInfo{
					Name: "Textures",
					Type: s.getMapInfo(24),
				},
				&service.FieldInfo{
					Name: "Framebuffers",
					Type: s.getMapInfo(9),
				},
				&service.FieldInfo{
					Name: "Buffers",
					Type: s.getMapInfo(1),
				},
				&service.FieldInfo{
					Name: "Shaders",
					Type: s.getMapInfo(21),
				},
				&service.FieldInfo{
					Name: "Programs",
					Type: s.getMapInfo(14),
				},
				&service.FieldInfo{
					Name: "VertexArrays",
					Type: s.getMapInfo(29),
				},
				&service.FieldInfo{
					Name: "Queries",
					Type: s.getMapInfo(15),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 20:
		e = service.CreateClassInfo(
			"Context",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "Identifier",
					Type: schema.U32,
				},
				&service.FieldInfo{
					Name: "Blending",
					Type: s.getClassInfo(16),
				},
				&service.FieldInfo{
					Name: "Rasterizing",
					Type: s.getClassInfo(17),
				},
				&service.FieldInfo{
					Name: "Clearing",
					Type: s.getClassInfo(18),
				},
				&service.FieldInfo{
					Name: "BoundFramebuffers",
					Type: s.getMapInfo(10),
				},
				&service.FieldInfo{
					Name: "BoundRenderbuffers",
					Type: s.getMapInfo(17),
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
					Type: s.getMapInfo(0),
				},
				&service.FieldInfo{
					Name: "TextureUnits",
					Type: s.getMapInfo(26),
				},
				&service.FieldInfo{
					Name: "ActiveTextureUnit",
					Type: s.getEnumInfo(35),
				},
				&service.FieldInfo{
					Name: "Capabilities",
					Type: s.getMapInfo(4),
				},
				&service.FieldInfo{
					Name: "GenerateMipmapHint",
					Type: s.getEnumInfo(54),
				},
				&service.FieldInfo{
					Name: "PixelStorage",
					Type: s.getMapInfo(13),
				},
				&service.FieldInfo{
					Name: "Instances",
					Type: s.getClassInfo(19),
				},
			},
			service.ClassInfoPtrArray{},
		)
	}
	s.classes[id] = e
	return e
}
