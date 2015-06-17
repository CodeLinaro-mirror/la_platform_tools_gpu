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
	"android.googlesource.com/platform/tools/gpu/log"
)

// Resource returns a Slice that wraps a resource stored in the database.
// resId is the identifier of the data and size is the size in bytes of the
// data.
func Resource(resId binary.ID, size uint64) Slice {
	return resource{resId, size}
}

type resource struct {
	resId binary.ID
	size  uint64
}

func (r resource) Get(d database.Database, l log.Logger) ([]byte, error) {
	data, err := database.ResolveBlob(r.resId, d, l)
	if err != nil {
		return nil, err
	}
	if r.size != uint64(len(data)) {
		return nil, fmt.Errorf("Loaded resource is unexpected size. Expected 0x%x, got 0x%x for resource %v",
			r.size, len(data), r.resId)
	}
	return data, nil
}

func (r resource) Slice(rng Range) Slice {
	return newResourceSlice(r, rng)
}

func (r resource) ResourceID(d database.Database, l log.Logger) (binary.ID, error) {
	return r.resId, nil
}

func (r resource) Size() uint64 {
	return r.size
}

func (r resource) String() string {
	return fmt.Sprintf("Resource[%v]", r.resId)
}

func newResourceSlice(src Slice, rng Range) Slice {
	if uint64(rng.Last()) > src.Size() {
		panic(fmt.Errorf("Slice range %v out of bounds %v", rng, Range{Base: 0, Size: src.Size()}))
	}
	return resourceSlice{src, rng}
}

type resourceSlice struct {
	src Slice
	rng Range
}

func (s resourceSlice) Get(d database.Database, l log.Logger) ([]byte, error) {
	if data, err := s.src.Get(d, l); err == nil {
		return data[s.rng.First() : s.rng.Last()+1], nil
	} else {
		return nil, err
	}
}

func (s resourceSlice) ResourceID(d database.Database, l log.Logger) (binary.ID, error) {
	panic("resourceSlice.ResourceID currently not implemented")
}

func (s resourceSlice) Slice(rng Range) Slice {
	return newResourceSlice(s.src, rng)
}

func (s resourceSlice) Size() uint64 {
	return s.rng.Size
}

func (s resourceSlice) String() string {
	return fmt.Sprintf("%v[%v]", s.src, s.rng)
}
