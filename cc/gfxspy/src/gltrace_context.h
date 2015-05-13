/*
 * Copyright (C) 2015 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef ANDROID_GLTRACE_CONTEXT_H
#define ANDROID_GLTRACE_CONTEXT_H

#include "gltrace_transport.h"
#include "hooks.h"

#include <pthread.h>
#include <stdint.h>
#include <unordered_map>

namespace android {
namespace gltrace {

class GLTraceState;

// From platform/system/core/include/utils/Timers.h
typedef int64_t nsecs_t;
enum {
    SYSTEM_TIME_REALTIME = 0,  // system-wide realtime clock
    SYSTEM_TIME_MONOTONIC = 1, // monotonic time since unspecified starting point
    SYSTEM_TIME_PROCESS = 2,   // high-resolution per-process clock
    SYSTEM_TIME_THREAD = 3,    // high-resolution per-thread clock
    SYSTEM_TIME_BOOTTIME = 4   // same as SYSTEM_TIME_MONOTONIC, but including CPU suspend time
};

class ElementArrayBuffer {
public:
    ElementArrayBuffer();
    ElementArrayBuffer(GLvoid *buf, GLsizeiptr size);
    ~ElementArrayBuffer();

    void updateSubBuffer(GLintptr offset, const GLvoid* data, GLsizeiptr size);
    GLvoid *getBuffer();
    GLsizeiptr getSize();

private:
    GLvoid *mBuf;
    GLsizeiptr mSize;
};

/** GL Trace Context info associated with each EGLContext */
class GLTraceContext {
public:
    GLTraceContext(int id, GLTraceState *state, BufferedOutputStream *stream, gl_hooks_t *hooks);
    ~GLTraceContext();
    int getId();
    int getVersionMajor();
    int getVersionMinor();
    GLTraceState *getGlobalTraceState();
    gl_hooks_t *getHooks();
    nsecs_t getSystemTime(int clock);

    // Methods to work with element array buffers
    void bindBuffer(GLuint bufferId, GLvoid *data, GLsizeiptr size);
    void getBuffer(GLuint bufferId, GLvoid **data, GLsizeiptr *size);
    void updateBufferSubData(GLuint bufferId, GLintptr offset, GLvoid *data, GLsizeiptr size);
    void deleteBuffer(GLuint bufferId);

    void traceGLMessage(GLMessage *msg);

private:
    /* Parses the GL version string returned from glGetString(GL_VERSION) to get find the major and
       minor versions of the GLES API. The context must be current before calling. */
    void parseGlesVersion();

    int mId;                    /* unique context id */
    int mVersionMajor;          /* GL major version. Lazily parsed in getVersionX(). */
    int mVersionMinor;          /* GL minor version. Lazily parsed in getVersionX(). */
    bool mVersionParsed;        /* True if major and minor versions have been parsed. */
    GLTraceState *mState;       /* parent GL Trace state (for per process GL Trace State Info) */
    BufferedOutputStream *mBufferedOutputStream; /* stream where trace info is sent */
    gl_hooks_t *mHooks;

    /* list of element array buffers in use. */
    std::tr1::unordered_map<GLuint, ElementArrayBuffer*> mElementArrayBuffers;
};

/** Per process trace state. */
class GLTraceState {
public:
    GLTraceState(TCPStream *stream);
    ~GLTraceState();

    GLTraceContext *createTraceContext(EGLContext c, gl_hooks_t *hooks);
    GLTraceContext *getTraceContext(EGLContext c);
    TCPStream *getStream();

private:
    int mTraceContextIds;
    TCPStream *mStream;
    std::tr1::unordered_map<EGLContext, GLTraceContext*> mPerContextState;
};

void setupTraceContextThreadSpecific(GLTraceContext *context);
GLTraceContext *getGLTraceContext();
void releaseContext();

} // end of namespace gltrace
} // end of namespace android

#endif // ANDROID_GLTRACE_CONTEXT_H
