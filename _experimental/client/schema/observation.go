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

package schema

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type Observations struct {
	Reads  []Observation
	Writes []Observation
}

// Decode decodes an Observations structure from the decode d.
func (o *Observations) Decode(d binary.Decoder) error {
	if count, err := d.Uint32(); err == nil {
		o.Reads = make([]Observation, count)
		for i := range o.Reads {
			if err := o.Reads[i].Decode(d); err != nil {
				return err
			}
		}
	} else {
		return err
	}

	if count, err := d.Uint32(); err == nil {
		o.Writes = make([]Observation, count)
		for i := range o.Writes {
			if err := o.Writes[i].Decode(d); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

type Observation struct {
	Range memory.Range // Memory range that was observed.
	ID    binary.ID    // The resource identifier of the observed data.
}

// Decode decodes an Observations structure from the decode d.
func (o *Observation) Decode(d binary.Decoder) error {
	if base, err := d.Uint64(); err == nil {
		o.Range.Base = memory.Pointer(base)
	} else {
		return err
	}

	if size, err := d.Uint64(); err == nil {
		o.Range.Size = size
	} else {
		return err
	}

	if id, err := d.ID(); err == nil {
		o.ID = id
	} else {
		return err
	}

	return nil
}
