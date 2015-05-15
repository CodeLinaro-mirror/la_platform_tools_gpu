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

#include "resource_disk_cache.h"
#include "resource_provider.h"
#include "server_connection.h"

#include <gapic/log.h>

#include <errno.h>
#include <stdio.h>
#include <sys/stat.h>

#include <memory>
#include <string>
#include <utility>
#include <vector>

#if TARGET_OS == GAPID_OS_WINDOWS
#include <direct.h>
#define mkdir(path, mode) _mkdir(path)
#else
static const mode_t MKDIR_MODE = 0755;
#endif

namespace gapir {
namespace {

int mkdirAll(const std::string& path) {
    if (0 != mkdir(path.c_str(), MKDIR_MODE)) {
        switch (errno) {
            case ENOENT: {  // Non-existent parent(s).
                size_t pos = path.find_last_of("/\\");
                if (pos == std::string::npos) {
                    return -1;
                }
                mkdirAll(path.substr(0, pos));
                return mkdir(path.c_str(), MKDIR_MODE);  // Retry.
            }
            case EEXIST:  // Already exists, return success.
                return 0;
            default:  // Something went wrong, return failure.
                return -1;
        }
    }
    return 0;
}

}  // anonymous namespace

std::unique_ptr<ResourceProvider> ResourceDiskCache::create(
        std::unique_ptr<ResourceProvider> fallbackProvider, const std::string& path) {
    if (0 != mkdirAll(path)) {
        GAPID_WARNING("Couldn't access/create cache directory; disabling disk cache.");
        return fallbackProvider;  // Disk path was inaccessible.
    } else {
        std::string diskPath = path;
        if (diskPath.back() != PATH_DELIMITER) {
            diskPath.push_back(PATH_DELIMITER);
        }

        return std::unique_ptr<ResourceProvider>(
                new ResourceDiskCache(std::move(fallbackProvider), std::move(diskPath)));
    }
}

ResourceDiskCache::ResourceDiskCache(std::unique_ptr<ResourceProvider> fallbackProvider,
                                     const std::string& path) :
        mFallbackProvider(std::move(fallbackProvider)), mArchive(path + "resources") {
}

bool ResourceDiskCache::get(const ResourceId& id, const ServerConnection& gazer, void* target,
                            uint32_t size) {
    // Try loading the resource from the archive.
    if (mArchive.read(id, target, size)) {
        return true;
    }

    // Fetch the resource from the fall back provider if it's not in the archive or if its size
    // doesn't match the requested size.
    if (mFallbackProvider->get(id, gazer, target, size)) {
        // Try saving the resource to the archive and return success.
        mArchive.write(id, target, size);
        return true;
    }
    return false;
}

bool ResourceDiskCache::getUncached(const ResourceId& id, const ServerConnection& gazer,
                                   void* target, uint32_t size) {
    return mFallbackProvider->getUncached(id, gazer, target, size);
}

bool ResourceDiskCache::prefetch(const ResourceList& resources,
                                 const ServerConnection& gazer, void* buffer, uint32_t size) {
    size_t querySumSize = 0;
    ResourceList query;

    // Batch the resource requests into batches where the sum size of each batch fit into the size
    // of the buffer available for prefetching
    for (const auto& res : resources) {
        // Check if the resource is already in the archive.
        if (mArchive.contains(res.first)) continue;

        if (res.second > size) {
            GAPID_WARNING("Can't prefetch resource because of limited buffer size");
        } else {
            // If next resource not fit into this batch then fetch the current batch
            if (querySumSize + res.second > size) {
                if (!fetch(gazer, buffer, size, query)) {
                    return false;
                }
                query.clear();
                querySumSize = 0;
            }
            query.push_back(res);
            querySumSize += res.second;
        }
    }

    // Fetch the last batch if it isn't empty
    if (query.size() > 0) {
        return fetch(gazer, buffer, size, query);
    }

    return true;
}

bool ResourceDiskCache::fetch(const ServerConnection& gazer,
                              void* buffer, uint32_t size,
                              const ResourceList& query) {
    // Request the resources from the fall back provider
    if (!mFallbackProvider->get(query, gazer, buffer, size)) {
        return false;
    }

    // Save each resources to the archive from the correct buffer offset.
    for (size_t i = 0, offset = 0; i < query.size(); ++i) {
        mArchive.write(query[i].first, static_cast<uint8_t*>(buffer) + offset, query[i].second);
        offset += query[i].second;
    }
    return true;
}

}  // namespace gapir
