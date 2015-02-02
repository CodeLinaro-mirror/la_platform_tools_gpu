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
#include "GazerConnection.h"
#include "ResourceDiskCache.h"
#include "ResourceProvider.h"

#include <errno.h>
#include <stdio.h>
#include <sys/stat.h>

#include <memory>
#include <string>
#include <utility>
#include <vector>

#if TARGET_OS == CAZE_OS_WINDOWS
#include <direct.h>
#define mkdir(path, mode) _mkdir(path)
#else
static const mode_t MKDIR_MODE = 0755;
#endif

namespace android {
namespace caze {
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

}  // end of anonymous namespace

std::unique_ptr<ResourceProvider> ResourceDiskCache::create(
        std::unique_ptr<ResourceProvider> fallbackProvider, const std::string& path) {
    if (0 != mkdirAll(path)) {
        CAZE_WARNING("Couldn't access/create cache directory; disabling disk cache.");
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
        mFallbackProvider(std::move(fallbackProvider)), mPath(path) {
}

bool ResourceDiskCache::get(const ResourceId& id, const GazerConnection& gazer, void* target,
                            uint32_t size) {
    const std::string filepath(mPath + id);
    if (FILE* fi = fopen(filepath.c_str(), "rb")) {
        // Resource already in the disk cache. Load it from disk
        const size_t length = fread(target, 1, size, fi);
        fclose(fi);
        // Check the amount of data read from the disk against the amount of data requested
        if (length == size) {
            return true;
        }
    }

    // Fetch the resource from the fall back provider if it isn't in the disk cache or the size
    // doesn't match with the requested size (mostly corrupted writes)
    if (mFallbackProvider->get(id, gazer, target, size)) {
        // Save the resource to the disk if possible
        if (FILE* fo = fopen(filepath.c_str(), "wb")) {
            fwrite(target, size, 1, fo);
            fclose(fo);
        }
        return true;
    } else {
        return false;
    }
}

bool ResourceDiskCache::prefetch(const ResourceList& resources,
                                 const GazerConnection& gazer, void* buffer, uint32_t size) {
    size_t querySumSize = 0;
    ResourceList query;

    // Batch the resource requests into batches where the sum size of each batch fit into the size
    // of the buffer available for prefetching
    for (const auto& res : resources) {
        const std::string filepath(mPath + res.first);
        // Check if the resource is already in the disk cache
        if (FILE* fi = fopen(filepath.c_str(), "rb")) {
            fclose(fi);
        } else if (res.second > size) {
            CAZE_WARNING("Can't prefetch resource because of limited buffer size");
        } else {
            // If next resource not fit into this batch then fetch the current batch
            if (querySumSize + res.second > size) {
                if (!fetch(gazer, buffer, query)) {
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
        return fetch(gazer, buffer, query);
    }

    return true;
}

bool ResourceDiskCache::fetch(const GazerConnection& gazer, void* buffer,
                              const ResourceList& query) {
    // Request the resources from the fall back provider
    if (!mFallbackProvider->get(query, gazer, buffer)) {
        return false;
    }

    // Save each resources to the disk from the correct offset
    for (size_t i = 0, offset = 0; i < query.size(); ++i) {
        const std::string filepath(mPath + query[i].first);
        if (FILE* fo = fopen(filepath.c_str(), "wb")) {
            fwrite(static_cast<uint8_t*>(buffer) + offset, query[i].second, 1, fo);
            fclose(fo);
        }
        offset += query[i].second;
    }
    return true;
}

}  // end of namespace caze
}  // end of namespace android
