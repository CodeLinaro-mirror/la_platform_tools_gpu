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

package gfxapi

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// State represents the graphics state across all contexts.
type State struct {
	binary.Object

	// Architecture holds information about the device architecture that was used
	// to create the capture.
	Architecture device.Architecture

	// Memory holds the memory state of the application.
	Memory map[memory.PoolID]*memory.Pool

	// NextPoolID hold the identifier of the next Pool to be created.
	NextPoolID memory.PoolID

	// APIs holds the per-API context states.
	APIs map[API]binary.Object
}

func (s State) String() string {
	mem := make([]string, 0, len(s.Memory))
	for i, p := range s.Memory {
		mem = append(mem, fmt.Sprintf("    %d: %v", i, strings.Replace(p.String(), "\n", "\n      ", -1)))
	}
	apis := make([]string, 0, len(s.APIs))
	for a, s := range s.APIs {
		apis = append(apis, fmt.Sprintf("    %v: %v", a, s))
	}
	return fmt.Sprintf("State{\n  %v\n  Memory:\n%v\n  APIs:\n%v\n}",
		s.Architecture, strings.Join(mem, "\n"), strings.Join(apis, "\n"))
}
func NewState() *State {
	return &State{
		Architecture: device.Architecture{
			PointerAlignment: 4,
			PointerSize:      4,
			IntegerSize:      4,
			ByteOrder:        endian.Little,
		},
		Memory: map[memory.PoolID]*memory.Pool{
			memory.ApplicationPool: &memory.Pool{},
		},
		NextPoolID: memory.ApplicationPool + 1,
	}
}

// MemoryDecoder returns a flat decoder backed by an endian reader that uses
// the byte-order of the capture device to decode from the slice s.
func (st State) MemoryDecoder(s memory.Slice, d database.Database, l log.Logger) binary.Decoder {
	br := memory.Reader(s, d, l)
	er := endian.Reader(br, st.Architecture.ByteOrder)
	return flat.Decoder(er)
}

// MemoryEncoder returns a flat encoder backed by an endian reader that uses
// the byte-order of the capture device to encode to the pool p, for the range
// rng.
func (st State) MemoryEncoder(p *memory.Pool, rng memory.Range) binary.Encoder {
	bw := memory.Writer(p, rng)
	ew := endian.Writer(bw, st.Architecture.ByteOrder)
	return flat.Encoder(ew)
}
