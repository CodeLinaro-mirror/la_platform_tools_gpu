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

package builder

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns the *service.MemoryInfo resulting from the given
// GetMemoryInfo request.
func (r *GetMemoryInfo) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	atoms, err := ResolveAtoms(r.After.Atoms, d, l)
	if err != nil {
		return nil, err
	}

	if r.After.Index >= uint64(len(atoms)) {
		return nil, fmt.Errorf("After (%d) parameter is out of bounds. [0-%d]", r.After, len(atoms)-1)
	}

	s := gfxapi.NewState()
	pool := s.Memory[memory.ApplicationPool]

	for _, a := range atoms[:r.After.Index] {
		a.Mutate(s, d, l)
	}

	var reads, writes memory.RangeList
	pool.OnRead = func(rng memory.Range) {
		if rng.Overlaps(r.Range) {
			interval.Merge(&reads, rng.Window(r.Range).Span(), false)
		}
	}
	pool.OnWrite = func(rng memory.Range) {
		if rng.Overlaps(r.Range) {
			interval.Merge(&writes, rng.Window(r.Range).Span(), false)
		}
	}
	atoms[r.After.Index].Mutate(s, d, l)

	slice := pool.Slice(r.Range)
	data, err := slice.Get(d, l)
	if err != nil {
		return nil, err
	}

	observed := slice.ValidRanges()

	return &service.MemoryInfo{
		Data:     data,
		Reads:    reads,
		Writes:   writes,
		Observed: observed,
	}, nil
}
