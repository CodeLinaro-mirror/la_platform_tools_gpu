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

#include "gltrace_context.h"
#include "gltrace.pb.h"
#include "log/log.h"

#include <pthread.h>
#include <time.h>

namespace android {
namespace gltrace {

static pthread_key_t sTLSKey = -1;
static pthread_once_t sPthreadOnceKey = PTHREAD_ONCE_INIT;

void createTLSKey() {
    pthread_key_create(&sTLSKey, (void (*)(void*))&releaseContext);
}

GLTraceContext *getGLTraceContext() {
    GLTraceContext* ctx = (GLTraceContext*) pthread_getspecific(sTLSKey);
    return ctx;
}

void setGLTraceContext(GLTraceContext *c) {
    pthread_setspecific(sTLSKey, c);
}

void setupTraceContextThreadSpecific(GLTraceContext *context) {
    pthread_once(&sPthreadOnceKey, createTLSKey);
    setGLTraceContext(context);
}

void releaseContext() {
    GLTraceContext *c = getGLTraceContext();
    if (c != NULL) {
        delete c;
        setGLTraceContext(NULL);
    }
}

GLTraceState::GLTraceState(TCPStream *stream) : mTraceContextIds(0), mStream(stream) {
}

GLTraceState::~GLTraceState() {
    if (mStream) mStream->closeStream();

    // Clean up state contexts.
    pthread_setspecific(sTLSKey, NULL);
    for (auto it : mPerContextState) {
        delete it.second;
    }
}

TCPStream *GLTraceState::getStream() {
    return mStream;
}

GLTraceContext *GLTraceState::createTraceContext(EGLContext eglContext, gl_hooks_t *hooks) {
    const int id = __sync_fetch_and_add(&mTraceContextIds, 1);
    const size_t DEFAULT_BUFFER_SIZE = 8192;
    BufferedOutputStream *stream = new BufferedOutputStream(mStream, DEFAULT_BUFFER_SIZE);
    GLTraceContext *traceContext = new GLTraceContext(id, this, stream, hooks);

    auto it = mPerContextState.find(eglContext);
    if (it != mPerContextState.end()) {
        delete it->second;
        it->second = traceContext;
    } else {
        mPerContextState.insert({eglContext, traceContext});
    }

    return traceContext;
}

GLTraceContext *GLTraceState::getTraceContext(EGLContext c) {
    auto it = mPerContextState.find(c);
    if (it != mPerContextState.end()) {
        return it->second;
    }
    return NULL;
}

GLTraceContext::GLTraceContext(int id, GLTraceState *state, BufferedOutputStream *stream,
                               gl_hooks_t *hooks) :
        mId(id), mVersionMajor(0), mVersionMinor(0), mVersionParsed(false), mState(state),
        mBufferedOutputStream(stream), mHooks(hooks) {
}

GLTraceContext::~GLTraceContext() {
    for (auto it : mElementArrayBuffers) {
        delete it.second;
    }
}

int GLTraceContext::getId() {
    return mId;
}

int GLTraceContext::getVersionMajor() {
    if (!mVersionParsed) {
        parseGlesVersion();
        mVersionParsed = true;
    }
    return mVersionMajor;
}

int GLTraceContext::getVersionMinor() {
    if (!mVersionParsed) {
        parseGlesVersion();
        mVersionParsed = true;
    }
    return mVersionMinor;
}

GLTraceState *GLTraceContext::getGlobalTraceState() {
    return mState;
}

gl_hooks_t *GLTraceContext::getHooks() {
    return mHooks;
}

void GLTraceContext::parseGlesVersion() {
    const char* str = (const char*)mHooks->gl.glGetString(GL_VERSION);
    int major, minor;
    if (sscanf(str, "OpenGL ES-CM %d.%d", &major, &minor) != 2) {
        if (sscanf(str, "OpenGL ES %d.%d", &major, &minor) != 2) {
            ALOGW("Unable to parse GL_VERSION string: \"%s\"", str);
            major = 1;
            minor = 0;
        }
    }
    mVersionMajor = major;
    mVersionMinor = minor;
}

void GLTraceContext::traceGLMessage(GLMessage *msg) {
    mBufferedOutputStream->send(msg);

    switch (msg->function()) {
        case GLMessage::eglCreateContext:
        case GLMessage::eglMakeCurrent:
        case GLMessage::eglSwapBuffers:
        case GLMessage::glDrawArrays:
        case GLMessage::glDrawElements:
            mBufferedOutputStream->flush();
        default:
            break;
    }
}

void GLTraceContext::bindBuffer(GLuint bufferId, GLvoid *data, GLsizeiptr size) {
    auto it = mElementArrayBuffers.find(bufferId);
    if (it != mElementArrayBuffers.end()) {
        delete it->second;
        it->second = new ElementArrayBuffer(data, size);
    } else {
        mElementArrayBuffers.insert({bufferId, new ElementArrayBuffer(data, size)});
    }
}

void GLTraceContext::getBuffer(GLuint bufferId, GLvoid **data, GLsizeiptr *size) {
    auto it = mElementArrayBuffers.find(bufferId);
    if (it != mElementArrayBuffers.end()) {
        *data = it->second->getBuffer();
        *size = it->second->getSize();
    } else {
        *data = NULL;
        *size = 0;
    }
}

void GLTraceContext::updateBufferSubData(GLuint bufferId, GLintptr offset, GLvoid *data,
                                         GLsizeiptr size) {
    auto it = mElementArrayBuffers.find(bufferId);
    if (it != mElementArrayBuffers.end()) {
        it->second->updateSubBuffer(offset, data, size);
    }
}

void GLTraceContext::deleteBuffer(GLuint bufferId) {
    auto it = mElementArrayBuffers.find(bufferId);
    if (it != mElementArrayBuffers.end()) {
        mElementArrayBuffers.erase(it);
        delete it->second;
    }
}

nsecs_t GLTraceContext::getSystemTime(int clock) {
    static const clockid_t clocks[] = {
            CLOCK_REALTIME,
            CLOCK_MONOTONIC,
            CLOCK_PROCESS_CPUTIME_ID,
            CLOCK_THREAD_CPUTIME_ID,
            CLOCK_BOOTTIME
    };
    struct timespec t;
    t.tv_sec = t.tv_nsec = 0;
    clock_gettime(clocks[clock], &t);
    return nsecs_t(t.tv_sec)*1000000000LL + t.tv_nsec;
}

ElementArrayBuffer::ElementArrayBuffer() : mBuf(NULL), mSize(0) {
}

ElementArrayBuffer::ElementArrayBuffer(GLvoid *buf, GLsizeiptr size) :
        mBuf(malloc(size)), mSize(size) {
    if (buf != NULL) {
        memcpy(mBuf, buf, size);
    }
}

ElementArrayBuffer::~ElementArrayBuffer() {
    if (mBuf != NULL) {
        free(mBuf);
    }
}

void ElementArrayBuffer::updateSubBuffer(GLintptr offset, const GLvoid* data, GLsizeiptr size) {
    if (offset + size <= mSize) {
        memcpy((char*)mBuf + offset, data, size);
    }
}

GLvoid *ElementArrayBuffer::getBuffer() {
    return mBuf;
}

GLsizeiptr ElementArrayBuffer::getSize() {
    return mSize;
}

} // end of namespace gltrace
} // end of namespace android
