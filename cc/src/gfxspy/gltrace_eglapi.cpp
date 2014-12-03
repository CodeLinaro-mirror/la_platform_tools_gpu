/*
 * Copyright 2011, The Android Open Source Project
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

#include <arpa/inet.h>
#include <stdlib.h>
#include <cutils/log.h>

#include "hooks.h"
#include "glestrace.h"

#include "gltrace_context.h"
#include "gltrace_egl.h"
#include "gltrace_hooks.h"
#include "gltrace_transport.h"

#include "EGL/Loader.h"

#include "EGL/egldefs.h"
#include "EGL/egl_object.h"

namespace android {
namespace gltrace {

static pthread_mutex_t sGlTraceStateLock = PTHREAD_MUTEX_INITIALIZER;

static int sGlTraceInProgress = 0;
static GLTraceState *sGLTraceState;
static pthread_t sReceiveThreadId;

egl_connection_t gEGLImpl;
::android::gl_hooks_t gHooks[2];

/**
 * Task that monitors the control stream from the host and updates
 * the trace status according to commands received from the host.
 */
static void *commandReceiveTask(void *arg) {
    GLTraceState *state = (GLTraceState *)arg;
    TCPStream *stream = state->getStream();

    // The control stream always receives an integer size of the
    // command buffer, followed by the actual command buffer.
    uint32_t cmdSize;

    // Command Buffer
    void *cmdBuf = NULL;
    uint32_t cmdBufSize = 0;

    enum TraceSettingsMasks {
        READ_TEXTURE_DATA_ON_GLTEXIMAGE_MASK = 1 << 2,
    };

    while (true) {
        // read command size
        if (stream->receive(&cmdSize, sizeof(uint32_t)) < 0) {
            break;
        }
        cmdSize = ntohl(cmdSize);

        // ensure command buffer is of required size
        if (cmdBufSize < cmdSize) {
            free(cmdBuf);
            cmdBufSize = cmdSize;
            cmdBuf = malloc(cmdSize);
            if (cmdBuf == NULL)
                break;
        }

        // receive the command
        if (stream->receive(cmdBuf, cmdSize) < 0) {
            break;
        }

        if (cmdSize != sizeof(uint32_t)) {
            // Currently, we only support commands that are a single integer,
            // so we skip all other commands
            continue;
        }

        uint32_t cmd = ntohl(*(uint32_t*)cmdBuf);

        bool collectTextureData = (cmd & READ_TEXTURE_DATA_ON_GLTEXIMAGE_MASK) != 0;

        state->setCollectTextureDataOnGlTexImage(collectTextureData);

        ALOGD("trace options: texImage: %d", collectTextureData);
    }

    ALOGE("Stopping OpenGL Trace Command Receiver\n");

    free(cmdBuf);
    return NULL;
}

/**
 * Starts Trace Server and waits for connection from the host.
 * Returns -1 in case of connection error, 0 otherwise.
 */
int GLTrace_start() {
    int status = 0;
    int clientSocket = -1;
    TCPStream *stream = NULL;

    pthread_mutex_lock(&sGlTraceStateLock);

    if (sGlTraceInProgress) {
        goto done;
    }

    ALOGD("GLTrace_start().");
    memset(&gEGLImpl, sizeof(gEGLImpl), 0);
    memset(gHooks, sizeof(gHooks), 0);

#if !defined(GLTRACE_DISABLE_SERVER)
    // Disabling the server is handy for debugging the tracer
    // because it lets you test loading/running the tracer
    // without needint to connect with a server.
    clientSocket = gltrace::acceptClientConnection(const_cast<char*>("gltrace"));
#endif

    if (clientSocket < 0) {
        ALOGE("Error creating GLTrace server socket. Tracing disabled.");
        status = -1;
#if defined(GLTRACE_DISABLE_SERVER)
        // Must create a trace state when the server is disabled
        // because the rest of the tracing code expects/needs
        // a GLTraceState object.
        sGLTraceState = new GLTraceState(NULL);
        sGlTraceInProgress = 1;
#endif
        goto done;
    }
    sGlTraceInProgress = 1;

    // create communication channel to the host
    stream = new TCPStream(clientSocket);

    // initialize tracing state
    sGLTraceState = new GLTraceState(stream);

    pthread_create(&sReceiveThreadId, NULL, commandReceiveTask, sGLTraceState);

done:
    pthread_mutex_unlock(&sGlTraceStateLock);
    return status;
}

