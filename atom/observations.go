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

package atom

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// Observations is a collection of reads and write observations performed by an
// atom.
type Observations struct {
	binary.Generate
	Reads  []Observation
	Writes []Observation
}

func (o Observations) String() string {
	return fmt.Sprintf("Reads: %v, Writes: %v", o.Reads, o.Writes)
}

// ApplyReads applies all the observed reads to memory pool p.
func (o Observations) ApplyReads(p *memory.Pool) {
	for _, r := range o.Reads {
		p.Write(r.Range.Base, memory.Resource(r.ID, r.Range.Size))
	}
}

// ApplyReads applies all the observed writes to the memory pool p.
func (o Observations) ApplyWrites(p *memory.Pool) {
	for _, w := range o.Writes {
		p.Write(w.Range.Base, memory.Resource(w.ID, w.Range.Size))
	}
}

// Observation represents a single read or write observation made by an atom.
type Observation struct {
	binary.Generate
	Range memory.Range // Memory range that was observed.
	ID    binary.ID    // The resource identifier of the observed data.
}

func (o Observation) String() string {
	return fmt.Sprintf("Range: %v, ID: %v", o.Range, o.ID)
}
