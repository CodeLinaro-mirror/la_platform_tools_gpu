// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memory

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
)

type resourceData struct {
	resId binary.ID
	size  uint64
}

func (r resourceData) Get(db database.Database, logger log.Logger) ([]byte, error) {
	binary := store.Blob{}
	if err := db.Load(r.resId, logger, &binary); err != nil {
		return nil, err
	}
	if r.size != uint64(len(binary.Data)) {
		return nil, fmt.Errorf("Loaded resource is unexpected size. Expected 0x%x, got 0x%x for resource %v",
			r.size, len(binary.Data), r.resId)
	}
	return binary.Data, nil
}

func (r resourceData) Slice(rng Range) DataSlicer {
	return slice(r, rng)
}

func (r resourceData) Size() uint64 {
	return r.size
}

// ResourceData returns a DataSlicer that wraps a resource. resId is the
// identifier of the resource and size is the size in bytes of the resource.
func ResourceData(resId binary.ID, size uint64) DataSlicer {
	return resourceData{resId, size}
}
