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

#if TARGET_OS == CAZE_OS_ANDROID
const char* listenerPort = "9285";  // Hard coded port for communication with the server
const char* cachePath = "/sdcard/caze_cache";
#else  // TARGET_OS == CAZE_OS_ANDROID
const char* listenerPort = "9284";  // Hard coded port for communication with the server
const char* cachePath = "data/cache";
#endif  // TARGET_OS == CAZE_OS_ANDROID

void listenConnections(MemoryManager* memoryManager) {
    std::unique_ptr<Connection> listenConn = SocketConnection::create("127.0.0.1", listenerPort);
    if (listenConn == nullptr) {
        CAZE_FATAL("Failed to create listening socket\n");
    }
    GazerListener listener(std::move(listenConn));

    std::unique_ptr<ResourceInMemoryCache> resourceProvider(ResourceInMemoryCache::create(
            ResourceDiskCache::create(ResourceRequester::create(), cachePath),
            memoryManager->getBaseAddress()));

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
    ::android::caze::MemoryManager memoryManager({256U * 1024 * 1024});  // 256MB
    ::android::caze::listenConnections(&memoryManager);
}

#else  // TARGET_OS == CAZE_OS_ANDROID
// Main function for PC
int main(int, char* []) {
    ::android::caze::MemoryManager memoryManager({2U * 1024 * 1024 * 1024});  // 2GB
    ::android::caze::listenConnections(&memoryManager);
    return EXIT_SUCCESS;
}

#endif  // TARGET_OS == CAZE_OS_ANDROID
