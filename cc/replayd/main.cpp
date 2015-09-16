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

#include <gapir/context.h>
#include <gapir/memory_manager.h>
#include <gapir/resource_disk_cache.h>
#include <gapir/resource_in_memory_cache.h>
#include <gapir/resource_requester.h>
#include <gapir/server_connection.h>
#include <gapir/server_listener.h>

#include <gapic/connection.h>
#include <gapic/log.h>
#include <gapic/socket_connection.h>
#include <gapic/target.h>

#include <memory>
#include <stdlib.h>
#include <string.h>

#if TARGET_OS == GAPID_OS_ANDROID
#include <android_native_app_glue.h>
#endif  // TARGET_OS == GAPID_OS_ANDROID

using namespace gapic;
using namespace gapir;

namespace {

std::vector<uint32_t> memorySizes{
    2 * 1024 * 1024 * 1024U,  // 2GB
    1 * 1024 * 1024 * 1024U,  // 1GB
         512 * 1024 * 1024U,  // 512MB
         256 * 1024 * 1024U,  // 256MB
         128 * 1024 * 1024U,  // 128MB
};

// createResourceProvider constructs and returns a ResourceInMemoryCache.
// If cachePath is non-null then the ResourceInMemoryCache will be backed by a
// disk-cache.
std::unique_ptr<ResourceInMemoryCache> createResourceProvider(
        const char* cachePath, MemoryManager* memoryManager) {
    if (cachePath != nullptr) {
        return std::unique_ptr<ResourceInMemoryCache>(
            ResourceInMemoryCache::create(
                ResourceDiskCache::create(ResourceRequester::create(), cachePath),
                memoryManager->getBaseAddress()));
    } else {
        return std::unique_ptr<ResourceInMemoryCache>(
            ResourceInMemoryCache::create(
                ResourceRequester::create(), memoryManager->getBaseAddress()));
    }
}

void listenConnections(const char* listenerPort, const char* cachePath,
                       MemoryManager* memoryManager) {
    std::unique_ptr<Connection> conn = SocketConnection::createSocket("127.0.0.1", listenerPort);
    if (conn == nullptr) {
        GAPID_FATAL("Failed to create listening socket");
    }
    ServerListener listener(std::move(conn), memoryManager->getSize());

    std::unique_ptr<ResourceInMemoryCache> resourceProvider(
            createResourceProvider(cachePath, memoryManager));

    while (true) {
        std::unique_ptr<ServerConnection> acceptedConn(listener.acceptConnection());
        if (!acceptedConn) {
            GAPID_INFO("Shutting down");
            break;
        }

        std::unique_ptr<Context> context =
                Context::create(*acceptedConn, resourceProvider.get(), memoryManager);
        if (context == nullptr) {
            GAPID_WARNING("Loading Context failed!");
            continue;
        }

        resourceProvider->updateSize(context->getInMemoryCacheSize());
        if (!context->interpret()) {
          GAPID_DEBUG("Interpret returned false");
        }
    }
}

}  // anonymous namespace

#if TARGET_OS == GAPID_OS_ANDROID
// Main function for android
void android_main(struct android_app*) {
    app_dummy();
    MemoryManager memoryManager(memorySizes);
    listenConnections("9285", "/sdcard/gapir_cache", &memoryManager);
}

#else  // TARGET_OS == GAPID_OS_ANDROID
// Main function for PC
int main(int argc, char* argv[]) {
    GAPID_LOGGER_INIT("logs/gapir.log");

    bool useCache = true;
    const char* portStr = "9284";
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--nocache") == 0) {
            printf("Disabling cache\n");
            useCache = false;
        }
        if (strcmp(argv[i], "--port") == 0) {
            if (i + 1 >= argc) {
                fprintf(stderr, "Usage: --port <port_num>");
                exit(1);
            }
            portStr = argv[i + 1];
        }
    }
    const char* cachePath = useCache ? ("data" PATH_DELIMITER_STR "ccache") : nullptr;
    MemoryManager memoryManager(memorySizes);
    GAPID_INFO("gapir listening on port %s", portStr);
    listenConnections(portStr, cachePath, &memoryManager);
    return EXIT_SUCCESS;
}

#endif  // TARGET_OS == GAPID_OS_ANDROID
