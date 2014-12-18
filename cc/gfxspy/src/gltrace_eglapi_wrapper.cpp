/*
 * Copyright 2014, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//////////////////////////////////////////////////////////////////////////
// This file contains wrapper functions for all EGL functions
// found in egl_trace.in.  Currently, only 3 of these functions do
// anything other than print a message and call the corresponding
// EGL function.  These 3 functions are:
//     EGLTrace_wrapper_eglCreateContext
//     EGLTrace_wrapper_eglMakeCurrent
//     EGLTrace_wrapper_eglSwapBuffers
//
//
// To use Visual Studio to expand egl_trace.in into the wrappers
// found in this file, perform a find and replace with regular
// expressions:
//
// Find what:
// TRACE_EGL\({[^,]+}, {[^,]+}, {\([^\)]+\)}, {\([^\)]*\)}\)
//
// REplace with:
// EGLAPI \1 EGLTrace_wrapper_\2\3 {\n    LOG_EGL_CALL("EGLTrace_wrapper_\2\3...");\n    \1 return_value = gEGLImpl.egl.\2\4;\n    return return_value;\n}\n
//////////////////////////////////////////////////////////////////////////

#include "EGL/egldefs.h"
#include "EGL/Loader.h"
#include "glestrace.h"
#include "gltrace_context.h"
#include "gltrace_egl.h"
#include "gltrace_transport.h"
#include "hooks.h"
#include "log/log.h"


namespace android {
namespace gltrace {

extern egl_connection_t gEGLImpl;

#if defined(GLTRACE_PRINT_EGL_CALLS)
#define LOG_EGL_CALL ALOGD
#else
#define LOG_EGL_CALL(...)
#endif

static int guessGlesVersion(EGLContext context, const EGLint *attrib_list) {
    if (context != EGL_NO_CONTEXT) {
        if (attrib_list) {
            while (*attrib_list != EGL_NONE) {
                const GLint attr = *attrib_list++;
                const GLint value = *attrib_list++;
                if (attr == EGL_CONTEXT_CLIENT_VERSION) {
                    switch (value) {
                        case 1:
                            return egl_connection_t::GLESv1_INDEX;
                            break;
                        case 2:
                        case 3:
                            return egl_connection_t::GLESv2_INDEX;
                            break;
                        default:
                            ALOGD("<gfxspy> Unknown GLES version %d, skipping.", value);
                    }
                }
            }
        }
    }

    ALOGD("<gfxspy> GLES version unspecified, defaulting to GLESv1.");
    return egl_connection_t::GLESv1_INDEX;
}

EGLAPI EGLint EGLTrace_wrapper_eglGetError(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetError(void)...[%p]", gEGLImpl.egl.eglGetError);
    EGLint return_value = gEGLImpl.egl.eglGetError();
    return return_value;
}

EGLAPI EGLDisplay EGLTrace_wrapper_eglGetDisplay(EGLNativeDisplayType display_id) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetDisplay(EGLNativeDisplayType display_id)...[%p]", gEGLImpl.egl.eglGetDisplay);
    EGLDisplay return_value = gEGLImpl.egl.eglGetDisplay(display_id);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglInitialize(EGLDisplay dpy, EGLint *major, EGLint *minor) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglInitialize(EGLDisplay dpy, EGLint *major, EGLint *minor)...[%p]", gEGLImpl.egl.eglInitialize);
    EGLBoolean return_value = gEGLImpl.egl.eglInitialize(dpy, major, minor);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglTerminate(EGLDisplay dpy) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglTerminate(EGLDisplay dpy)...[%p]", gEGLImpl.egl.eglTerminate);
    EGLBoolean return_value = gEGLImpl.egl.eglTerminate(dpy);
    return return_value;
}

EGLAPI const char * EGLTrace_wrapper_eglQueryString(EGLDisplay dpy, EGLint name) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQueryString(EGLDisplay dpy, EGLint name)...[%p]", gEGLImpl.egl.eglTerminate);
    const char * return_value = gEGLImpl.egl.eglQueryString(dpy, name);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglGetConfigs(EGLDisplay dpy, EGLConfig *configs, EGLint config_size, EGLint *num_config) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetConfigs(EGLDisplay dpy, EGLConfig *configs, EGLint config_size, EGLint *num_config)...[%p]", gEGLImpl.egl.eglGetConfigs);
    EGLBoolean return_value = gEGLImpl.egl.eglGetConfigs(dpy, configs, config_size, num_config);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglChooseConfig(EGLDisplay dpy, const EGLint *attrib_list, EGLConfig *configs, EGLint config_size, EGLint *num_config) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglChooseConfig(EGLDisplay dpy, const EGLint *attrib_list, EGLConfig *configs, EGLint config_size, EGLint *num_config)...[%p]", gEGLImpl.egl.eglChooseConfig);
    EGLBoolean return_value = gEGLImpl.egl.eglChooseConfig(dpy, attrib_list, configs, config_size, num_config);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglGetConfigAttrib(EGLDisplay dpy, EGLConfig config, EGLint attribute, EGLint *value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetConfigAttrib(EGLDisplay dpy, EGLConfig config, EGLint attribute, EGLint *value)...[%p]", gEGLImpl.egl.eglGetConfigAttrib);
    EGLBoolean return_value = gEGLImpl.egl.eglGetConfigAttrib(dpy, config, attribute, value);
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreateWindowSurface(EGLDisplay dpy, EGLConfig config, EGLNativeWindowType win, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreateWindowSurface(EGLDisplay dpy, EGLConfig config, EGLNativeWindowType win, const EGLint *attrib_list)...[%p]", gEGLImpl.egl.eglCreateWindowSurface);
    EGLSurface return_value = gEGLImpl.egl.eglCreateWindowSurface(dpy, config, win, attrib_list);
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreatePbufferSurface(EGLDisplay dpy, EGLConfig config, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreatePbufferSurface(EGLDisplay dpy, EGLConfig config, const EGLint *attrib_list)...[%p]", gEGLImpl.egl.eglCreatePbufferSurface);
    EGLSurface return_value = gEGLImpl.egl.eglCreatePbufferSurface(dpy, config, attrib_list);
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreatePixmapSurface(EGLDisplay dpy, EGLConfig config, EGLNativePixmapType pixmap, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreatePixmapSurface(EGLDisplay dpy, EGLConfig config, EGLNativePixmapType pixmap, const EGLint *attrib_list)...[%p]", gEGLImpl.egl.eglCreatePixmapSurface);
    EGLSurface return_value = gEGLImpl.egl.eglCreatePixmapSurface(dpy, config, pixmap, attrib_list);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglDestroySurface(EGLDisplay dpy, EGLSurface surface) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglDestroySurface(EGLDisplay dpy, EGLSurface surface)...[%p]", gEGLImpl.egl.eglDestroySurface);
    EGLBoolean return_value = gEGLImpl.egl.eglDestroySurface(dpy, surface);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglQuerySurface(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint *value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQuerySurface(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint *value)...[%p]", gEGLImpl.egl.eglQuerySurface);
    EGLBoolean return_value = gEGLImpl.egl.eglQuerySurface(dpy, surface, attribute, value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglBindAPI(EGLenum api) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglBindAPI(EGLenum api)...[%p]", gEGLImpl.egl.eglBindAPI);
    EGLBoolean return_value = gEGLImpl.egl.eglBindAPI(api);
    return return_value;
}

EGLAPI EGLenum EGLTrace_wrapper_eglQueryAPI(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQueryAPI(void)...[%p]", gEGLImpl.egl.eglQueryAPI);
    EGLenum return_value = gEGLImpl.egl.eglQueryAPI();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglWaitClient(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglWaitClient(void)...[%p]", gEGLImpl.egl.eglWaitClient);
    EGLBoolean return_value = gEGLImpl.egl.eglWaitClient();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglReleaseThread(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglReleaseThread(void)...[%p]", gEGLImpl.egl.eglReleaseThread);
    EGLBoolean return_value = gEGLImpl.egl.eglReleaseThread();
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreatePbufferFromClientBuffer(EGLDisplay dpy, EGLenum buftype, EGLClientBuffer buffer, EGLConfig config, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreatePbufferFromClientBuffer(EGLDisplay dpy, EGLenum buftype, EGLClientBuffer buffer, EGLConfig config, const EGLint *attrib_list)...[%p]", gEGLImpl.egl.eglCreatePbufferFromClientBuffer);
    EGLSurface return_value = gEGLImpl.egl.eglCreatePbufferFromClientBuffer(dpy, buftype, buffer, config, attrib_list);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglSurfaceAttrib(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglSurfaceAttrib(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint value)...[%p]", gEGLImpl.egl.eglSurfaceAttrib);
    EGLBoolean return_value = gEGLImpl.egl.eglSurfaceAttrib(dpy, surface, attribute, value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglBindTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglBindTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer)...[%p]", gEGLImpl.egl.eglBindTexImage);
    EGLBoolean return_value = gEGLImpl.egl.eglBindTexImage(dpy, surface, buffer);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglReleaseTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglReleaseTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer)...[%p]", gEGLImpl.egl.eglReleaseTexImage);
    EGLBoolean return_value = gEGLImpl.egl.eglReleaseTexImage(dpy, surface, buffer);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglSwapInterval(EGLDisplay dpy, EGLint interval) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglSwapInterval(EGLDisplay dpy, EGLint interval)...[%p]", gEGLImpl.egl.eglSwapInterval);
    EGLBoolean return_value = gEGLImpl.egl.eglSwapInterval(dpy, interval);
    return return_value;
}

EGLAPI EGLContext EGLTrace_wrapper_eglCreateContext(EGLDisplay dpy, EGLConfig config, EGLContext share_context, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreateContext(EGLDisplay dpy, EGLConfig config, EGLContext share_context, const EGLint *attrib_list)...[%p]", gEGLImpl.egl.eglCreateContext);
    EGLContext return_value = gEGLImpl.egl.eglCreateContext(dpy, config, share_context, attrib_list);
    GLTrace_eglCreateContext(guessGlesVersion(return_value, attrib_list), return_value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglDestroyContext(EGLDisplay dpy, EGLContext ctx) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglDestroyContext(EGLDisplay dpy, EGLContext ctx)...[%p]", gEGLImpl.egl.eglDestroyContext);
    EGLBoolean return_value = gEGLImpl.egl.eglDestroyContext(dpy, ctx);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglMakeCurrent(EGLDisplay dpy, EGLSurface draw, EGLSurface read, EGLContext ctx) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglMakeCurrent(EGLDisplay dpy, EGLSurface draw, EGLSurface read, EGLContext ctx)...[%p]", gEGLImpl.egl.eglMakeCurrent);
    EGLBoolean return_value = gEGLImpl.egl.eglMakeCurrent(dpy, draw, read, ctx);
    if (return_value != EGL_FALSE) {
        GLTrace_eglMakeCurrent(ctx);
    }
    return return_value;
}

EGLAPI EGLContext EGLTrace_wrapper_eglGetCurrentContext(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetCurrentContext(void)...[%p]", gEGLImpl.egl.eglGetCurrentContext);
    EGLContext return_value = gEGLImpl.egl.eglGetCurrentContext();
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglGetCurrentSurface(EGLint readdraw) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetCurrentSurface(EGLint readdraw)...[%p]", gEGLImpl.egl.eglGetCurrentSurface);
    EGLSurface return_value = gEGLImpl.egl.eglGetCurrentSurface(readdraw);
    return return_value;
}

EGLAPI EGLDisplay EGLTrace_wrapper_eglGetCurrentDisplay(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetCurrentDisplay(void)...[%p]", gEGLImpl.egl.eglGetCurrentDisplay);
    EGLDisplay return_value = gEGLImpl.egl.eglGetCurrentDisplay();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglQueryContext(EGLDisplay dpy, EGLContext ctx, EGLint attribute, EGLint *value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQueryContext(EGLDisplay dpy, EGLContext ctx, EGLint attribute, EGLint *value)...[%p]", gEGLImpl.egl.eglQueryContext);
    EGLBoolean return_value = gEGLImpl.egl.eglQueryContext(dpy, ctx, attribute, value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglWaitGL(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglWaitGL(void)...[%p]", gEGLImpl.egl.eglWaitGL);
    EGLBoolean return_value = gEGLImpl.egl.eglWaitGL();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglWaitNative(EGLint engine) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglWaitNative(EGLint engine)...[%p]", gEGLImpl.egl.eglWaitNative);
    EGLBoolean return_value = gEGLImpl.egl.eglWaitNative(engine);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglSwapBuffers(EGLDisplay dpy, EGLSurface surface) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglSwapBuffers(EGLDisplay dpy, EGLSurface surface)...[%p]", gEGLImpl.egl.eglSwapBuffers);
    EGLBoolean return_value = gEGLImpl.egl.eglSwapBuffers(dpy, surface);
    if (return_value) {
        GLTrace_eglSwapBuffers_internal(dpy, surface);
    }
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglCopyBuffers(EGLDisplay dpy, EGLSurface surface, EGLNativePixmapType target) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCopyBuffers(EGLDisplay dpy, EGLSurface surface, EGLNativePixmapType target)...[%p]", gEGLImpl.egl.eglCopyBuffers);
    EGLBoolean return_value = gEGLImpl.egl.eglCopyBuffers(dpy, surface, target);
    return return_value;
}

EGLAPI __eglMustCastToProperFunctionPointerType EGLTrace_wrapper_eglGetProcAddress(const char *procname) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetProcAddress(const char *procname)...[%p]", gEGLImpl.egl.eglGetProcAddress);
    __eglMustCastToProperFunctionPointerType return_value = gEGLImpl.egl.eglGetProcAddress(procname);
    return return_value;
}

#undef LOG_EGL_CALL

} // end of namespace gltrace
} // end of namespace android