void GLTrace_stop() {
    pthread_mutex_lock(&sGlTraceStateLock);

    if (sGlTraceInProgress) {
        sGlTraceInProgress = 0;
        delete sGLTraceState;
        sGLTraceState = NULL;
    }

    pthread_mutex_unlock(&sGlTraceStateLock);
}

void GLTrace_eglCreateContext(int version, EGLContext c) {
    pthread_mutex_lock(&sGlTraceStateLock);
    GLTraceState *state = sGLTraceState;
    pthread_mutex_unlock(&sGlTraceStateLock);

    if (state == NULL) return;

    // update trace state for new EGL context
    GLTraceContext *traceContext = state->createTraceContext(version, c);
    gltrace::setupTraceContextThreadSpecific(traceContext);

    // trace command through to the host
    gltrace::GLTrace_eglCreateContext(version, traceContext->getId());
}

// egl_init_drivers_locked() and egl_init_drivers() were copied
// (with minor modifications) from frameworks/native/opengl/libs/EGL/egl.cpp.
static EGLBoolean egl_init_drivers_locked() {
    //ALOGD("egl_init_drivers_locked()");
    // get our driver loader
    Loader& loader(Loader::getInstance());

    // dynamically load our EGL implementation
    egl_connection_t* cnx = &gEGLImpl;
    if (cnx->dso == NULL) {
        cnx->hooks[egl_connection_t::GLESv1_INDEX] = &gHooks[
            egl_connection_t::GLESv1_INDEX];
        cnx->hooks[egl_connection_t::GLESv2_INDEX] = &gHooks[
            egl_connection_t::GLESv2_INDEX];
        cnx->dso = loader.open(cnx);
    }

    return cnx->dso ? EGL_TRUE : EGL_FALSE;
}

// This mutex protects egl_init_drivers_locked().
static pthread_mutex_t sInitDriverMutex = PTHREAD_MUTEX_INITIALIZER;

EGLBoolean egl_init_drivers() {
    GLTrace_start();
    EGLBoolean res;
    pthread_mutex_lock(&sInitDriverMutex);
    res = egl_init_drivers_locked();
    pthread_mutex_unlock(&sInitDriverMutex);
    return res;
}

gl_hooks_t *loadHooks() {
    egl_init_drivers();
    return gEGLImpl.hooks[egl_connection_t::GLESv2_INDEX];
}

#if defined(GLTRACE_LOAD_HOOKS_AT_STARTUP)
// This function will be called automatically when this
// library (libgfxspy.so) is loaded.  This allows the hooks to
// be inserted before other libraries are loaded, and is necessary
// when tracing is done by pre-loading the gfxspy library using
// LD_PRELOAD.
__attribute__((constructor)) static void GLTrace_init() {
    ALOGD("GLTrace_init() called.");
    memset(&gEGLImpl, sizeof(gEGLImpl), 0);
    memset(gHooks, sizeof(gHooks), 0);

    egl_init_drivers();
}
#endif  // defined(GLTRACE_LOAD_HOOKS_AT_STARTUP)

void GLTrace_eglMakeCurrent(const unsigned version, gl_hooks_t *hooks, EGLContext c) {
    pthread_mutex_lock(&sGlTraceStateLock);
    GLTraceState *state = sGLTraceState;
    pthread_mutex_unlock(&sGlTraceStateLock);

    if (state == NULL) return;

    // setup per context state
    GLTraceContext *traceContext = state->getTraceContext(c);
    if (traceContext == NULL) {
        GLTrace_eglCreateContext(version, c);
        traceContext = state->getTraceContext(c);
    }

    if (hooks == NULL) {
        hooks = loadHooks();
    }

    traceContext->hooks = hooks;
    gltrace::setupTraceContextThreadSpecific(traceContext);

    // trace command through to the host
    gltrace::GLTrace_eglMakeCurrent(traceContext->getId());
}

void GLTrace_eglReleaseThread() {
    gltrace::releaseContext();
}

void GLTrace_eglSwapBuffers_internal(void *dpy, void *draw) {
    gltrace::GLTrace_eglSwapBuffers(dpy, draw);
}

gl_hooks_t *GLTrace_getGLHooks() {
    return gltrace::getGLHooks();
}

}  // namespace gltrace
}  // namespace android
