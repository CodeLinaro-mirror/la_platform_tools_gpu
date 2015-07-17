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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// ResolveMemoryRange returns the MemoryInfo describing the range of memory r in
// pool p immediately following the atom i in the list a.
func ResolveMemoryRange(
	a []atom.Atom,
	i atom.ID,
	p memory.PoolID,
	r memory.Range,
	d database.Database,
	l log.Logger) (*service.MemoryInfo, error) {

	s := gfxapi.NewState()

	for _, a := range a[:i] {
		a.Mutate(s, d, l)
	}

	pool, ok := s.Memory[p]
	if !ok {
		return nil, fmt.Errorf("Pool %d not found", p)
	}

	var reads, writes memory.RangeList
	pool.OnRead = func(rng memory.Range) {
		if rng.Overlaps(r) {
			interval.Merge(&reads, rng.Window(r).Span(), false)
		}
	}
	pool.OnWrite = func(rng memory.Range) {
		if rng.Overlaps(r) {
			interval.Merge(&writes, rng.Window(r).Span(), false)
		}
	}
	a[i].Mutate(s, d, l)

	slice := pool.Slice(r)
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
