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

#include "Log.h"
#include "SocketConnection.h"

#include <gapic/target.h>

#include <stdint.h>
#include <string.h>

#if TARGET_OS == GAPID_OS_WINDOWS

#define _WSPIAPI_EMIT_LEGACY
#include <winsock2.h>
#include <ws2tcpip.h>

#else  // TARGET_OS == GAPID_OS_WINDOWS

#include <errno.h>
#include <netdb.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>

#endif  // TARGET_OS == GAPID_OS_WINDOWS

namespace android {
namespace caze {
namespace {

void close(int fd) {
#if TARGET_OS == GAPID_OS_WINDOWS
    ::closesocket(fd);
#else  // TARGET_OS == GAPID_OS_WINDOWS
    ::close(fd);
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

size_t recv(int sockfd, void *buf, size_t len, int flags) {
#if TARGET_OS == GAPID_OS_WINDOWS
    return ::recv(sockfd, static_cast<char*>(buf), static_cast<int>(len), flags);
#else  // TARGET_OS == GAPID_OS_WINDOWS
    return ::recv(sockfd, buf, len, flags);
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

size_t send(int sockfd, const void *buf, size_t len, int flags) {
#if TARGET_OS == GAPID_OS_WINDOWS
    return ::send(sockfd, static_cast<const char*>(buf), static_cast<int>(len), flags);
#else  // TARGET_OS == GAPID_OS_WINDOWS
    return ::send(sockfd, buf, len, flags);
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

int accept(int sockfd, struct sockaddr *addr, size_t *addrlen) {
    if (addrlen == nullptr) {
        return static_cast<int>(::accept(sockfd, addr, nullptr));
    }

#if TARGET_OS == GAPID_OS_WINDOWS
    int addrlenTmp = static_cast<int>(*addrlen);
#else  // TARGET_OS == GAPID_OS_WINDOWS
    socklen_t addrlenTmp = *addrlen;
#endif  // TARGET_OS == GAPID_OS_WINDOWS

    int ret = static_cast<int>(::accept(sockfd, addr, &addrlenTmp));
    *addrlen = addrlenTmp;
    return ret;
}

int getaddrinfo(const char *node, const char *service, const struct addrinfo *hints,
                struct addrinfo **res) {
    return ::getaddrinfo(node, service, hints, res);
}

void freeaddrinfo(struct addrinfo *res) {
    ::freeaddrinfo(res);
}

int setsockopt(int sockfd, int level, int optname, const void *optval, size_t optlen) {
#if TARGET_OS == GAPID_OS_WINDOWS
    return ::setsockopt(sockfd, level, optname, static_cast<const char*>(optval),
            static_cast<int>(optlen));
#else  // TARGET_OS == GAPID_OS_WINDOWS
    return ::setsockopt(sockfd, level, optname, optval, optlen);
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

int socket(int domain, int type, int protocol) {
    return static_cast<int>(::socket(domain, type, protocol));
}

int listen(int sockfd, int backlog) {
    return ::listen(sockfd, backlog);
}

int bind(int sockfd, const struct sockaddr *addr, size_t addrlen) {
#if TARGET_OS == GAPID_OS_WINDOWS
    return ::bind(sockfd, addr, static_cast<int>(addrlen));
#else  // TARGET_OS == GAPID_OS_WINDOWS
    return ::bind(sockfd, addr, addrlen);
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

int error() {
#if TARGET_OS == GAPID_OS_WINDOWS
    return ::WSAGetLastError();
#else  // TARGET_OS == GAPID_OS_WINDOWS
    return errno;
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

}  // end of anonymous namespace

SocketConnection::SocketConnection(int socket) : mSocket(socket) {
}

SocketConnection::~SocketConnection() {
    caze::close(mSocket);
}

size_t SocketConnection::send(const void* data, size_t size) {
    return caze::send(mSocket, data, size, 0);
}

size_t SocketConnection::recv(void* data, size_t size) {
    return caze::recv(mSocket, data, size, MSG_WAITALL);
}

const char* SocketConnection::error() {
    return strerror(caze::error());
}

std::unique_ptr<Connection> SocketConnection::accept() {
    int clientSocket = caze::accept(mSocket, nullptr, nullptr);
    if (-1 == clientSocket) {
        GAPID_WARNING("Failed to accept incoming connection: %s\n", strerror(caze::error()));
        return nullptr;
    }
    return std::unique_ptr<Connection>(new SocketConnection(clientSocket));
}

std::unique_ptr<Connection> SocketConnection::create(const char* hostname, const char* port) {
    // Network initializer to ensure that the network driver is initialized during the lifetime of
    // the create function. If the connection created successfully then the new connection will
    // hold a reference to a networkInitializer struct to ensure that the network is initialized
    NetworkInitializer networkInitializer;

    struct addrinfo* addr;
    struct addrinfo hints;

    memset(&hints, 0, sizeof(addrinfo));
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;

    const int getaddrinfoRes = caze::getaddrinfo(hostname, port, &hints, &addr);
    if (0 != getaddrinfoRes) {
        GAPID_WARNING("getaddrinfo() failed: %d - %s.\n", getaddrinfoRes, strerror(caze::error()));
        return nullptr;
    }
    auto addrDeleter = [](struct addrinfo* ptr) { caze::freeaddrinfo(ptr); };  // deferred.
    std::unique_ptr<struct addrinfo, decltype(addrDeleter)> addrScopeGuard(addr, addrDeleter);

    const int sock = caze::socket(addr->ai_family, addr->ai_socktype, addr->ai_protocol);
    if (-1 == sock) {
        GAPID_WARNING("socket() failed: %s.\n", strerror(caze::error()));
        return nullptr;
    }
    auto socketCloser = [](const int* ptr) { caze::close(*ptr); };  // deferred.
    std::unique_ptr<const int, decltype(socketCloser)> sockScopeGuard(&sock, socketCloser);

    const int one = 1;
    if (-1 == caze::setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, (const char*)&one, sizeof(int))) {
        GAPID_WARNING("setsockopt() failed: %s\n", strerror(caze::error()));
        return nullptr;
    }

    if (-1 == caze::bind(sock, addr->ai_addr, addr->ai_addrlen)) {
        GAPID_WARNING("bind() failed: %s.\n", strerror(caze::error()));
        return nullptr;
    }

    if (-1 == caze::listen(sock, 10)) {
        GAPID_WARNING("listen() failed: %s.\n", strerror(caze::error()));
        return nullptr;
    }

    sockScopeGuard.release();
    return std::unique_ptr<Connection>(new SocketConnection(sock));
}

SocketConnection::NetworkInitializer::NetworkInitializer() {
#if TARGET_OS == GAPID_OS_WINDOWS
    WSADATA wsaData;
    int wsaInitRes = ::WSAStartup(MAKEWORD(2, 2), &wsaData);
    if (wsaInitRes != 0) {
        GAPID_FATAL("WSAStartup failed with error code: %d\n", wsaInitRes);
    }
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

SocketConnection::NetworkInitializer::~NetworkInitializer() {
#if TARGET_OS == GAPID_OS_WINDOWS
    ::WSACleanup();
#endif  // TARGET_OS == GAPID_OS_WINDOWS
}

}  // end of namespace caze
}  // end of namespace android
