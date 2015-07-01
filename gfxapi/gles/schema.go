////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

func init() {
	sc_EglInitialize := schema.Of((*EglInitialize)(nil).Class())
	sc_EglInitialize.Metadata = append(sc_EglInitialize.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglInitialize.xhtml]",
	})

	sc_EglCreateContext := schema.Of((*EglCreateContext)(nil).Class())
	sc_EglCreateContext.Metadata = append(sc_EglCreateContext.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml]",
	})

	sc_EglMakeCurrent := schema.Of((*EglMakeCurrent)(nil).Class())
	sc_EglMakeCurrent.Metadata = append(sc_EglMakeCurrent.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml]",
	})

	sc_EglSwapBuffers := schema.Of((*EglSwapBuffers)(nil).Class())
	sc_EglSwapBuffers.Metadata = append(sc_EglSwapBuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0 | atom.EndOfFrame,
		DocumentationUrl: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml]",
	})

	sc_EglQuerySurface := schema.Of((*EglQuerySurface)(nil).Class())
	sc_EglQuerySurface.Metadata = append(sc_EglQuerySurface.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlXCreateContext := schema.Of((*GlXCreateContext)(nil).Class())
	sc_GlXCreateContext.Metadata = append(sc_GlXCreateContext.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlXCreateNewContext := schema.Of((*GlXCreateNewContext)(nil).Class())
	sc_GlXCreateNewContext.Metadata = append(sc_GlXCreateNewContext.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlXMakeContextCurrent := schema.Of((*GlXMakeContextCurrent)(nil).Class())
	sc_GlXMakeContextCurrent.Metadata = append(sc_GlXMakeContextCurrent.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlXMakeCurrent := schema.Of((*GlXMakeCurrent)(nil).Class())
	sc_GlXMakeCurrent.Metadata = append(sc_GlXMakeCurrent.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlXSwapBuffers := schema.Of((*GlXSwapBuffers)(nil).Class())
	sc_GlXSwapBuffers.Metadata = append(sc_GlXSwapBuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0 | atom.EndOfFrame,
		DocumentationUrl: "[]",
	})

	sc_GlXQueryDrawable := schema.Of((*GlXQueryDrawable)(nil).Class())
	sc_GlXQueryDrawable.Metadata = append(sc_GlXQueryDrawable.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_WglCreateContext := schema.Of((*WglCreateContext)(nil).Class())
	sc_WglCreateContext.Metadata = append(sc_WglCreateContext.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374379(v=vs.85).aspx]",
	})

	sc_WglCreateContextAttribsARB := schema.Of((*WglCreateContextAttribsARB)(nil).Class())
	sc_WglCreateContextAttribsARB.Metadata = append(sc_WglCreateContextAttribsARB.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.opengl.org/registry/specs/ARB/wgl_create_context.txt]",
	})

	sc_WglMakeCurrent := schema.Of((*WglMakeCurrent)(nil).Class())
	sc_WglMakeCurrent.Metadata = append(sc_WglMakeCurrent.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374387(v=vs.85).aspx]",
	})

	sc_WglSwapBuffers := schema.Of((*WglSwapBuffers)(nil).Class())
	sc_WglSwapBuffers.Metadata = append(sc_WglSwapBuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0 | atom.EndOfFrame,
		DocumentationUrl: "[http://msdn.microsoft.com/en-us/library/dd369060(v=vs.85)]",
	})

	sc_CGLCreateContext := schema.Of((*CGLCreateContext)(nil).Class())
	sc_CGLCreateContext.Metadata = append(sc_CGLCreateContext.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGL_OpenGL/index.html#//apple_ref/c/func/CGLCreateContext]",
	})

	sc_CGLSetCurrentContext := schema.Of((*CGLSetCurrentContext)(nil).Class())
	sc_CGLSetCurrentContext.Metadata = append(sc_CGLSetCurrentContext.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CGLGetSurface := schema.Of((*CGLGetSurface)(nil).Class())
	sc_CGLGetSurface.Metadata = append(sc_CGLGetSurface.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CGSGetSurfaceBounds := schema.Of((*CGSGetSurfaceBounds)(nil).Class())
	sc_CGSGetSurfaceBounds.Metadata = append(sc_CGSGetSurfaceBounds.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CGLFlushDrawable := schema.Of((*CGLFlushDrawable)(nil).Class())
	sc_CGLFlushDrawable.Metadata = append(sc_CGLFlushDrawable.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0 | atom.EndOfFrame,
		DocumentationUrl: "[]",
	})

	sc_GlEnableClientState := schema.Of((*GlEnableClientState)(nil).Class())
	sc_GlEnableClientState.Metadata = append(sc_GlEnableClientState.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
	})

	sc_GlDisableClientState := schema.Of((*GlDisableClientState)(nil).Class())
	sc_GlDisableClientState.Metadata = append(sc_GlDisableClientState.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
	})

	sc_GlGetProgramBinaryOES := schema.Of((*GlGetProgramBinaryOES)(nil).Class())
	sc_GlGetProgramBinaryOES.Metadata = append(sc_GlGetProgramBinaryOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
	})

	sc_GlProgramBinaryOES := schema.Of((*GlProgramBinaryOES)(nil).Class())
	sc_GlProgramBinaryOES.Metadata = append(sc_GlProgramBinaryOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
	})

	sc_GlStartTilingQCOM := schema.Of((*GlStartTilingQCOM)(nil).Class())
	sc_GlStartTilingQCOM.Metadata = append(sc_GlStartTilingQCOM.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
	})

	sc_GlEndTilingQCOM := schema.Of((*GlEndTilingQCOM)(nil).Class())
	sc_GlEndTilingQCOM.Metadata = append(sc_GlEndTilingQCOM.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
	})

	sc_GlDiscardFramebufferEXT := schema.Of((*GlDiscardFramebufferEXT)(nil).Class())
	sc_GlDiscardFramebufferEXT.Metadata = append(sc_GlDiscardFramebufferEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_discard_framebuffer.txt]",
	})

	sc_GlInsertEventMarkerEXT := schema.Of((*GlInsertEventMarkerEXT)(nil).Class())
	sc_GlInsertEventMarkerEXT.Metadata = append(sc_GlInsertEventMarkerEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})

	sc_GlPushGroupMarkerEXT := schema.Of((*GlPushGroupMarkerEXT)(nil).Class())
	sc_GlPushGroupMarkerEXT.Metadata = append(sc_GlPushGroupMarkerEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})

	sc_GlPopGroupMarkerEXT := schema.Of((*GlPopGroupMarkerEXT)(nil).Class())
	sc_GlPopGroupMarkerEXT.Metadata = append(sc_GlPopGroupMarkerEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
	})

	sc_GlTexStorage1DEXT := schema.Of((*GlTexStorage1DEXT)(nil).Class())
	sc_GlTexStorage1DEXT.Metadata = append(sc_GlTexStorage1DEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})

	sc_GlTexStorage2DEXT := schema.Of((*GlTexStorage2DEXT)(nil).Class())
	sc_GlTexStorage2DEXT.Metadata = append(sc_GlTexStorage2DEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})

	sc_GlTexStorage3DEXT := schema.Of((*GlTexStorage3DEXT)(nil).Class())
	sc_GlTexStorage3DEXT.Metadata = append(sc_GlTexStorage3DEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})

	sc_GlTextureStorage1DEXT := schema.Of((*GlTextureStorage1DEXT)(nil).Class())
	sc_GlTextureStorage1DEXT.Metadata = append(sc_GlTextureStorage1DEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})

	sc_GlTextureStorage2DEXT := schema.Of((*GlTextureStorage2DEXT)(nil).Class())
	sc_GlTextureStorage2DEXT.Metadata = append(sc_GlTextureStorage2DEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})

	sc_GlTextureStorage3DEXT := schema.Of((*GlTextureStorage3DEXT)(nil).Class())
	sc_GlTextureStorage3DEXT.Metadata = append(sc_GlTextureStorage3DEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
	})

	sc_GlGenVertexArraysOES := schema.Of((*GlGenVertexArraysOES)(nil).Class())
	sc_GlGenVertexArraysOES.Metadata = append(sc_GlGenVertexArraysOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})

	sc_GlBindVertexArrayOES := schema.Of((*GlBindVertexArrayOES)(nil).Class())
	sc_GlBindVertexArrayOES.Metadata = append(sc_GlBindVertexArrayOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})

	sc_GlDeleteVertexArraysOES := schema.Of((*GlDeleteVertexArraysOES)(nil).Class())
	sc_GlDeleteVertexArraysOES.Metadata = append(sc_GlDeleteVertexArraysOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})

	sc_GlIsVertexArrayOES := schema.Of((*GlIsVertexArrayOES)(nil).Class())
	sc_GlIsVertexArrayOES.Metadata = append(sc_GlIsVertexArrayOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
	})

	sc_GlEGLImageTargetTexture2DOES := schema.Of((*GlEGLImageTargetTexture2DOES)(nil).Class())
	sc_GlEGLImageTargetTexture2DOES.Metadata = append(sc_GlEGLImageTargetTexture2DOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
	})

	sc_GlEGLImageTargetRenderbufferStorageOES := schema.Of((*GlEGLImageTargetRenderbufferStorageOES)(nil).Class())
	sc_GlEGLImageTargetRenderbufferStorageOES.Metadata = append(sc_GlEGLImageTargetRenderbufferStorageOES.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
	})

	sc_GlGetGraphicsResetStatusEXT := schema.Of((*GlGetGraphicsResetStatusEXT)(nil).Class())
	sc_GlGetGraphicsResetStatusEXT.Metadata = append(sc_GlGetGraphicsResetStatusEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_robustness.txt]",
	})

	sc_GlBindAttribLocation := schema.Of((*GlBindAttribLocation)(nil).Class())
	sc_GlBindAttribLocation.Metadata = append(sc_GlBindAttribLocation.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindAttribLocation.xml]",
	})

	sc_GlBlendFunc := schema.Of((*GlBlendFunc)(nil).Class())
	sc_GlBlendFunc.Metadata = append(sc_GlBlendFunc.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFunc.xml]",
	})

	sc_GlBlendFuncSeparate := schema.Of((*GlBlendFuncSeparate)(nil).Class())
	sc_GlBlendFuncSeparate.Metadata = append(sc_GlBlendFuncSeparate.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFuncSeparate.xml]",
	})

	sc_GlBlendEquation := schema.Of((*GlBlendEquation)(nil).Class())
	sc_GlBlendEquation.Metadata = append(sc_GlBlendEquation.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquation.xml]",
	})

	sc_GlBlendEquationSeparate := schema.Of((*GlBlendEquationSeparate)(nil).Class())
	sc_GlBlendEquationSeparate.Metadata = append(sc_GlBlendEquationSeparate.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquationSeparate.xml]",
	})

	sc_GlBlendColor := schema.Of((*GlBlendColor)(nil).Class())
	sc_GlBlendColor.Metadata = append(sc_GlBlendColor.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendColor.xml]",
	})

	sc_GlEnableVertexAttribArray := schema.Of((*GlEnableVertexAttribArray)(nil).Class())
	sc_GlEnableVertexAttribArray.Metadata = append(sc_GlEnableVertexAttribArray.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml]",
	})

	sc_GlDisableVertexAttribArray := schema.Of((*GlDisableVertexAttribArray)(nil).Class())
	sc_GlDisableVertexAttribArray.Metadata = append(sc_GlDisableVertexAttribArray.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml]",
	})

	sc_GlVertexAttribPointer := schema.Of((*GlVertexAttribPointer)(nil).Class())
	sc_GlVertexAttribPointer.Metadata = append(sc_GlVertexAttribPointer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttribPointer.xml]",
	})

	sc_GlGetActiveAttrib := schema.Of((*GlGetActiveAttrib)(nil).Class())
	sc_GlGetActiveAttrib.Metadata = append(sc_GlGetActiveAttrib.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveAttrib.xml]",
	})

	sc_GlGetActiveUniform := schema.Of((*GlGetActiveUniform)(nil).Class())
	sc_GlGetActiveUniform.Metadata = append(sc_GlGetActiveUniform.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveUniform.xml]",
	})

	sc_GlGetError := schema.Of((*GlGetError)(nil).Class())
	sc_GlGetError.Metadata = append(sc_GlGetError.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetError.xml]",
	})

	sc_GlGetProgramiv := schema.Of((*GlGetProgramiv)(nil).Class())
	sc_GlGetProgramiv.Metadata = append(sc_GlGetProgramiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgram.xml]",
	})

	sc_GlGetShaderiv := schema.Of((*GlGetShaderiv)(nil).Class())
	sc_GlGetShaderiv.Metadata = append(sc_GlGetShaderiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderiv.xml]",
	})

	sc_GlGetUniformLocation := schema.Of((*GlGetUniformLocation)(nil).Class())
	sc_GlGetUniformLocation.Metadata = append(sc_GlGetUniformLocation.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniformLocation.xml]",
	})

	sc_GlGetAttribLocation := schema.Of((*GlGetAttribLocation)(nil).Class())
	sc_GlGetAttribLocation.Metadata = append(sc_GlGetAttribLocation.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttribLocation.xml]",
	})

	sc_GlPixelStorei := schema.Of((*GlPixelStorei)(nil).Class())
	sc_GlPixelStorei.Metadata = append(sc_GlPixelStorei.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPixelStorei.xml]",
	})

	sc_GlTexParameteri := schema.Of((*GlTexParameteri)(nil).Class())
	sc_GlTexParameteri.Metadata = append(sc_GlTexParameteri.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
	})

	sc_GlTexParameterf := schema.Of((*GlTexParameterf)(nil).Class())
	sc_GlTexParameterf.Metadata = append(sc_GlTexParameterf.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
	})

	sc_GlGetTexParameteriv := schema.Of((*GlGetTexParameteriv)(nil).Class())
	sc_GlGetTexParameteriv.Metadata = append(sc_GlGetTexParameteriv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
	})

	sc_GlGetTexParameterfv := schema.Of((*GlGetTexParameterfv)(nil).Class())
	sc_GlGetTexParameterfv.Metadata = append(sc_GlGetTexParameterfv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
	})

	sc_GlUniform1i := schema.Of((*GlUniform1i)(nil).Class())
	sc_GlUniform1i.Metadata = append(sc_GlUniform1i.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform2i := schema.Of((*GlUniform2i)(nil).Class())
	sc_GlUniform2i.Metadata = append(sc_GlUniform2i.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform3i := schema.Of((*GlUniform3i)(nil).Class())
	sc_GlUniform3i.Metadata = append(sc_GlUniform3i.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform4i := schema.Of((*GlUniform4i)(nil).Class())
	sc_GlUniform4i.Metadata = append(sc_GlUniform4i.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform1iv := schema.Of((*GlUniform1iv)(nil).Class())
	sc_GlUniform1iv.Metadata = append(sc_GlUniform1iv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform2iv := schema.Of((*GlUniform2iv)(nil).Class())
	sc_GlUniform2iv.Metadata = append(sc_GlUniform2iv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform3iv := schema.Of((*GlUniform3iv)(nil).Class())
	sc_GlUniform3iv.Metadata = append(sc_GlUniform3iv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform4iv := schema.Of((*GlUniform4iv)(nil).Class())
	sc_GlUniform4iv.Metadata = append(sc_GlUniform4iv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform1f := schema.Of((*GlUniform1f)(nil).Class())
	sc_GlUniform1f.Metadata = append(sc_GlUniform1f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform2f := schema.Of((*GlUniform2f)(nil).Class())
	sc_GlUniform2f.Metadata = append(sc_GlUniform2f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform3f := schema.Of((*GlUniform3f)(nil).Class())
	sc_GlUniform3f.Metadata = append(sc_GlUniform3f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform4f := schema.Of((*GlUniform4f)(nil).Class())
	sc_GlUniform4f.Metadata = append(sc_GlUniform4f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform1fv := schema.Of((*GlUniform1fv)(nil).Class())
	sc_GlUniform1fv.Metadata = append(sc_GlUniform1fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform2fv := schema.Of((*GlUniform2fv)(nil).Class())
	sc_GlUniform2fv.Metadata = append(sc_GlUniform2fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform3fv := schema.Of((*GlUniform3fv)(nil).Class())
	sc_GlUniform3fv.Metadata = append(sc_GlUniform3fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniform4fv := schema.Of((*GlUniform4fv)(nil).Class())
	sc_GlUniform4fv.Metadata = append(sc_GlUniform4fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniformMatrix2fv := schema.Of((*GlUniformMatrix2fv)(nil).Class())
	sc_GlUniformMatrix2fv.Metadata = append(sc_GlUniformMatrix2fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniformMatrix3fv := schema.Of((*GlUniformMatrix3fv)(nil).Class())
	sc_GlUniformMatrix3fv.Metadata = append(sc_GlUniformMatrix3fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlUniformMatrix4fv := schema.Of((*GlUniformMatrix4fv)(nil).Class())
	sc_GlUniformMatrix4fv.Metadata = append(sc_GlUniformMatrix4fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
	})

	sc_GlGetUniformfv := schema.Of((*GlGetUniformfv)(nil).Class())
	sc_GlGetUniformfv.Metadata = append(sc_GlGetUniformfv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
	})

	sc_GlGetUniformiv := schema.Of((*GlGetUniformiv)(nil).Class())
	sc_GlGetUniformiv.Metadata = append(sc_GlGetUniformiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
	})

	sc_GlVertexAttrib1f := schema.Of((*GlVertexAttrib1f)(nil).Class())
	sc_GlVertexAttrib1f.Metadata = append(sc_GlVertexAttrib1f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib2f := schema.Of((*GlVertexAttrib2f)(nil).Class())
	sc_GlVertexAttrib2f.Metadata = append(sc_GlVertexAttrib2f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib3f := schema.Of((*GlVertexAttrib3f)(nil).Class())
	sc_GlVertexAttrib3f.Metadata = append(sc_GlVertexAttrib3f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib4f := schema.Of((*GlVertexAttrib4f)(nil).Class())
	sc_GlVertexAttrib4f.Metadata = append(sc_GlVertexAttrib4f.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib1fv := schema.Of((*GlVertexAttrib1fv)(nil).Class())
	sc_GlVertexAttrib1fv.Metadata = append(sc_GlVertexAttrib1fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib2fv := schema.Of((*GlVertexAttrib2fv)(nil).Class())
	sc_GlVertexAttrib2fv.Metadata = append(sc_GlVertexAttrib2fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib3fv := schema.Of((*GlVertexAttrib3fv)(nil).Class())
	sc_GlVertexAttrib3fv.Metadata = append(sc_GlVertexAttrib3fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlVertexAttrib4fv := schema.Of((*GlVertexAttrib4fv)(nil).Class())
	sc_GlVertexAttrib4fv.Metadata = append(sc_GlVertexAttrib4fv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
	})

	sc_GlGetShaderPrecisionFormat := schema.Of((*GlGetShaderPrecisionFormat)(nil).Class())
	sc_GlGetShaderPrecisionFormat.Metadata = append(sc_GlGetShaderPrecisionFormat.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml]",
	})

	sc_GlDepthMask := schema.Of((*GlDepthMask)(nil).Class())
	sc_GlDepthMask.Metadata = append(sc_GlDepthMask.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthMask.xml]",
	})

	sc_GlDepthFunc := schema.Of((*GlDepthFunc)(nil).Class())
	sc_GlDepthFunc.Metadata = append(sc_GlDepthFunc.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthFunc.xml]",
	})

	sc_GlDepthRangef := schema.Of((*GlDepthRangef)(nil).Class())
	sc_GlDepthRangef.Metadata = append(sc_GlDepthRangef.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthRangef.xml]",
	})

	sc_GlColorMask := schema.Of((*GlColorMask)(nil).Class())
	sc_GlColorMask.Metadata = append(sc_GlColorMask.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glColorMask.xml]",
	})

	sc_GlStencilMask := schema.Of((*GlStencilMask)(nil).Class())
	sc_GlStencilMask.Metadata = append(sc_GlStencilMask.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMask.xml]",
	})

	sc_GlStencilMaskSeparate := schema.Of((*GlStencilMaskSeparate)(nil).Class())
	sc_GlStencilMaskSeparate.Metadata = append(sc_GlStencilMaskSeparate.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMaskSeparate.xml]",
	})

	sc_GlStencilFuncSeparate := schema.Of((*GlStencilFuncSeparate)(nil).Class())
	sc_GlStencilFuncSeparate.Metadata = append(sc_GlStencilFuncSeparate.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilFuncSeparate.xml]",
	})

	sc_GlStencilOpSeparate := schema.Of((*GlStencilOpSeparate)(nil).Class())
	sc_GlStencilOpSeparate.Metadata = append(sc_GlStencilOpSeparate.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilOpSeparate.xml]",
	})

	sc_GlFrontFace := schema.Of((*GlFrontFace)(nil).Class())
	sc_GlFrontFace.Metadata = append(sc_GlFrontFace.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFrontFace.xml]",
	})

	sc_GlViewport := schema.Of((*GlViewport)(nil).Class())
	sc_GlViewport.Metadata = append(sc_GlViewport.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glViewport.xml]",
	})

	sc_GlScissor := schema.Of((*GlScissor)(nil).Class())
	sc_GlScissor.Metadata = append(sc_GlScissor.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glScissor.xml]",
	})

	sc_GlActiveTexture := schema.Of((*GlActiveTexture)(nil).Class())
	sc_GlActiveTexture.Metadata = append(sc_GlActiveTexture.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glActiveTexture.xml]",
	})

	sc_GlGenTextures := schema.Of((*GlGenTextures)(nil).Class())
	sc_GlGenTextures.Metadata = append(sc_GlGenTextures.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenTextures.xml]",
	})

	sc_GlDeleteTextures := schema.Of((*GlDeleteTextures)(nil).Class())
	sc_GlDeleteTextures.Metadata = append(sc_GlDeleteTextures.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteTextures.xml]",
	})

	sc_GlIsTexture := schema.Of((*GlIsTexture)(nil).Class())
	sc_GlIsTexture.Metadata = append(sc_GlIsTexture.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsTexture.xml]",
	})

	sc_GlBindTexture := schema.Of((*GlBindTexture)(nil).Class())
	sc_GlBindTexture.Metadata = append(sc_GlBindTexture.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindTexture.xml]",
	})

	sc_GlTexImage2D := schema.Of((*GlTexImage2D)(nil).Class())
	sc_GlTexImage2D.Metadata = append(sc_GlTexImage2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexImage2D.xml]",
	})

	sc_GlTexSubImage2D := schema.Of((*GlTexSubImage2D)(nil).Class())
	sc_GlTexSubImage2D.Metadata = append(sc_GlTexSubImage2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexSubImage2D.xml]",
	})

	sc_GlCopyTexImage2D := schema.Of((*GlCopyTexImage2D)(nil).Class())
	sc_GlCopyTexImage2D.Metadata = append(sc_GlCopyTexImage2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexImage2D.xml]",
	})

	sc_GlCopyTexSubImage2D := schema.Of((*GlCopyTexSubImage2D)(nil).Class())
	sc_GlCopyTexSubImage2D.Metadata = append(sc_GlCopyTexSubImage2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml]",
	})

	sc_GlCompressedTexImage2D := schema.Of((*GlCompressedTexImage2D)(nil).Class())
	sc_GlCompressedTexImage2D.Metadata = append(sc_GlCompressedTexImage2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexImage2D.xml]",
	})

	sc_GlCompressedTexSubImage2D := schema.Of((*GlCompressedTexSubImage2D)(nil).Class())
	sc_GlCompressedTexSubImage2D.Metadata = append(sc_GlCompressedTexSubImage2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml]",
	})

	sc_GlGenerateMipmap := schema.Of((*GlGenerateMipmap)(nil).Class())
	sc_GlGenerateMipmap.Metadata = append(sc_GlGenerateMipmap.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenerateMipmap.xml]",
	})

	sc_GlReadPixels := schema.Of((*GlReadPixels)(nil).Class())
	sc_GlReadPixels.Metadata = append(sc_GlReadPixels.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReadPixels.xml]",
	})

	sc_GlGenFramebuffers := schema.Of((*GlGenFramebuffers)(nil).Class())
	sc_GlGenFramebuffers.Metadata = append(sc_GlGenFramebuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenFramebuffers.xml]",
	})

	sc_GlBindFramebuffer := schema.Of((*GlBindFramebuffer)(nil).Class())
	sc_GlBindFramebuffer.Metadata = append(sc_GlBindFramebuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindFramebuffer.xml]",
	})

	sc_GlCheckFramebufferStatus := schema.Of((*GlCheckFramebufferStatus)(nil).Class())
	sc_GlCheckFramebufferStatus.Metadata = append(sc_GlCheckFramebufferStatus.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml]",
	})

	sc_GlDeleteFramebuffers := schema.Of((*GlDeleteFramebuffers)(nil).Class())
	sc_GlDeleteFramebuffers.Metadata = append(sc_GlDeleteFramebuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteFramebuffers.xml]",
	})

	sc_GlIsFramebuffer := schema.Of((*GlIsFramebuffer)(nil).Class())
	sc_GlIsFramebuffer.Metadata = append(sc_GlIsFramebuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsFramebuffer.xml]",
	})

	sc_GlGenRenderbuffers := schema.Of((*GlGenRenderbuffers)(nil).Class())
	sc_GlGenRenderbuffers.Metadata = append(sc_GlGenRenderbuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenRenderbuffers.xml]",
	})

	sc_GlBindRenderbuffer := schema.Of((*GlBindRenderbuffer)(nil).Class())
	sc_GlBindRenderbuffer.Metadata = append(sc_GlBindRenderbuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindRenderbuffer.xml]",
	})

	sc_GlRenderbufferStorage := schema.Of((*GlRenderbufferStorage)(nil).Class())
	sc_GlRenderbufferStorage.Metadata = append(sc_GlRenderbufferStorage.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glRenderbufferStorage.xml]",
	})

	sc_GlDeleteRenderbuffers := schema.Of((*GlDeleteRenderbuffers)(nil).Class())
	sc_GlDeleteRenderbuffers.Metadata = append(sc_GlDeleteRenderbuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml]",
	})

	sc_GlIsRenderbuffer := schema.Of((*GlIsRenderbuffer)(nil).Class())
	sc_GlIsRenderbuffer.Metadata = append(sc_GlIsRenderbuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsRenderbuffer.xml]",
	})

	sc_GlGetRenderbufferParameteriv := schema.Of((*GlGetRenderbufferParameteriv)(nil).Class())
	sc_GlGetRenderbufferParameteriv.Metadata = append(sc_GlGetRenderbufferParameteriv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml]",
	})

	sc_GlGenBuffers := schema.Of((*GlGenBuffers)(nil).Class())
	sc_GlGenBuffers.Metadata = append(sc_GlGenBuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml]",
	})

	sc_GlBindBuffer := schema.Of((*GlBindBuffer)(nil).Class())
	sc_GlBindBuffer.Metadata = append(sc_GlBindBuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindBuffer.xml]",
	})

	sc_GlBufferData := schema.Of((*GlBufferData)(nil).Class())
	sc_GlBufferData.Metadata = append(sc_GlBufferData.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferData.xml]",
	})

	sc_GlBufferSubData := schema.Of((*GlBufferSubData)(nil).Class())
	sc_GlBufferSubData.Metadata = append(sc_GlBufferSubData.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferSubData.xml]",
	})

	sc_GlDeleteBuffers := schema.Of((*GlDeleteBuffers)(nil).Class())
	sc_GlDeleteBuffers.Metadata = append(sc_GlDeleteBuffers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteBuffers.xml]",
	})

	sc_GlIsBuffer := schema.Of((*GlIsBuffer)(nil).Class())
	sc_GlIsBuffer.Metadata = append(sc_GlIsBuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsBuffer.xml]",
	})

	sc_GlGetBufferParameteriv := schema.Of((*GlGetBufferParameteriv)(nil).Class())
	sc_GlGetBufferParameteriv.Metadata = append(sc_GlGetBufferParameteriv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetBufferParameteriv.xml]",
	})

	sc_GlCreateShader := schema.Of((*GlCreateShader)(nil).Class())
	sc_GlCreateShader.Metadata = append(sc_GlCreateShader.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateShader.xml]",
	})

	sc_GlDeleteShader := schema.Of((*GlDeleteShader)(nil).Class())
	sc_GlDeleteShader.Metadata = append(sc_GlDeleteShader.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteShader.xml]",
	})

	sc_GlShaderSource := schema.Of((*GlShaderSource)(nil).Class())
	sc_GlShaderSource.Metadata = append(sc_GlShaderSource.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderSource.xml]",
	})

	sc_GlShaderBinary := schema.Of((*GlShaderBinary)(nil).Class())
	sc_GlShaderBinary.Metadata = append(sc_GlShaderBinary.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderBinary.xml]",
	})

	sc_GlGetShaderInfoLog := schema.Of((*GlGetShaderInfoLog)(nil).Class())
	sc_GlGetShaderInfoLog.Metadata = append(sc_GlGetShaderInfoLog.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderInfoLog.xml]",
	})

	sc_GlGetShaderSource := schema.Of((*GlGetShaderSource)(nil).Class())
	sc_GlGetShaderSource.Metadata = append(sc_GlGetShaderSource.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderSource.xml]",
	})

	sc_GlReleaseShaderCompiler := schema.Of((*GlReleaseShaderCompiler)(nil).Class())
	sc_GlReleaseShaderCompiler.Metadata = append(sc_GlReleaseShaderCompiler.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml]",
	})

	sc_GlCompileShader := schema.Of((*GlCompileShader)(nil).Class())
	sc_GlCompileShader.Metadata = append(sc_GlCompileShader.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompileShader.xml]",
	})

	sc_GlIsShader := schema.Of((*GlIsShader)(nil).Class())
	sc_GlIsShader.Metadata = append(sc_GlIsShader.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsShader.xml]",
	})

	sc_GlCreateProgram := schema.Of((*GlCreateProgram)(nil).Class())
	sc_GlCreateProgram.Metadata = append(sc_GlCreateProgram.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateProgram.xml]",
	})

	sc_GlDeleteProgram := schema.Of((*GlDeleteProgram)(nil).Class())
	sc_GlDeleteProgram.Metadata = append(sc_GlDeleteProgram.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteProgram.xml]",
	})

	sc_GlAttachShader := schema.Of((*GlAttachShader)(nil).Class())
	sc_GlAttachShader.Metadata = append(sc_GlAttachShader.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glAttachShader.xml]",
	})

	sc_GlDetachShader := schema.Of((*GlDetachShader)(nil).Class())
	sc_GlDetachShader.Metadata = append(sc_GlDetachShader.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDetachShader.xml]",
	})

	sc_GlGetAttachedShaders := schema.Of((*GlGetAttachedShaders)(nil).Class())
	sc_GlGetAttachedShaders.Metadata = append(sc_GlGetAttachedShaders.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttachedShaders.xml]",
	})

	sc_GlLinkProgram := schema.Of((*GlLinkProgram)(nil).Class())
	sc_GlLinkProgram.Metadata = append(sc_GlLinkProgram.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLinkProgram.xml]",
	})

	sc_GlGetProgramInfoLog := schema.Of((*GlGetProgramInfoLog)(nil).Class())
	sc_GlGetProgramInfoLog.Metadata = append(sc_GlGetProgramInfoLog.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgramInfoLog.xml]",
	})

	sc_GlUseProgram := schema.Of((*GlUseProgram)(nil).Class())
	sc_GlUseProgram.Metadata = append(sc_GlUseProgram.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUseProgram.xml]",
	})

	sc_GlIsProgram := schema.Of((*GlIsProgram)(nil).Class())
	sc_GlIsProgram.Metadata = append(sc_GlIsProgram.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsProgram.xml]",
	})

	sc_GlValidateProgram := schema.Of((*GlValidateProgram)(nil).Class())
	sc_GlValidateProgram.Metadata = append(sc_GlValidateProgram.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glValidateProgram.xml]",
	})

	sc_GlClearColor := schema.Of((*GlClearColor)(nil).Class())
	sc_GlClearColor.Metadata = append(sc_GlClearColor.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearColor.xml]",
	})

	sc_GlClearDepthf := schema.Of((*GlClearDepthf)(nil).Class())
	sc_GlClearDepthf.Metadata = append(sc_GlClearDepthf.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearDepthf.xml]",
	})

	sc_GlClearStencil := schema.Of((*GlClearStencil)(nil).Class())
	sc_GlClearStencil.Metadata = append(sc_GlClearStencil.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearStencil.xml]",
	})

	sc_GlClear := schema.Of((*GlClear)(nil).Class())
	sc_GlClear.Metadata = append(sc_GlClear.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClear.xml]",
	})

	sc_GlCullFace := schema.Of((*GlCullFace)(nil).Class())
	sc_GlCullFace.Metadata = append(sc_GlCullFace.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCullFace.xml]",
	})

	sc_GlPolygonOffset := schema.Of((*GlPolygonOffset)(nil).Class())
	sc_GlPolygonOffset.Metadata = append(sc_GlPolygonOffset.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPolygonOffset.xml]",
	})

	sc_GlLineWidth := schema.Of((*GlLineWidth)(nil).Class())
	sc_GlLineWidth.Metadata = append(sc_GlLineWidth.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLineWidth.xml]",
	})

	sc_GlSampleCoverage := schema.Of((*GlSampleCoverage)(nil).Class())
	sc_GlSampleCoverage.Metadata = append(sc_GlSampleCoverage.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glSampleCoverage.xml]",
	})

	sc_GlHint := schema.Of((*GlHint)(nil).Class())
	sc_GlHint.Metadata = append(sc_GlHint.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glHint.xml]",
	})

	sc_GlFramebufferRenderbuffer := schema.Of((*GlFramebufferRenderbuffer)(nil).Class())
	sc_GlFramebufferRenderbuffer.Metadata = append(sc_GlFramebufferRenderbuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml]",
	})

	sc_GlFramebufferTexture2D := schema.Of((*GlFramebufferTexture2D)(nil).Class())
	sc_GlFramebufferTexture2D.Metadata = append(sc_GlFramebufferTexture2D.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferTexture2D.xml]",
	})

	sc_GlGetFramebufferAttachmentParameteriv := schema.Of((*GlGetFramebufferAttachmentParameteriv)(nil).Class())
	sc_GlGetFramebufferAttachmentParameteriv.Metadata = append(sc_GlGetFramebufferAttachmentParameteriv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml]",
	})

	sc_GlDrawElements := schema.Of((*GlDrawElements)(nil).Class())
	sc_GlDrawElements.Metadata = append(sc_GlDrawElements.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0 | atom.DrawCall,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawElements.xml]",
	})

	sc_GlDrawArrays := schema.Of((*GlDrawArrays)(nil).Class())
	sc_GlDrawArrays.Metadata = append(sc_GlDrawArrays.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0 | atom.DrawCall,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawArrays.xml]",
	})

	sc_GlFlush := schema.Of((*GlFlush)(nil).Class())
	sc_GlFlush.Metadata = append(sc_GlFlush.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFlush.xml]",
	})

	sc_GlFinish := schema.Of((*GlFinish)(nil).Class())
	sc_GlFinish.Metadata = append(sc_GlFinish.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFinish.xml]",
	})

	sc_GlGetBooleanv := schema.Of((*GlGetBooleanv)(nil).Class())
	sc_GlGetBooleanv.Metadata = append(sc_GlGetBooleanv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})

	sc_GlGetFloatv := schema.Of((*GlGetFloatv)(nil).Class())
	sc_GlGetFloatv.Metadata = append(sc_GlGetFloatv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})

	sc_GlGetIntegerv := schema.Of((*GlGetIntegerv)(nil).Class())
	sc_GlGetIntegerv.Metadata = append(sc_GlGetIntegerv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
	})

	sc_GlGetString := schema.Of((*GlGetString)(nil).Class())
	sc_GlGetString.Metadata = append(sc_GlGetString.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetString.xml]",
	})

	sc_GlEnable := schema.Of((*GlEnable)(nil).Class())
	sc_GlEnable.Metadata = append(sc_GlEnable.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnable.xml]",
	})

	sc_GlDisable := schema.Of((*GlDisable)(nil).Class())
	sc_GlDisable.Metadata = append(sc_GlDisable.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisable.xml]",
	})

	sc_GlIsEnabled := schema.Of((*GlIsEnabled)(nil).Class())
	sc_GlIsEnabled.Metadata = append(sc_GlIsEnabled.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsEnabled.xml]",
	})

	sc_GlFenceSync := schema.Of((*GlFenceSync)(nil).Class())
	sc_GlFenceSync.Metadata = append(sc_GlFenceSync.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glFenceSync.xhtml]",
	})

	sc_GlDeleteSync := schema.Of((*GlDeleteSync)(nil).Class())
	sc_GlDeleteSync.Metadata = append(sc_GlDeleteSync.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteSync.xhtml]",
	})

	sc_GlWaitSync := schema.Of((*GlWaitSync)(nil).Class())
	sc_GlWaitSync.Metadata = append(sc_GlWaitSync.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glWaitSync.xhtml]",
	})

	sc_GlClientWaitSync := schema.Of((*GlClientWaitSync)(nil).Class())
	sc_GlClientWaitSync.Metadata = append(sc_GlClientWaitSync.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glClientWaitSync.xhtml]",
	})

	sc_GlMapBufferRange := schema.Of((*GlMapBufferRange)(nil).Class())
	sc_GlMapBufferRange.Metadata = append(sc_GlMapBufferRange.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
	})

	sc_GlUnmapBuffer := schema.Of((*GlUnmapBuffer)(nil).Class())
	sc_GlUnmapBuffer.Metadata = append(sc_GlUnmapBuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
	})

	sc_GlInvalidateFramebuffer := schema.Of((*GlInvalidateFramebuffer)(nil).Class())
	sc_GlInvalidateFramebuffer.Metadata = append(sc_GlInvalidateFramebuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glInvalidateFramebuffer.xhtml]",
	})

	sc_GlRenderbufferStorageMultisample := schema.Of((*GlRenderbufferStorageMultisample)(nil).Class())
	sc_GlRenderbufferStorageMultisample.Metadata = append(sc_GlRenderbufferStorageMultisample.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.opengl.org/registry/specs/EXT/framebuffer_multisample.txt]",
	})

	sc_GlBlitFramebuffer := schema.Of((*GlBlitFramebuffer)(nil).Class())
	sc_GlBlitFramebuffer.Metadata = append(sc_GlBlitFramebuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBlitFramebuffer.xhtml]",
	})

	sc_GlGenQueries := schema.Of((*GlGenQueries)(nil).Class())
	sc_GlGenQueries.Metadata = append(sc_GlGenQueries.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenQueries.xhtml]",
	})

	sc_GlBeginQuery := schema.Of((*GlBeginQuery)(nil).Class())
	sc_GlBeginQuery.Metadata = append(sc_GlBeginQuery.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBeginQuery.xhtml]",
	})

	sc_GlEndQuery := schema.Of((*GlEndQuery)(nil).Class())
	sc_GlEndQuery.Metadata = append(sc_GlEndQuery.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glEndQuery.xhtml]",
	})

	sc_GlDeleteQueries := schema.Of((*GlDeleteQueries)(nil).Class())
	sc_GlDeleteQueries.Metadata = append(sc_GlDeleteQueries.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteQueries.xhtml]",
	})

	sc_GlIsQuery := schema.Of((*GlIsQuery)(nil).Class())
	sc_GlIsQuery.Metadata = append(sc_GlIsQuery.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glIsQuery.xhtml]",
	})

	sc_GlGetQueryiv := schema.Of((*GlGetQueryiv)(nil).Class())
	sc_GlGetQueryiv.Metadata = append(sc_GlGetQueryiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryiv.xhtml]",
	})

	sc_GlGetQueryObjectuiv := schema.Of((*GlGetQueryObjectuiv)(nil).Class())
	sc_GlGetQueryObjectuiv.Metadata = append(sc_GlGetQueryObjectuiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryObjectuiv.xhtml]",
	})

	sc_GlGetActiveUniformBlockName := schema.Of((*GlGetActiveUniformBlockName)(nil).Class())
	sc_GlGetActiveUniformBlockName.Metadata = append(sc_GlGetActiveUniformBlockName.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniformBlockName.xhtml]",
	})

	sc_GlGetActiveUniformBlockiv := schema.Of((*GlGetActiveUniformBlockiv)(nil).Class())
	sc_GlGetActiveUniformBlockiv.Metadata = append(sc_GlGetActiveUniformBlockiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniformBlockiv.xhtml]",
	})

	sc_GlUniformBlockBinding := schema.Of((*GlUniformBlockBinding)(nil).Class())
	sc_GlUniformBlockBinding.Metadata = append(sc_GlUniformBlockBinding.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glUniformBlockBinding.xhtml]",
	})

	sc_GlGetActiveUniformsiv := schema.Of((*GlGetActiveUniformsiv)(nil).Class())
	sc_GlGetActiveUniformsiv.Metadata = append(sc_GlGetActiveUniformsiv.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniformsiv.xhtml]",
	})

	sc_GlBindBufferBase := schema.Of((*GlBindBufferBase)(nil).Class())
	sc_GlBindBufferBase.Metadata = append(sc_GlBindBufferBase.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBindBufferBase.xhtml]",
	})

	sc_GlGenVertexArrays := schema.Of((*GlGenVertexArrays)(nil).Class())
	sc_GlGenVertexArrays.Metadata = append(sc_GlGenVertexArrays.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenVertexArrays.xhtml]",
	})

	sc_GlBindVertexArray := schema.Of((*GlBindVertexArray)(nil).Class())
	sc_GlBindVertexArray.Metadata = append(sc_GlBindVertexArray.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBindVertexArray.xhtml]",
	})

	sc_GlDeleteVertexArrays := schema.Of((*GlDeleteVertexArrays)(nil).Class())
	sc_GlDeleteVertexArrays.Metadata = append(sc_GlDeleteVertexArrays.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteVertexArrays.xhtml]",
	})

	sc_GlGetQueryObjecti64v := schema.Of((*GlGetQueryObjecti64v)(nil).Class())
	sc_GlGetQueryObjecti64v.Metadata = append(sc_GlGetQueryObjecti64v.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlGetQueryObjectui64v := schema.Of((*GlGetQueryObjectui64v)(nil).Class())
	sc_GlGetQueryObjectui64v.Metadata = append(sc_GlGetQueryObjectui64v.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_GlGenQueriesEXT := schema.Of((*GlGenQueriesEXT)(nil).Class())
	sc_GlGenQueriesEXT.Metadata = append(sc_GlGenQueriesEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlBeginQueryEXT := schema.Of((*GlBeginQueryEXT)(nil).Class())
	sc_GlBeginQueryEXT.Metadata = append(sc_GlBeginQueryEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlEndQueryEXT := schema.Of((*GlEndQueryEXT)(nil).Class())
	sc_GlEndQueryEXT.Metadata = append(sc_GlEndQueryEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlDeleteQueriesEXT := schema.Of((*GlDeleteQueriesEXT)(nil).Class())
	sc_GlDeleteQueriesEXT.Metadata = append(sc_GlDeleteQueriesEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlIsQueryEXT := schema.Of((*GlIsQueryEXT)(nil).Class())
	sc_GlIsQueryEXT.Metadata = append(sc_GlIsQueryEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlQueryCounterEXT := schema.Of((*GlQueryCounterEXT)(nil).Class())
	sc_GlQueryCounterEXT.Metadata = append(sc_GlQueryCounterEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlGetQueryivEXT := schema.Of((*GlGetQueryivEXT)(nil).Class())
	sc_GlGetQueryivEXT.Metadata = append(sc_GlGetQueryivEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlGetQueryObjectivEXT := schema.Of((*GlGetQueryObjectivEXT)(nil).Class())
	sc_GlGetQueryObjectivEXT.Metadata = append(sc_GlGetQueryObjectivEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlGetQueryObjectuivEXT := schema.Of((*GlGetQueryObjectuivEXT)(nil).Class())
	sc_GlGetQueryObjectuivEXT.Metadata = append(sc_GlGetQueryObjectuivEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlGetQueryObjecti64vEXT := schema.Of((*GlGetQueryObjecti64vEXT)(nil).Class())
	sc_GlGetQueryObjecti64vEXT.Metadata = append(sc_GlGetQueryObjecti64vEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_GlGetQueryObjectui64vEXT := schema.Of((*GlGetQueryObjectui64vEXT)(nil).Class())
	sc_GlGetQueryObjectui64vEXT.Metadata = append(sc_GlGetQueryObjectui64vEXT.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
	})

	sc_Architecture := schema.Of((*Architecture)(nil).Class())
	sc_Architecture.Metadata = append(sc_Architecture.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_ReplayCreateRenderer := schema.Of((*ReplayCreateRenderer)(nil).Class())
	sc_ReplayCreateRenderer.Metadata = append(sc_ReplayCreateRenderer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_ReplayBindRenderer := schema.Of((*ReplayBindRenderer)(nil).Class())
	sc_ReplayBindRenderer.Metadata = append(sc_ReplayBindRenderer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_BackbufferInfo := schema.Of((*BackbufferInfo)(nil).Class())
	sc_BackbufferInfo.Metadata = append(sc_BackbufferInfo.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_StartTimer := schema.Of((*StartTimer)(nil).Class())
	sc_StartTimer.Metadata = append(sc_StartTimer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_StopTimer := schema.Of((*StopTimer)(nil).Class())
	sc_StopTimer.Metadata = append(sc_StopTimer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_FlushPostBuffer := schema.Of((*FlushPostBuffer)(nil).Class())
	sc_FlushPostBuffer.Metadata = append(sc_FlushPostBuffer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})
}
