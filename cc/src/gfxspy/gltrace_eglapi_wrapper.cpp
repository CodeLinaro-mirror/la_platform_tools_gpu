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
// EGLAPI \1 EGLTrace_wrapper_\2\3 {\n    LOG_EGL_CALL("EGLTrace_wrapper_\2\3...");\n    \1 return_value = getEGLHooks()->\2\4;\n    return return_value;\n}\n
//////////////////////////////////////////////////////////////////////////

#include <cutils/log.h>

#include "EGL/Loader.h"
#include "EGL/egldefs.h"
#include "EGL/egl_object.h"
#include "glestrace.h"
#include "gltrace_context.h"
#include "gltrace_egl.h"
#include "gltrace_hooks.h"
#include "gltrace_transport.h"
#include "hooks.h"

namespace android {
namespace gltrace {

extern egl_connection_t gEGLImpl;
extern EGLBoolean egl_init_drivers();

#if defined(GLTRACE_PRINT_EGL_CALLS)
#define LOG_EGL_CALL ALOGD
#else
#define LOG_EGL_CALL(...)
#endif

egl_t* getEGLHooks() {
    if (egl_init_drivers()) {
        return &gEGLImpl.egl;
    }
    return NULL;
}

EGLAPI EGLint EGLTrace_wrapper_eglGetError(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetError(void)...[%p]", getEGLHooks()->eglGetError);
    EGLint return_value = getEGLHooks()->eglGetError();
    return return_value;
}

EGLAPI EGLDisplay EGLTrace_wrapper_eglGetDisplay(EGLNativeDisplayType display_id) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetDisplay(EGLNativeDisplayType display_id)...[%p]", getEGLHooks()->eglGetDisplay);
    EGLDisplay return_value = getEGLHooks()->eglGetDisplay(display_id);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglInitialize(EGLDisplay dpy, EGLint *major, EGLint *minor) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglInitialize(EGLDisplay dpy, EGLint *major, EGLint *minor)...[%p]", getEGLHooks()->eglInitialize);
    EGLBoolean return_value = getEGLHooks()->eglInitialize(dpy, major, minor);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglTerminate(EGLDisplay dpy) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglTerminate(EGLDisplay dpy)...[%p]", getEGLHooks()->eglTerminate);
    EGLBoolean return_value = getEGLHooks()->eglTerminate(dpy);
    return return_value;
}

EGLAPI const char * EGLTrace_wrapper_eglQueryString(EGLDisplay dpy, EGLint name) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQueryString(EGLDisplay dpy, EGLint name)...[%p]", getEGLHooks()->eglTerminate);
    const char * return_value = getEGLHooks()->eglQueryString(dpy, name);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglGetConfigs(EGLDisplay dpy, EGLConfig *configs, EGLint config_size, EGLint *num_config) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetConfigs(EGLDisplay dpy, EGLConfig *configs, EGLint config_size, EGLint *num_config)...[%p]", getEGLHooks()->eglGetConfigs);
    EGLBoolean return_value = getEGLHooks()->eglGetConfigs(dpy, configs, config_size, num_config);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglChooseConfig(EGLDisplay dpy, const EGLint *attrib_list, EGLConfig *configs, EGLint config_size, EGLint *num_config) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglChooseConfig(EGLDisplay dpy, const EGLint *attrib_list, EGLConfig *configs, EGLint config_size, EGLint *num_config)...[%p]", getEGLHooks()->eglChooseConfig);
    EGLBoolean return_value = getEGLHooks()->eglChooseConfig(dpy, attrib_list, configs, config_size, num_config);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglGetConfigAttrib(EGLDisplay dpy, EGLConfig config, EGLint attribute, EGLint *value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetConfigAttrib(EGLDisplay dpy, EGLConfig config, EGLint attribute, EGLint *value)...[%p]", getEGLHooks()->eglGetConfigAttrib);
    EGLBoolean return_value = getEGLHooks()->eglGetConfigAttrib(dpy, config, attribute, value);
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreateWindowSurface(EGLDisplay dpy, EGLConfig config, EGLNativeWindowType win, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreateWindowSurface(EGLDisplay dpy, EGLConfig config, EGLNativeWindowType win, const EGLint *attrib_list)...[%p]", getEGLHooks()->eglCreateWindowSurface);
    EGLSurface return_value = getEGLHooks()->eglCreateWindowSurface(dpy, config, win, attrib_list);
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreatePbufferSurface(EGLDisplay dpy, EGLConfig config, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreatePbufferSurface(EGLDisplay dpy, EGLConfig config, const EGLint *attrib_list)...[%p]", getEGLHooks()->eglCreatePbufferSurface);
    EGLSurface return_value = getEGLHooks()->eglCreatePbufferSurface(dpy, config, attrib_list);
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreatePixmapSurface(EGLDisplay dpy, EGLConfig config, EGLNativePixmapType pixmap, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreatePixmapSurface(EGLDisplay dpy, EGLConfig config, EGLNativePixmapType pixmap, const EGLint *attrib_list)...[%p]", getEGLHooks()->eglCreatePixmapSurface);
    EGLSurface return_value = getEGLHooks()->eglCreatePixmapSurface(dpy, config, pixmap, attrib_list);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglDestroySurface(EGLDisplay dpy, EGLSurface surface) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglDestroySurface(EGLDisplay dpy, EGLSurface surface)...[%p]", getEGLHooks()->eglDestroySurface);
    EGLBoolean return_value = getEGLHooks()->eglDestroySurface(dpy, surface);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglQuerySurface(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint *value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQuerySurface(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint *value)...[%p]", getEGLHooks()->eglQuerySurface);
    EGLBoolean return_value = getEGLHooks()->eglQuerySurface(dpy, surface, attribute, value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglBindAPI(EGLenum api) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglBindAPI(EGLenum api)...[%p]", getEGLHooks()->eglBindAPI);
    EGLBoolean return_value = getEGLHooks()->eglBindAPI(api);
    return return_value;
}

EGLAPI EGLenum EGLTrace_wrapper_eglQueryAPI(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQueryAPI(void)...[%p]", getEGLHooks()->eglQueryAPI);
    EGLenum return_value = getEGLHooks()->eglQueryAPI();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglWaitClient(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglWaitClient(void)...[%p]", getEGLHooks()->eglWaitClient);
    EGLBoolean return_value = getEGLHooks()->eglWaitClient();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglReleaseThread(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglReleaseThread(void)...[%p]", getEGLHooks()->eglReleaseThread);
    EGLBoolean return_value = getEGLHooks()->eglReleaseThread();
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglCreatePbufferFromClientBuffer(EGLDisplay dpy, EGLenum buftype, EGLClientBuffer buffer, EGLConfig config, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreatePbufferFromClientBuffer(EGLDisplay dpy, EGLenum buftype, EGLClientBuffer buffer, EGLConfig config, const EGLint *attrib_list)...[%p]", getEGLHooks()->eglCreatePbufferFromClientBuffer);
    EGLSurface return_value = getEGLHooks()->eglCreatePbufferFromClientBuffer(dpy, buftype, buffer, config, attrib_list);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglSurfaceAttrib(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglSurfaceAttrib(EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint value)...[%p]", getEGLHooks()->eglSurfaceAttrib);
    EGLBoolean return_value = getEGLHooks()->eglSurfaceAttrib(dpy, surface, attribute, value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglBindTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglBindTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer)...[%p]", getEGLHooks()->eglBindTexImage);
    EGLBoolean return_value = getEGLHooks()->eglBindTexImage(dpy, surface, buffer);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglReleaseTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglReleaseTexImage(EGLDisplay dpy, EGLSurface surface, EGLint buffer)...[%p]", getEGLHooks()->eglReleaseTexImage);
    EGLBoolean return_value = getEGLHooks()->eglReleaseTexImage(dpy, surface, buffer);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglSwapInterval(EGLDisplay dpy, EGLint interval) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglSwapInterval(EGLDisplay dpy, EGLint interval)...[%p]", getEGLHooks()->eglSwapInterval);
    EGLBoolean return_value = getEGLHooks()->eglSwapInterval(dpy, interval);
    return return_value;
}

EGLAPI EGLContext EGLTrace_wrapper_eglCreateContext(EGLDisplay dpy, EGLConfig config, EGLContext share_context, const EGLint *attrib_list) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCreateContext(EGLDisplay dpy, EGLConfig config, EGLContext share_context, const EGLint *attrib_list)...[%p]", getEGLHooks()->eglCreateContext);
    EGLContext return_value = getEGLHooks()->eglCreateContext(dpy, config, share_context, attrib_list);
    GLTrace_eglCreateContext(android::egl_connection_t::GLESv2_INDEX, return_value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglDestroyContext(EGLDisplay dpy, EGLContext ctx) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglDestroyContext(EGLDisplay dpy, EGLContext ctx)...[%p]", getEGLHooks()->eglDestroyContext);
    EGLBoolean return_value = getEGLHooks()->eglDestroyContext(dpy, ctx);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglMakeCurrent(EGLDisplay dpy, EGLSurface draw, EGLSurface read, EGLContext ctx) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglMakeCurrent(EGLDisplay dpy, EGLSurface draw, EGLSurface read, EGLContext ctx)...[%p]", getEGLHooks()->eglMakeCurrent);
    EGLBoolean return_value = getEGLHooks()->eglMakeCurrent(dpy, draw, read, ctx);
    if (return_value != EGL_FALSE) {
        GLTrace_eglMakeCurrent(android::egl_connection_t::GLESv2_INDEX, NULL, ctx);
    }
    return return_value;
}

EGLAPI EGLContext EGLTrace_wrapper_eglGetCurrentContext(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetCurrentContext(void)...[%p]", getEGLHooks()->eglGetCurrentContext);
    EGLContext return_value = getEGLHooks()->eglGetCurrentContext();
    return return_value;
}

EGLAPI EGLSurface EGLTrace_wrapper_eglGetCurrentSurface(EGLint readdraw) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetCurrentSurface(EGLint readdraw)...[%p]", getEGLHooks()->eglGetCurrentSurface);
    EGLSurface return_value = getEGLHooks()->eglGetCurrentSurface(readdraw);
    return return_value;
}

EGLAPI EGLDisplay EGLTrace_wrapper_eglGetCurrentDisplay(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetCurrentDisplay(void)...[%p]", getEGLHooks()->eglGetCurrentDisplay);
    EGLDisplay return_value = getEGLHooks()->eglGetCurrentDisplay();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglQueryContext(EGLDisplay dpy, EGLContext ctx, EGLint attribute, EGLint *value) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglQueryContext(EGLDisplay dpy, EGLContext ctx, EGLint attribute, EGLint *value)...[%p]", getEGLHooks()->eglQueryContext);
    EGLBoolean return_value = getEGLHooks()->eglQueryContext(dpy, ctx, attribute, value);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglWaitGL(void) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglWaitGL(void)...[%p]", getEGLHooks()->eglWaitGL);
    EGLBoolean return_value = getEGLHooks()->eglWaitGL();
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglWaitNative(EGLint engine) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglWaitNative(EGLint engine)...[%p]", getEGLHooks()->eglWaitNative);
    EGLBoolean return_value = getEGLHooks()->eglWaitNative(engine);
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglSwapBuffers(EGLDisplay dpy, EGLSurface surface) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglSwapBuffers(EGLDisplay dpy, EGLSurface surface)...[%p]", getEGLHooks()->eglSwapBuffers);
    EGLBoolean return_value = getEGLHooks()->eglSwapBuffers(dpy, surface);
    if (return_value) {
        GLTrace_eglSwapBuffers_internal(dpy, surface);
    }
    return return_value;
}

EGLAPI EGLBoolean EGLTrace_wrapper_eglCopyBuffers(EGLDisplay dpy, EGLSurface surface, EGLNativePixmapType target) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglCopyBuffers(EGLDisplay dpy, EGLSurface surface, EGLNativePixmapType target)...[%p]", getEGLHooks()->eglCopyBuffers);
    EGLBoolean return_value = getEGLHooks()->eglCopyBuffers(dpy, surface, target);
    return return_value;
}

EGLAPI __eglMustCastToProperFunctionPointerType EGLTrace_wrapper_eglGetProcAddress(const char *procname) {
    LOG_EGL_CALL("EGLTrace_wrapper_eglGetProcAddress(const char *procname)...[%p]", getEGLHooks()->eglGetProcAddress);
    __eglMustCastToProperFunctionPointerType return_value = getEGLHooks()->eglGetProcAddress(procname);
    return return_value;
}

#undef LOG_EGL_CALL

}  // namespace gltrace
}  // namespace android
