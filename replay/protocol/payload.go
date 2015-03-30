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
// See the License for the specific language governing permissions ands
// limitations under the License.

package protocol

import (
	"android.googlesource.com/platform/tools/gpu/binary"
)

// ResourceInfo describes a resource used by a Payload.
type ResourceInfo struct {
	ID   binary.ID // The resource identifier as a string.
	Size uint32    // The size in bytes of the resource.
}

// Payload contains all the information to perform a replay. The encoded form
// is what is passed to the replay system.
type Payload struct {
	StackSize          uint32         // Maximum number of values.
	VolatileMemorySize uint32         // In bytes.
	Constants          []byte         // The constant buffer.
	Resources          []ResourceInfo // Resources used by this replay payload.
	Opcodes            []byte         // The encoded list of opcodes.
}

func (p *Payload) Encode(e binary.Encoder) error {
	if err := e.Uint32(p.StackSize); err != nil {
		return err
	}
	if err := e.Uint32(p.VolatileMemorySize); err != nil {
		return err
	}
	if err := e.Uint32(uint32(len(p.Constants))); err != nil {
		return err
	}
	if err := e.Data(p.Constants); err != nil {
		return err
	}
	if err := e.Uint32(uint32(len(p.Resources))); err != nil {
		return err
	}
	for _, r := range p.Resources {
		if err := e.String(r.ID.String()); err != nil {
			return err
		}
		if err := e.Uint32(r.Size); err != nil {
			return err
		}
	}
	if err := e.Uint32(uint32(len(p.Opcodes))); err != nil {
		return err
	}
	if err := e.Data(p.Opcodes); err != nil {
		return err
	}
	return nil
}

func (p *Payload) Decode(d binary.Decoder) (err error) {
	if p.StackSize, err = d.Uint32(); err != nil {
		return err
	}
	if p.VolatileMemorySize, err = d.Uint32(); err != nil {
		return err
	}
	if count, err := d.Uint32(); err != nil {
		return err
	} else {
		p.Constants = make([]byte, count)
		if err = d.Data(p.Constants); err != nil {
			return err
		}
	}
	if count, err := d.Uint32(); err != nil {
		return err
	} else {
		p.Resources = make([]ResourceInfo, count)
		for i := range p.Resources {
			r := &p.Resources[i]
			if s, err := d.String(); err != nil {
				return err
			} else if r.ID, err = binary.ParseID(s); err != nil {
				return err
			}
			if r.Size, err = d.Uint32(); err != nil {
				return err
			}
		}
	}
	if count, err := d.Uint32(); err != nil {
		return err
	} else {
		p.Opcodes = make([]byte, count)
		if err = d.Data(p.Opcodes); err != nil {
			return err
		}
	}
	return nil
}
