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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
)

const TypeIDObservation atom.TypeID = 0xfffe

func init() {
	atom.Register(atom.TypeInfo{ID: TypeIDObservation, New: func() atom.Atom { return &Observation{} }})
}

// Observation is an Atom describing a region of application space memory that
// was observed at capture time.
type Observation struct {
	binary.Generate
	Context    atom.ContextID // The context on which the observation was made.
	Range      Range          // The memory range that was observed.
	ResourceID binary.ID      // The resource identifier holding the memory that was observed.
}

func (a *Observation) String() string {
	r := a.Range
	return fmt.Sprintf("[0x%.16x-0x%.16x] ResID: %s",
		r.First(), r.Last(), a.ResourceID)
}

// Atom compliance
func (a *Observation) TypeID() atom.TypeID       { return TypeIDObservation }
func (a *Observation) ContextID() atom.ContextID { return a.Context }
func (a *Observation) Flags() atom.Flags         { return 0 }
