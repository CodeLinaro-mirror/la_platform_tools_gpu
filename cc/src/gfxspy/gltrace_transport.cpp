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

#include "gltrace_transport.h"
#include "gltrace.pb.h"
#include "log/log.h"

#include <errno.h>
#include <netinet/in.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <unistd.h>

namespace android {
namespace gltrace {

// Accepts one client connection on the UNIX socket sockname. Returns its socket, or -1 on failure.
int acceptClientConnection(char *sockname) {
    int serverSocket = socket(AF_LOCAL, SOCK_STREAM, 0);
    if (serverSocket < 0) {
        ALOGE("Error while creating socket: %s. Check if app has network permissions.",
              strerror(errno));
        return -1;
    }

    struct sockaddr_un server, client;

    memset(&server, 0, sizeof server);
    server.sun_family = AF_UNIX;
    // the first byte of sun_path should be '\0' for abstract namespace
    strncpy(server.sun_path + 1, sockname, sizeof(server.sun_path) - 2);

    // note that sockaddr_len should be set to the exact size of the buffer that is used.
    socklen_t sockaddr_len = sizeof(server);
    if (bind(serverSocket, (struct sockaddr *) &server, sockaddr_len) < 0) {
        ALOGE("Failed to bind the server socket: %s", strerror(errno));
        close(serverSocket);
        return -1;
    }

    if (listen(serverSocket, 1) < 0) {
        ALOGE("Failed to listen on server socket: %s", strerror(errno));
        close(serverSocket);
        return -1;
    }

    ALOGD("gltrace::acceptClientConnection: server listening @ path %s", sockname);
    sockaddr_len = sizeof(client);
    memset(&client, 0, sizeof(client));
    int clientSocket = accept(serverSocket, (struct sockaddr *)&client, &sockaddr_len);
    if (clientSocket < 0) {
        ALOGE("Failed to accept client connection: %s", strerror(errno));
        close(serverSocket);
        return -1;
    }

    struct ucred cr;
    socklen_t cr_len = sizeof(cr);
    if (getsockopt(clientSocket, SOL_SOCKET, SO_PEERCRED, &cr, &cr_len) != 0) {
        ALOGE("Error obtaining credentials of peer: %s", strerror(errno));
        close(clientSocket);
        close(serverSocket);
        return -1;
    }

    // Constants from system/core/include/private/android_filesystem_config.h.
    // According to android_filesystem_config.h, these constants should never change.
    static const int AID_ROOT = 0;      /* traditional unix root user */
    static const int AID_SHELL = 2000;  /* adb and debug shell user */

    // Only accept connects from the shell (adb forward comes to us as shell user),
    // or the root user.
    if (cr.uid != AID_SHELL && cr.uid != AID_ROOT) {
        ALOGE("Unknown peer type (%d), expected shell to be the peer", cr.uid);
        close(clientSocket);
        close(serverSocket);
        return -1;
    }
    ALOGD("gltrace::acceptClientConnection: client connected.");

    // do not accept any more incoming connections
    close(serverSocket);

    return clientSocket;
}

TCPStream::TCPStream(int socket) : mSocket(socket) {
    pthread_mutex_init(&mSocketWriteMutex, NULL);
}

TCPStream::~TCPStream() {
    pthread_mutex_destroy(&mSocketWriteMutex);
}

void TCPStream::closeStream() {
    if (mSocket >= 0) {
        close(mSocket);
        mSocket = -1;
    }
}

int TCPStream::send(void *buf, size_t len) {
    if (mSocket < 0) {
        return -1;
    }

    pthread_mutex_lock(&mSocketWriteMutex);
    const int n = write(mSocket, buf, len);
    pthread_mutex_unlock(&mSocketWriteMutex);
    if (n < 0) {
        ALOGE("Error sending data to stream: %s", strerror(errno));
    }

    return n;
}

int TCPStream::receive(void *data, size_t len) {
    if (mSocket < 0) {
        return -1;
    }

    size_t totalRead = 0;
    while (totalRead < len) {
        const int n = read(mSocket, (uint8_t*)data + totalRead, len - totalRead);
        if (n < 0) {
            ALOGE("Error receiving data from stream: %s", strerror(errno));
            return -1;
        }

        totalRead += n;
    }

    return 0;
}

BufferedOutputStream::BufferedOutputStream(TCPStream *stream, size_t bufferSize) :
        mStream(stream), mBufferSize(bufferSize) {
    mStringBuffer.reserve(mBufferSize);
}

int BufferedOutputStream::flush() {
    if (mStringBuffer.size() == 0 || mStream == NULL) {
        // Clear the data when we're not streaming it so the buffer
        // doesn't continue to grow.
        mStringBuffer.clear();
        return 0;
    }

    int n = mStream->send((void *)mStringBuffer.data(), mStringBuffer.size());
    mStringBuffer.clear();
    return n;
}

void BufferedOutputStream::enqueueMessage(GLMessage *msg) {
    const uint32_t len = msg->ByteSize();

    // Send the message size, then the message.
    mStringBuffer.append((const char *)&len, sizeof(len));
    msg->AppendToString(&mStringBuffer);
}

int BufferedOutputStream::send(GLMessage *msg) {
    enqueueMessage(msg);

    if (mStringBuffer.size() > mBufferSize) {
        return flush();
    }

    return 0;
}

}  // end of namespace gltrace
}  // end of namespace android
