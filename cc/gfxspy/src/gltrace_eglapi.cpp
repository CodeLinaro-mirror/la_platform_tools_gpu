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

#include "egldefs.h"
#include "glestrace.h"
#include "gltrace_context.h"
#include "gltrace_egl.h"
#include "gltrace_transport.h"
#include "hooks.h"
#include "Loader.h"
#include "log/log.h"

#include <arpa/inet.h>

namespace android {
namespace gltrace {

static pthread_once_t sGlTraceDriverInitOnce = PTHREAD_ONCE_INIT;
static pthread_mutex_t sGlTraceStateLock = PTHREAD_MUTEX_INITIALIZER;

static int sGlTraceInProgress = 0;
static GLTraceState *sGLTraceState;
static pthread_t sReceiveThreadId;

egl_connection_t gEGLImpl;

// Driver initializer that should be called exactly once, hence the nested struct and pthread_once.
static void initializeDrivers() {
    struct DriverInitializer {
        static void initialize() {
            memset(&gEGLImpl, sizeof(gEGLImpl), 0);
            Loader::open(&gEGLImpl);
            GLTrace_start();
        }
    };
    pthread_once(&sGlTraceDriverInitOnce, &DriverInitializer::initialize);
}

/**
 * Task that monitors the control stream from the host and updates
 * the trace status according to commands received from the host.
 * Note: this is now a no-op placeholder for future gfxspy commands.
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

    while (true) {
        // Read the command size, break if the connection is closed.
        if (stream->receive(&cmdSize, sizeof(uint32_t)) < 0) {
            break;
        }
        cmdSize = ntohl(cmdSize);

        // Grow the command buffer size if necessary.
        if (cmdBufSize < cmdSize) {
            free(cmdBuf);
            cmdBuf = malloc(cmdSize);
            cmdBufSize = cmdSize;
            if (cmdBuf == NULL) break;
        }

        // Read the command, break if the connection is closed.
        if (stream->receive(cmdBuf, cmdSize) < 0) {
            break;
        }
        const uint32_t cmd = ntohl(*(uint32_t*)cmdBuf);

        // Placeholder for gfxspy commands.
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

    pthread_mutex_lock(&sGlTraceStateLock);
    if (sGlTraceInProgress) goto done;

#if defined(GLTRACE_DISABLE_SERVER)
    // Must create a trace state when the server is disabled because the rest of the tracing code
    // expects/needs a GLTraceState object.
    sGLTraceState = new GLTraceState(NULL);
    sGlTraceInProgress = 1;
    goto done;
#else
    // Disabling the server is handy for debugging the tracer because it lets you test
    // loading/running the tracer without needing to connect with a server.
    clientSocket = gltrace::acceptClientConnection(const_cast<char*>("gltrace"));
    if (clientSocket < 0) {
        ALOGE("Error creating GLTrace server socket. Tracing disabled.");
        status = -1;
        goto done;
    }
#endif

    // Initialize the tracing state and create a communication channel to the host.
    sGlTraceInProgress = 1;
    sGLTraceState = new GLTraceState(new TCPStream(clientSocket));
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
    initializeDrivers();
    GLTraceContext *traceContext = state->createTraceContext(c, gEGLImpl.hooks[version]);
    gltrace::setupTraceContextThreadSpecific(traceContext);

    // trace command through to the host
    gltrace::GLTrace_eglCreateContext(version, traceContext->getId());
}

void GLTrace_eglMakeCurrent(EGLContext c) {
    pthread_mutex_lock(&sGlTraceStateLock);
    GLTraceState *state = sGLTraceState;
    pthread_mutex_unlock(&sGlTraceStateLock);

    if (state == NULL) return;

    // setup per context state
    GLTraceContext *traceContext = state->getTraceContext(c);
    if (traceContext == NULL) {
        ALOGD("<gfxspy> GLES version unspecified, defaulting to GLESv1.");
        GLTrace_eglCreateContext(egl_connection_t::GLESv1_INDEX, c);
        traceContext = state->getTraceContext(c);
    }
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

#ifndef GLTRACE_SHOULDNT_LOAD_HOOKS_AT_STARTUP
__attribute__((constructor))
static void GLTrace_init() {
    initializeDrivers();
}
__attribute__((destructor))
static void GLTrace_exit() {
    GLTrace_stop();
}
#endif // GLTRACE_SHOULDNT_LOAD_HOOKS_AT_STARTUP

}  // end of namespace gltrace
}  // end of namespace android
