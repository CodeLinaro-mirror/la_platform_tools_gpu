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

#include "Connection.h"
#include "Context.h"
#include "GazerConnection.h"
#include "GazerListener.h"
#include "Log.h"
#include "MemoryManager.h"
#include "ResourceDiskCache.h"
#include "ResourceInMemoryCache.h"
#include "ResourceRequester.h"
#include "SocketConnection.h"
#include "Target.h"

#include <memory>
#include <stdlib.h>

#if TARGET_OS == CAZE_OS_ANDROID
#include <android_native_app_glue.h>
#endif  // TARGET_OS == CAZE_OS_ANDROID

namespace android {
namespace caze {
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
    std::unique_ptr<Connection> listenConn = SocketConnection::create("127.0.0.1", listenerPort);
    if (listenConn == nullptr) {
        CAZE_FATAL("Failed to create listening socket\n");
    }
    GazerListener listener(std::move(listenConn), memoryManager->getSize());

    std::unique_ptr<ResourceInMemoryCache> resourceProvider(
            createResourceProvider(cachePath, memoryManager));

    while (true) {
        std::unique_ptr<GazerConnection> gazer(listener.acceptConnection());
        if (!gazer) {
            break;
        }

        std::unique_ptr<Context> context =
                Context::create(*gazer, resourceProvider.get(), memoryManager);
        if (context == nullptr) {
            CAZE_WARNING("Loading Context failed!\n");
            continue;
        }

        resourceProvider->updateSize(context->getInMemoryCacheSize());
        context->interpret();
    }
}

}  // end of anonymous namespace
}  // end of namespace caze
}  // end of namespace android

#if TARGET_OS == CAZE_OS_ANDROID
// Main function for android
void android_main(struct android_app*) {
    app_dummy();
    ::android::caze::MemoryManager memoryManager(::android::caze::memorySizes);
    ::android::caze::listenConnections("9285", "/sdcard/caze_cache", &memoryManager);
}

#else  // TARGET_OS == CAZE_OS_ANDROID
// Main function for PC
int main(int argc, char* argv[]) {
    bool useCache = true;
    for (int i = 1; i < argc; i++) {
        if (strstr(argv[i], "--nocache") == argv[i]) {
            printf("Disabling cache\n");
            useCache = false;
        }
    }
    const char* cachePath = useCache ? ("data" PATH_DELIMITER_STR "ccache") : nullptr;
    ::android::caze::MemoryManager memoryManager(::android::caze::memorySizes);
    ::android::caze::listenConnections("9284", cachePath, &memoryManager);
    return EXIT_SUCCESS;
}

#endif  // TARGET_OS == CAZE_OS_ANDROID
